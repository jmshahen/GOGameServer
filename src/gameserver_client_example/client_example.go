package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//fmt.Printf("1\n")
	var (
		host   = "127.0.0.1"
		port   = "9989"
		remote = host + ":" + port
		msg    = "Doing some stuff\n"
	)
	//fmt.Printf("2(%s)\n", remote)
	con, error := net.Dial("tcp", remote)
	//fmt.Printf("3(%s)\n", con)

	if error != nil {
		fmt.Printf("Host not found: %s\n", error)
		os.Exit(1)
	}

	//fmt.Printf("4\n")
	in, error := con.Write([]byte(msg))
	if error != nil {
		fmt.Printf("Error sending data: %s, in: %d\n", error, in)
		os.Exit(2)
	}

	status, error := bufio.NewReader(con).ReadString('\n')
	if error != nil {
		fmt.Printf("Error sending data: %s, in: %d\n", error, in)
		os.Exit(2)
	}

	fmt.Printf("Response: %s", status)

	fmt.Fprintf(con, "quit\r\n")
	con.Close()
}
