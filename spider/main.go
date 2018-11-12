package main

import (
	"github.com/xxg3053/go-spider/spider/engine"
	"github.com/xxg3053/go-spider/spider/zhenai/parser"
	"github.com/xxg3053/go-spider/spider/scheduler"
	"github.com/xxg3053/go-spider/spider/persist"
)

func main()  {
	/**
	engine.SimpleEngine{}.Run(engine.Request{
		Url: "http://city.zhenai.com/",
		ParserFunc: parser.ParserCityList,
	})
	**/
	/**
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url: "http://city.zhenai.com/",
		ParserFunc: parser.ParserCityList,
	})**/

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url: "http://city.zhenai.com/",
		ParserFunc: parser.ParserCityList,
		//Url: "http://city.zhenai.com/shenzhen",
		//ParserFunc: parser.ParserCity,
	})
}

