package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// check the origin of our connection
	CheckOrigin: func(r *http.Request) bool { return true }}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

// define a reader which will listne for new messages being sent to the Websocket endpoint
// func Reader(conn *websocket.Conn) {
// 	for {
// 		// read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print out the message for clarity
// 		fmt.Println(string(p))
// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func Writer(conn *websocket.Conn) {
// 	for {
// 		fmt.Println("Sending")
// 		messageType, r, err := conn.NextReader()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		w, err := conn.NextWriter(messageType)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if _, err := io.Copy(w, r); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }
