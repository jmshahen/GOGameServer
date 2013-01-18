package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var version string = "0.0.2"

func main() {
	fmt.Println("GO Game Server Client Example", version, "by Jonathan Shahen 2013")
	rand.Seed(time.Now().UnixNano())
	var (
		host   = "127.0.0.1"
		port   = "9989"
		remote = host + ":" + port
		msg    = "Random Number: " + strconv.Itoa(rand.Intn(9999)) + "\n"
	)

	con, error := net.Dial("tcp", remote)

	if error != nil {
		fmt.Printf("Host not found: %s\n", error)
		os.Exit(1)
	}

	fmt.Println("Message:", msg)
	in, error := con.Write([]byte(msg))
	if error != nil {
		fmt.Println("Error sending data:", error, ", in:", in)
		os.Exit(2)
	}

	status, error := bufio.NewReader(con).ReadString('\n')
	if error != nil {
		fmt.Println("Error sending data:", error, ", in:", in)
		os.Exit(2)
	}

	fmt.Println("Response:", status)

	fmt.Fprintf(con, "quit\r\n")
	con.Close()
}
