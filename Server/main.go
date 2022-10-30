package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go_irc")
	cmdlineArgs()
}

/* Verify & pass commandline args */
func cmdlineArgs() int {
	// if arglen := len(os.Args[1:]); arglen < 2 {
	// 	fmt.Println("Not enough args ")
	// 	return
	// }

	var mode string

	// flags declaration using flag package
	flag.StringVar(&mode, "mode", "client", "Specify mode to run program (client or server). Default is client")

	//client vars
	var ip, port string
	flag.StringVar(&ip, "ip", "", "Specify ip to connect to (client only)")
	flag.StringVar(&port, "port", "", "Specify port to connect to (client) or port to listen on (server) ")

	flag.Parse()

	//client check
	if (ip == "" || port == "") && (mode == "client") {
		fmt.Println("Requires ip and port for client")
		flag.PrintDefaults()
		return 1
	}

	if (port == "") && (mode == "server") {
		fmt.Println("Requires port for server")
		flag.PrintDefaults()
		return 1
	}

	switch mode {
	case "client":
		if err := initClient(ip, port); err != nil {
			fmt.Println("Error starting Client: ", err)
			return 0
		}
	case "server":
		if err := initServer(port); err != nil {
			fmt.Println("Error starting Server: ", err)
			return 0
		}
	}

	fmt.Println("Invalid mode, specify mode to run: client, server")
	flag.PrintDefaults()
	return 1
}
