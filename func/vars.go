package functions

import (
	"net"
	"sync"
)

type Message struct {
	Name string
	Text string
}

type History struct {
	arrhistory []Message
}

var (
	delclients = make(map[net.Conn]string) // Все отключенные клиенты
	clients    = make(map[net.Conn]string) // Все подключенные клиенты
	messages   = make(chan Message)        // Все входящие сообщения клиента
	history    = History{}
	mutex      sync.Mutex
)
