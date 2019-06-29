package scheduler

import "newproject/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(
	w chan engine.Request){
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan=make(chan chan engine.Request)
	s.requestChan=make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//如果requestQ和workerQ同时不为空，则将
			//activeRequest初始化为requestQ[0]、activeWorker初始化为workerQ[0]
			//这样activeWorker才能接收activeRequest传过来的Request
			if len(requestQ)>0 &&
				len(workerQ)>0{
				activeRequest=requestQ[0]
				activeWorker=workerQ[0]
			}
			//收到worker和收到request是独立事件，所以使用select
			select {
			case r:=<-s.requestChan:
				//收到request缓存起来
				requestQ=append(requestQ,r)
			case w:=<-s.workerChan:
				//收到worker缓存起来
				workerQ=append(workerQ,w)
			case activeWorker<-activeRequest:
				workerQ=workerQ[1:]
				requestQ=requestQ[1:]
			}
		}
	}()
}

