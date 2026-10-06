package main

import (
	"sync"
	"net"
	"log"
)

type Client struct {
	Conn net.Conn
	name string
}

type ClientManager struct {
	clients map[*Client]struct{}
	mu	  sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[*Client]struct{}),
	}
}

func (cm *ClientManager) AddClient(client *Client) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[client] = struct{}{}
}

func (cm *ClientManager) RemoveClient(client *Client) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	if _, exists := cm.clients[client]; exists {
		delete(cm.clients, client)
		log.Printf("Client %s removed from the manager", client.name)
	} else {
		log.Printf("Client %s not found in the manager", client.name)
	}
}

func (cm *ClientManager) BroadcastMessage(message string, sender *Client) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for client := range cm.clients {
		if client != sender {
			_, err := client.Conn.Write([]byte(message))
			if err != nil {
				log.Printf("Error sending message to client %s: %v", client.name, err)
			}
		}
	}
}