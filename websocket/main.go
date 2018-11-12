package main

import (
	"net/http"
	"fmt"
	"github.com/xxg3053/learngo/websocket/handler"
	"html/template"
	"net/url"
)

var (
	tpl = template.Must(template.ParseGlob("./websocket/views/*"))
)

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	if err := tpl.ExecuteTemplate(w, "client.html", nil); err != nil{
		http.Error(w, "error 500:"+" "+err.Error(), http.StatusInternalServerError)
	}
}
func sendHandler(w http.ResponseWriter, r *http.Request)()  {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil{
		http.Error(w, "send error: " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(queryForm["name"]) == 0{
		http.Error(w, "please, put name arg", http.StatusBadRequest)
		return
	}
	if len(queryForm["msg"]) == 0{
		http.Error(w, "please, put msg arg", http.StatusBadRequest)
		return
	}

	handler.Push2TaskCh(handler.Task{Operator: queryForm["name"][0], Content: queryForm["msg"][0]})
	w.Write([]byte("send ok"))
}

func wsHandler(w http.ResponseWriter, r *http.Request)()  {
	hub := handler.NewHub()
	go hub.Run()
	handler.ServeWS(hub, w, r)
}

//client http://localhost:8899/?name=3
//send mssage http://localhost:8899/send?name=3&msg=abc
func main()  {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/ws", wsHandler)

	fmt.Printf("Server starting on %d\n", 8899)
	if err := http.ListenAndServe(":8899", nil); err != nil{
		fmt.Printf("Server start error: %v\n", err)
	}
}
