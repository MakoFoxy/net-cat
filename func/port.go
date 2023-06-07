package functions

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func Port() {
	localhost := os.Args[1:]
	listenarr := []string{}
	listenErr := "[USAGE]: ./TCPChat $port"

	listenstr := ""
	if len(localhost) < 1 {
		listenarr = append(listenarr, "8989")
	} else if len(localhost) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else {
		for _, check := range localhost {
			if check >= "0" && check <= "9" {
				num, err := strconv.Atoi(check)
				if err != nil {
					log.Fatal(err)
				}
				if num >= 1024 && num <= 65535 {
					listenarr = append(listenarr, strconv.Itoa(num))
				} else if num > 65535 || num < 1024 {
					fmt.Println("ERROR: not correct port, input from port 1024 to 65535")
					return
				}
			} else {
				for i := 0; i <= len(localhost); i++ {
					if localhost[0][i] == '-' || localhost[0][i] == '+' {
						listenstr = listenstr + listenErr
						fmt.Println(listenstr)
						return
					} else if localhost[0][i] >= 'a' && localhost[0][i] <= 'z' {
						listenstr = listenstr + listenErr
						fmt.Println(listenstr)
						return
					} else if localhost[0][i] >= 'A' && localhost[0][i] <= 'Z' {
						listenstr = listenstr + listenErr
						fmt.Println(listenstr)
						return
					}
				}
			}
		}
	}

	for _, localstr := range listenarr {
		listenstr = listenstr + "localhost:" + localstr
	}
	listener, err := net.Listen("tcp", listenstr)
	// fmt.Println(listenstr)

	log.Println("Listening on the port:", listenstr)

	if err != nil {
		log.Fatal(err)
	}

	go Broadcaster()
	for {
		conn, err := listener.Accept() // консоль открывает
		if err != nil {
			log.Print(err)
			continue
		}
		Retryconn(conn)
	}
}
