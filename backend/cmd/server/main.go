package main

import (
	"log"
	"net/http"
	"whatsapp-clone/pkg/handlers"
	"whatsapp-clone/pkg/websocket"
)

func main() {
	go websocket.HubInstance.Run()
	http.HandleFunc("/", handlers.ServerHome)
	http.HandleFunc("/Ws", handlers.SeveWs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
