package websocket

import (
	
	"log"
	"net/http"
   "github.com/gorilla/websocket"
	
)
var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	checkOrigin: func(r *http.Request) bool {
		return true // Allow all connection by default
	},
}
	type Client struct {
		conn *websocket.Com
		send chan[]byte
	}
	type Hub struct {
		Client map[*client]bool
		Broadcast chan []byte
		Register chan *client 
		unregister chan *client 
	}
	var HubInstance = &Hub{
		Client:   make(map[*Client]bool),
		Broadcast:  make(chan[]byte),
		Register:  make(chan *client),
		Unregister: make(chan *client),
	}
	func (c *Client) Write() {
		defer func(){
			HubInstance.Unregister <- c
			c.Conn.Close()
		}()
		for {
			_, message,err := c.Conn.ReadMessage()
			if err != nil {
				log.Println("read error:",err)
				break
			}
			HubInstance.Broadcast <- message
		}
	}
	func(c *Client) Writer(){
		defer func(){
			c.Conn.Close()
		} ()
		for {
			select {
			case message, ok := <-c.Send:
			if !ok{
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return 
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
func (h *Hub) Run(){
	for {
		select {
		case client := <-h.Register:
			h.clients[client]= true 
		case client := <-h.Unregister:
			if _, ok := h.Client[client]; ok{
				delete(h.Client, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
func ServeWs(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil 
	log.Println("upgrade error:", err)
	return 
}
client := &Client{Conn: conn, Send: make(chan []byte)}
HubInstance.Register <- client
go client.Read()
go client.Write()
}