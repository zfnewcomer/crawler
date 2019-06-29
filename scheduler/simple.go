package scheduler

import "newproject/crawler/engine"

//使用IDE的implement interface功能自动生成Submit方法
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan=make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	//submit request to WokerChan
	//当request多了的时候Submit会卡住，使用groutine解决
	//还可以使用队列解决
	go func() {
		s.workerChan <- r
	}()
}




