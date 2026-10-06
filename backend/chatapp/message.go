package main

type MessageType int

const (
	ClientConnected  MessageType = iota
	ClientDisconnected
	ClientSendMessage
)

type Message struct {
	Id int
	MessageType MessageType
	Message string
}