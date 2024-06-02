package handlers

import (
	"net/http"
	"whatsapp-clone/pkg/websocket"
)

func ServerHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
func SeveWs(w http.ResponseWriter, r *http.Request) {
	websocket.ServeWs(w, r)
}
