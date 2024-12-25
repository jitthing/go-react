package main

import (
	"fmt"
	"net/http"

	"github.com/jitthing/realtime-chat-go-react/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	// upgrade this connection to a WebSocket
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	// listen indefinitely for new messages coming through on our WebSocket connection
	// go websocket.Writer(ws)
	// websocket.Reader(ws)

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Simple Server")
	// })
	pool := websocket.NewPool()
	go pool.Start()

	// map the '/ws' endpoint to the serveWs function
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
