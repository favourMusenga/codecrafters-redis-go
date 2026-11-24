package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment the code below to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		if !scanner.Scan() {
			break // Client disconnected (EOF), timed out, or read error
		}

		receivedText := strings.TrimSpace(scanner.Text())
		fmt.Println("Received:", receivedText)

		if _, err := conn.Write([]byte("+PONG\r\n")); err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Connection closed with error:", err)
	}
}
