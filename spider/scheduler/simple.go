package scheduler

import "github.com/xxg3053/go-spider/spider/engine"

type SimpleScheduler struct{
	workerChian chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request)  {

}

func (s *SimpleScheduler) Submit(request engine.Request)  {
	go func() {
		s.workerChian <- request
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request{
	return s.workerChian
}

func (s *SimpleScheduler) Run()  {
	s.workerChian = make(chan engine.Request)
}
