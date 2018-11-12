package handler

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"net/url"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)


type Client struct {
	name  string
	hub   *Hub
	conn  *websocket.Conn
	send  chan []byte
	close chan bool
}

func (c *Client) read()  {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	taskChan := GetOrAddTask(c.name)
	go func() {
	L:
		for {
			select {
			case t := <- taskChan:
				task,_ := json.Marshal(t)
				c.hub.broadcast <- Message{toClients: []*Client{c}, content: task}
			case <-c.close:
				break L
			}
		}
	}()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("[WS] websocket closede : %v\n", err)
			}
			break
		}
		c.send <- data
	}

}

func (c *Client) write()  {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//握手
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil{
		http.Error(w, "send error: " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(queryForm["name"]) == 0{
		http.Error(w, "please, put name arg", http.StatusBadRequest)
		return
	}
	name := queryForm["name"][0]
	conn, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		fmt.Printf("[WS] websocket upgrade failed: %v\n", e)
		return
	}
	client := &Client{name: name, hub: hub,conn:conn, send: make(chan []byte, 256), close:make(chan bool, 1)}
	client.hub.register <- client

	go client.write()
	go client.read()
}

// Hub maintains the set of active clients and broadcasts messages to specific clients.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

// Message contains message content and message receiver
type Message struct {
	toClients []*Client
	content   []byte
}

// NewHub constructs a hub instance.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run starts hub.
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				c.close <- true
				close(c.send)
			}
		case message := <-h.broadcast:
			for _, c := range message.toClients {
				if _, ok := h.clients[c]; ok {
					c.send <- message.content
				}
			}
		}
	}
}

//https://github.com/owenliang/go-push
