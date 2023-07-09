package main

import (
	"net/http"
	"log"
	"os"
	"github.com/gorilla/websocket"
)

// the Message struct
type Message struct {
	message string // the actual message
}

// User.WebsockerConnection

// a map of all connections and if they are active 
var connections = map[*websocket.Conn]bool

// the channel used to receive and send messages
var wschannel = make(chan Message)

var logger = log.New(os.Stdout,"[EVENT] ",log.Ltime)

// upgrader for Websockets
var upgrader = websocket.Upgrader{}

func handleConnection(wsconn *websocket.Conn) {
	defer wsconn.Close()
	 	
	// continiusly read and write messages
	for {
		_, msg, err_msg := wsconn.ReadMessage()
		if err_msg != nil {
			logger.Printf("%s\n", err_msg)
			break
		}
		// store the message in the database
		logger.Printf("Message -> %s\n", string(msg))
	}
}

func main() {
	// File server
	http.Handle("/", http.FileServer(http.Dir("./client")))
	
	// handler for websockets
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		// upgrade to Websockets
		wsconn, err_wsconn := upgrader.Upgrade(w,r,nil)	
		if err_wsconn != nil {
			logger.Printf("%s\n", err_wsconn)
			// Maybe respond with a http.StatusBadRequest
			return
		}
		go handleConnection(wsconn)
	})
	
	logger.Printf("Live on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Printf("%s\n", err)
	}
}
