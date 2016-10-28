package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

/**
* This is the main function. It recieves two arguments,
* the first one is the IP.
* The second one is the Port.
 */
func main() {
  connIp := os.Args[1]
  port := os.Args[2]
	conn := conn_server(connIp, port)
	commandProcess(conn)
}

/**
* This function connects to the server based, in TCP protocol.
* It tries to connect. If some error occur, returns the error.
* Otherwise, return the connection.
 */
func conn_server(host string, port string) net.Conn {
	fmt.Print("\nTrying to connect to " + host + "....")

	conn, connErr := net.Dial("tcp", host + ":" + port)
	if connErr != nil {
		log.Fatal("\nAn error occured: ", connErr)
	}

	fmt.Println("\nDone")

	return conn
}

/**
* This function read and process the command
* read from the user.
 */
func commandProcess(conn net.Conn) {
	for {
		bufferIn := bufio.NewReader(os.Stdin)
		fmt.Print("\nAngraDB> ")

		messageIn, InErr := bufferIn.ReadString(';')
		if InErr != nil {
			log.Fatal("\nAn error occured: ", InErr)
		}
		fmt.Fprint(conn, messageIn)

		bufferOut := bufio.NewReader(conn)
    fmt.Println("Entrando no reader da saida:")
		_, _,OutErr := bufferOut.ReadLine()
    fmt.Println("Saindo do reader da saida")
		if OutErr != nil {
			log.Fatal("\nAn error occured: ", OutErr)
		}
		// fmt.Println(messageOut)
	}
}
