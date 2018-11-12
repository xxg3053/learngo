package main

import (
	"net/http"
	"github.com/xxg3053/go-spider/frontend/controller"
)

func main()  {
	// 搜索首页
	http.Handle("/",
		http.FileServer(http.Dir("./frontend/view")))
	// 搜索结果页面
	http.Handle("/search",
		controller.CreateSearchResultHandler("./frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}



