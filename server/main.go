package main

// REFS:
// https://gist.github.com/drewolson/3950226
// https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
// https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	// accept connection on port
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err) //TODO
		}
		// output message received
		fmt.Print("Message Received:", string(message))

		// THIS IS THE SILLY RETURN STUFF
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
