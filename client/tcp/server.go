package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
)

// ConnectServer connects to a server based in TCP protocol.
// It tries to connect. If some error occur, returns the error.
// Otherwise, return the connection.
func ConnectServer(host string, port string) net.Conn {
	fmt.Print("\nTrying to connect to " + host + ":" + port + "....")

	conn, connErr := net.Dial("tcp", host+":"+port)
	if connErr != nil {
		log.Fatal("\nAn error occured: ", connErr)
	}

	fmt.Println("\nDone")

	return conn
}

// DisconnectServer will close the connection with the server
func DisconnectServer(conn net.Conn) {
	conn.Close()
	os.Exit(0)
}
