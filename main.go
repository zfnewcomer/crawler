package main

import (
	"newproject/crawler/engine"
	"newproject/crawler/scheduler"
	"newproject/crawler/zhenai/parser"
)

func main() {
	e:=engine.ConcurrentEngine{
		//scheduler.SimpleScheduler实现了Scheduler接口
		//SimpleScheduler的方法接收者是指针，所以接口变量自带指针，赋值时右边需要使用&取地址
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:10,
	}
	//Run是指针接收者，直接使用engine.ConcurrentEngine{}.Run()编译报错，
	//因为engine.ConcurrentEngine{}无法取到地址，定义一个临时变量e解决
	//e.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun",
	//	ParserFunc:	parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc:parser.ParseCity,
	})

}


