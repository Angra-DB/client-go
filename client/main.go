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
	conn := conn_server(os.Args[1], os.Args[2])
	commandProcess(conn)
}

/**
* This function connects to the server based, in TCP protocol.
* It tries to connect. If some error occur, returns the error.
* Otherwise, return the connection.
 */
func conn_server(host string, port string) {
	fmt.Print("Trying to connect to " + host + "....")

	conn, connErr := net.Dial("tcp", host+port)
	if connErr != nil {
		log.Fatal("An error occured: " + connErr)
	}

	fmt.Println("Done")

	return conn
}

/**
* This function read and process the command
* read from the user.
 */
func commandProcess(conn Conn) {
	for {
		bufferIn := bufio.NewReader(os.Stdin)
		fmt.Print("AngraDB> ")

		messageIn, InErr := bufferIn.ReadString(";\n")
		if InErr != nil {
			log.Fatal("An error occured: " + InErr)
		}
		fmt.Fprint(conn, messageIn)

		bufferOut := bufio.NewReader(conn)
		messageOut, OutErr := bufferOut.ReadString("\n")
		if OutErr != nil {
			log.Fatal("An error occured: " + OutErr)
		}
		fmt.Printfln(messageOut)
	}
}
