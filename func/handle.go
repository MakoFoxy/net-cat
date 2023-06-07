package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func Retryconn(conn net.Conn) {
	mutex.Lock()
	if len(clients) < 10 {
		go HandleConn(conn)
	} else {
		arrconn := []net.Conn{}
		arrconn = append(arrconn, conn)
		for _, wcon := range arrconn {
			if wcon == conn {
				fmt.Fprintln(conn, "ERROR: not places for you in this Chat")
				wcon = conn
				del := "del"
				delclients[wcon] = del
				delete(delclients, wcon)
				wcon.Close()
				// DelWN(delclients, wcon)
				// Retryconn(wcon)
				// fmt.Println(len(delclients))
				break
			}
		}
	}
	mutex.Unlock()
}

func HandleConn(conn net.Conn) {
	// cwnow := RetryW(wnow)
	// mutex.Lock()
	resname := Retry(conn)
	//	Nicknameretry(conn, resname)
	if IsUniqueName(resname) {
		mutex.Lock()
		clients[conn] = resname
		mutex.Unlock()
	} else {
		resname = Retry(conn)
	}
	// mutex.Unlock()
	// fmt.Println(len(clients))
	allstory := []Message{}

	for _, allhistory := range history.arrhistory {
		allstory = append(allstory, allhistory)
	}

	for _, w := range allstory {
		if w.Text != "{" || w.Text != "}" {
			fmt.Fprintln(conn, w.Text)
		} else {
			break
		}
	}

	messages <- Message{Name: "", Text: resname + " conecting"}
	history.arrhistory = append(history.arrhistory, Message{Text: resname + " conecting"})

	inputtxt := bufio.NewScanner(conn)
	var newMess Message

	for inputtxt.Scan() {
		if strings.TrimSpace(inputtxt.Text()) == "" {
			fmt.Fprintln(conn, "You not can send message")
			conn.Write([]byte(MakeFormat(resname, inputtxt.Text())))
			// fmt.Sprintln(conn, MakeFormat(resname, inputtxt.Text()))
			// messages <- newMess
		} else {
			for _, checkrune := range inputtxt.Text() {
				if checkrune >= rune(32) && checkrune <= rune(126) {
					newMess = Message{Name: resname, Text: MakeFormat(resname, inputtxt.Text())}
					messages <- newMess
					// arrhistory = append(arrhistory, inputtxt.Text())
					history.arrhistory = append(history.arrhistory, Message{Text: MakeFormat(resname, inputtxt.Text())})
					break
				} else {
					fmt.Fprintln(conn, "You not can used others symbols")
					conn.Write([]byte(MakeFormat(resname, inputtxt.Text())))
					// fmt.Sprintln(conn, MakeFormat(resname, inputtxt.Text()))
					// fmt.Fprintln(conn, Message{Text: MakeFormat(resname, newMess.Text)})
					// messages <- newMess
					break
				}
			}
		}
	}
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	newMess = Message{Name: resname, Text: resname + " disconnecting"}
	messages <- newMess

	conn.Close()
}

func MakeFormat(resname, txt string) string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + resname + ": " + strings.TrimSpace(txt)
}
