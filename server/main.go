package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8881"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server is running at :" + SERVER_PORT)
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	defer server.Close()
	
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
	}

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for connection...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go processClient(conn)
	}

}

func processClient(conn net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = conn.Write([]byte("Thank you! I got your message: " + string(buffer[:mLen])))
	conn.Close()
}
