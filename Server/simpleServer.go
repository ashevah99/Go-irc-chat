package main

import (
	"bufio"
	"fmt"
	"net"
)

// Simple server connection (returns nil if successful)
func initServer(port string) error {
	fmt.Println("Launching Server ...")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// accept connection on p
	conn, err := ln.Accept()
	if err != nil {
		return err
	}

	// loop for new connections
	for {
		//listen for message
		msg, err := bufio.NewReader(conn).ReadString('\n')
		// output message received
		if err == nil {
			fmt.Print("Message Received:", string(msg))
		} else {
			fmt.Print("Error: ", err)
		}

		//write back
		conn.Write([]byte("thanks!\n"))
	}
}
