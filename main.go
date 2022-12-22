package main

import (
	"bufio"
	"log"
	"net"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	openConnections = make(map[net.Conn]bool)
	newConnections  = make(chan net.Conn)
	deadConnections = make(chan net.Conn)
)

func main() {
	ln, err := net.Listen("tcp", ":8888")
	check(err)
	defer ln.Close()

	go func() {
		for {
			conn, err := ln.Accept()
			check(err)
			openConnections[conn] = true
			newConnections <- conn
		}
	}()

	for {
		select {
		case conn := <-newConnections:
			go clientMessage(conn)

		case conn := <-deadConnections:
			for item := range openConnections {
				if item == conn {
					break
				}
			}
			delete(openConnections, conn)
		}
	}
}

func clientMessage(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			break
		}

		for item := range openConnections {
			if item != conn {
				item.Write([]byte(message))
			}
		}
	}

}
