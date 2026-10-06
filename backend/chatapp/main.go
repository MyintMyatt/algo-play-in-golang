package main

import (
	"log"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "6969"
	SERVER_TYPE = "tcp"
)

func main() {
	clientManager := NewClientManager()
	server, err := net.Listen(SERVER_TYPE, ":"+SERVER_PORT)

	if err != nil {
		log.Fatal("Error starting server:", err)
		return
	}

	log.Printf("Server is running on : %s \n", SERVER_PORT)
	defer server.Close()
	
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, clientManager)
	}
}

func handleConnection(conn net.Conn, clientManager *ClientManager) {
	defer conn.Close()
	client := &Client{
		Conn: conn,
		name: conn.RemoteAddr().String(),
	}
	clientManager.AddClient(client)
	defer func() {
        clientManager.RemoveClient(client)
        conn.Close()
    }()
	log.Printf("Client connected: %v", conn.RemoteAddr())


	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from client:", err)
			return
		}
		broadcastMessage(string(buffer[:n]), client, clientManager)
	}
}

func broadcastMessage(message string, sender *Client, clientManager *ClientManager) {
	clientManager.BroadcastMessage(message, sender)
}