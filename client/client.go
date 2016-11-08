package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connIp := os.Args[1]
	port := os.Args[2]
	conn := tcp.ConnectServer(connIp, port)

	communicateWithServer(conn)
}

func communicateWithServer(conn net.Conn) {
	for {
		bufferOut := bufio.NewReader(os.Stdin)
		fmt.Print("\nAngraDB> ")

		userMessage, messageErr := bufferOut.ReadString(';')
		if messageErr != nil {
			log.Fatal("\nAn error occured: ", messageErr)
		}
		userMessage = userMessage[:len(userMessage)-1]

		if userMessage == "quit" {
			tcp.Disconnect(conn)
		}

		if tcp.ProcessCommand(userMessage) {
			fmt.Fprint(conn, userMessage)

			bufferIn := bufio.NewReader(conn)
			response, responseErr := bufferIn.ReadString('\n')
			if responseErr != nil {
				log.Fatal("\nAn error occured: ", responseErr)
			}
			fmt.Println(response)
		}
	}
}
