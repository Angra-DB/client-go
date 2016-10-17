package client

import (
  "fmt"
  "net"
  "os"
  "bufio"
)

func main() {
  conn := conn_server(os.Args[1], os.Args[2])
  processaComandos(conn)
}

def conn_server(host string, port string) {
  fmt.Print("Trying to connect to " + host + "....")

  conn, connErr := net.Dial("tcp", host + port)
  if connErr != nil {
    log.Fatal("An error occured: " + connErr)
  }

  fmt.Println("Done")

  return conn
}

def processaComandos(conn Conn) {
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
