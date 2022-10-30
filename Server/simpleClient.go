package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func initClient(ip string, port string) error {
	fmt.Println("Starting client")
	//connect to socket
	conn, _ := net.Dial("tcp", ip+":"+port)
	for {
		//read input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		//listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
