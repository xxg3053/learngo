package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	// init
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i<e.WorkerCount; i++{
		createrWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// submit
	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out
		for _, item := range result.Items{
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createrWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}