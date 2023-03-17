package helpers

import (
	"log"

	"github.com/gorilla/websocket"
)


type Client struct{
	connection *websocket.Conn
	manager *Manager
}

type ClientList map[*Client] bool


func NewClient(conn *websocket.Conn, manager *Manager) *Client{
	return &Client{
		connection: conn,
		manager: manager,
	}
}


func (c *Client) RDMessage(){
	defer func ()  {
		//clean up connection
		c.manager.removeClient(c)
	}()

	for{
		messageType, payload, err:=c.connection.ReadMessage()

		if err !=nil{
			log.Println("Error", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
				log.Printf("error reading message: %v", err)
			}
			break

		}

		log.Println(messageType)
		log.Println(string(payload ))
	}
}