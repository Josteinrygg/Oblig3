package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("udp", ":17")
	checkErr(err)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hit enter to connect!")
	txt, _ := reader.ReadString('\n')
	fmt.Fprint(conn, txt+"\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}