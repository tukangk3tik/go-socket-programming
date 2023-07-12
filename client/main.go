package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8881"
	SERVER_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hellow Felix! Greetings"))
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
}
