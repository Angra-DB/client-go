package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"client/utils"
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

// ProcessCommand processes the command read from the user. It separates the payload
// from the message, and verifies if it are a valid JSON string.
func ProcessCommand(message string) bool {

	var payload string
	var verifyJSON bool
	if strings.HasPrefix(message, "save ") {

		payload = message[len("save "):len(message)]
		verifyJSON = true

	} else if strings.HasPrefix(message, "update ") {

		payload = message[len("update "):len(message)]
		verifyJSON = true

	} else if strings.HasPrefix(message, "delete ") && strings.HasPrefix(message, "lookup ") &&
		strings.HasPrefix(message, "create_db ") && strings.HasPrefix(message, "connect ") {

		verifyJSON = false

	}

	if verifyJSON && !utils.IsJSON(payload) {
		fmt.Println("Não é JSON. Digite novamente")
		return false
	}

	return true
}
