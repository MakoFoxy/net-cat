package functions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

func Retry(conn net.Conn) string {
	filename, err := ioutil.ReadFile("pingwin.txt")
	if err != nil {
		fmt.Fprint(conn, "ERROR: can't read file")
	}
	fmt.Fprintln(conn, string(filename))

	eyn := "[ENTER YOUR NAME]:"
	fmt.Fprint(conn, eyn)

	resname := ""
	input := bufio.NewScanner(conn)
	if input.Scan() {

		text := strings.TrimSpace(input.Text())
		if text >= string(rune(32)) && text <= string(rune(126)) {
			if len(text) >= 3 {
				resname = resname + "[" + text + "]"
			} else {
				fmt.Fprintln(conn, "Please enter you name with length 3 and more symbols")
				resname = Retry(conn)
			}
		} else {
			fmt.Fprintln(conn, "Please enter you name latin symbols")
			resname = Retry(conn)
		}
		if IsUniqueName(resname) == false {
			fmt.Fprintln(conn, "You not can append in the chat, this nickname busy, create new nickname")
			resname = Retry(conn)
		}
	}
	return resname
}
