package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", ":8888")
	check(err)
	defer connection.Close()

	fmt.Println("Type ur name:")
	username, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)

	username = strings.Trim(username, " \r\n")

	startmessage := fmt.Sprintf("%s welcome to TCP chat \n", username)

	fmt.Println(startmessage)
	go read(connection)
	write(connection, username)

}

func read(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			conn.Close()
			fmt.Println("conn was closed")
			os.Exit(0)
		}
		fmt.Println(message)

	}

}

func write(conn net.Conn, username string) {
	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			break
		}
		message = fmt.Sprintf("%s:- %s\n", username, strings.Trim(message, " \r\n"))
		conn.Write([]byte(message))
	}
}
