package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("tcp", ":17")

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
	conn.Close()
}