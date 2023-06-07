package functions

import (
	"fmt"
	"log"
)

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for conn, user := range clients {
				if msg.Text == "" {
					if user == msg.Name {
						_, err := fmt.Fprint(conn, "\033[2K\r"+MakeFormat(user, ""))
						if err != nil {
							log.Println(err)
							return
						}
					}
					continue
				}
				if user != msg.Name {
					if msg.Text != "You not can send message" {
						_, err := fmt.Fprint(conn, "\033[2K\r"+msg.Text+"\n")
						if err != nil {
							log.Println(err)
							return
						}
					}
				} else {
				}
				_, err := fmt.Fprint(conn, "\033[2K\r"+MakeFormat(user, ""))
				if err != nil {
					log.Println(err)
					return
				}
			}
			mutex.Unlock()
		}
	}
}
