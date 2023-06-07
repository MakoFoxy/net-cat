package functions

import "net"

func DelWN(delclients map[net.Conn]string, wcon net.Conn) {
	delete(delclients, wcon)
	wcon.Close()
}

func IsUniqueName(resname string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	for _, userName := range clients {
		if userName == resname {
			return false
		}
	}
	return true
}
