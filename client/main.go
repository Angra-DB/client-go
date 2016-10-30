package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

/**
* This is the main function. It recieves two arguments,
* the first one is the 	 IP.
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

	conn, connErr := net.Dial("tcp", host+":"+port)
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
		messageIn = messageIn[:len(messageIn)-1]

		if messageIn == "quit" {
			conn.Close()
			os.Exit(0)
		}

		var json_buffer string
		if strings.HasPrefix(messageIn, "save") {

			json_buffer = messageIn[len("save "):len(messageIn)]

		} else if strings.HasPrefix(messageIn, "lookup") {

			json_bufferB, _ := json.Marshal(messageIn[len("lookup "):len(messageIn)])
			json_buffer = string(json_bufferB)

		} else if strings.HasPrefix(messageIn, "delete") {

			json_bufferB, _ := json.Marshal(messageIn[len("delete "):len(messageIn)])
			json_buffer = string(json_bufferB)

		} else if strings.HasPrefix(messageIn, "update") {
			
			json_buffer = messageIn[len("update "):len(messageIn)]

		} else if strings.HasPrefix(messageIn, "create_db") {

			json_buffer = messageIn[len("create_db "):len(messageIn)]

		} else if strings.HasPrefix(messageIn, "connect") {

			json_buffer = messageIn[len("connect "):len(messageIn)]

		}

		if isJSON(json_buffer) || isJSONString(json_buffer) {
			fmt.Println("É JSON.")

			fmt.Fprint(conn, messageIn)

			bufferOut := bufio.NewReader(conn)
			messageOut, OutErr := bufferOut.ReadString('\n')
			if OutErr != nil {
				log.Fatal("\nAn error occured: ", OutErr)
			}
			fmt.Println(messageOut)
		} else {
			fmt.Println("Não é JSON.")
		}
	}
}
