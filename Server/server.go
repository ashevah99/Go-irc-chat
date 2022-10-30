/*
Server to handle connections:
	- open connection to receive process
	- write string
	- close connection
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	port = ":6969"
)

func Open(addr string) (*bufio.ReadWriter, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil

}

func Listen() (net.Listener, error) {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return ln, nil
}

func handleStrings(rw *bufio.ReadWriter) {
	log.Print("Receive STRING message:")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection.\n", err)
	}
	s = strings.Trim(s, "\n ")
	log.Println(s)
	_, err = rw.WriteString("Thank you.\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}
