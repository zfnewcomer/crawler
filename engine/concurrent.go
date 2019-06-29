package engine
//engine的功能是在各模块之间传递数据
//go中各模块传递数据使用channel
//接口(这里指interface)由使用者(engine)定义
//操作模块(scheduler)的接口(ConfigureMasterWokerChan、Submit)由scheduler模块具体实现
/*
对面向接口编程的理解：
使用者(caller)需要被调用者(callee)拥有功能A,B，则定义一个包含A()、B()的接口，被调用者去实现该接口
使用者如何获取接口呢？可以将接口定义为使用者struct的一个字段，比如本项目中给ConcurrentEngine
定义了一个Scheduler接口类型的字段
*/
/*
为什么要使用接口，而不是直接使用 "callee.A()"的形式？
假设有2个被调用者callee1、callee2，caller需要使用callee1.A()、callee2.A()去调用
具体到本项目中，则是有多个版本的scheduler，显然这样做是不合适的，抽象出一个接口，
只对接口进行编程是更好的做法
*/

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

//Scheduler接口需要重构
//统一SimpleScheduler和QueuedScheduler的方法
//SimpleScheduler和QueuedScheduler的区别在于：
//在SimpleScheduler中所有worker共用一个channel
//QueuedScheduler中每个worker有自己的channel
//思路：传入一个worker，让scheduler返回channel
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	out:=make(chan ParseResult)
	e.Scheduler.Run()

	//启动WorkerCount个groutine
	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),
			out,e.Scheduler)
	}

	//Submit接口由Scheduler模块提供，request:engine->scheduler
	for _,r :=range seeds{
		e.Scheduler.Submit(r)
	}
	//request:worker->engine
	for{
		result:=<- out
		for _,item:= range result.Items{
			log.Printf("Got item: %v",item)
		}
		for _,request:=range result.Requests{
			//request:engine->scheduler
			e.Scheduler.Submit(request)
		}
	}
}

//本来是将scheduler接口传进来，但是太重了，
//单独拆分有WorkerReady方法的ReadyNotifier接口
func createWorker(in chan Request,
	out chan ParseResult,ready ReadyNotifier)  {
	go func() {
		for{
			ready.WorkerReady(in)
			request:=<-in
			//request：scheduler->worker
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			out<- result
		}
	}()
}

