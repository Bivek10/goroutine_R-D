package helpers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	client ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		client: make(ClientList),
	}

}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("=========NEW CONNECION==========")

	//upgrade the regular http to websocket

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	
	if err != nil {
		log.Println("Error:", err)
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)

	go client.RDMessage()
	//conn.Close()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.client[client] = true

	// if _, ok := m.client[client]; ok{

	// }
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.client[client]; ok {
		client.connection.Close()
		delete(m.client, client)
	}
}
