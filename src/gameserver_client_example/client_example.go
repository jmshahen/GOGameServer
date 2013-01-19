package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var version string = "0.0.5"
var terminator byte = '|'

func main() {
	fmt.Println("GO Game Server Client Example", version, "by Jonathan Shahen 2013")
	rand.Seed(time.Now().UnixNano())
	var (
		host   = "127.0.0.1"
		port   = "9989"
		remote = host + ":" + port
		msg    = "Random Number: " + strconv.Itoa(rand.Intn(9999)) + "|"
	)

	con, error := net.Dial("tcp", remote)

	if error != nil {
		fmt.Printf("Host not found: %s\n", error)
		os.Exit(1)
	}

	var bconn = bufio.NewReader(con)
	status, error := bconn.ReadString(terminator)
	if error != nil {
		fmt.Println("Error reading data:", error, ", in:", status)
		os.Exit(2)
	}
	if status == "Success"+string(terminator) {
		fmt.Println("Successfully connected to the server!")
	} else {
		fmt.Println("An internal server error occured:", status)
		os.Exit(500)
	}

	for {
		fmt.Println("Message:", msg)
		in, error := con.Write([]byte(msg))
		if error != nil {
			fmt.Println("Error sending data:", error, ", in:", in)
			os.Exit(2)
		}

		fmt.Println("in:", in)

		status, error := bconn.ReadString(terminator)
		if error != nil {
			fmt.Println("Error reading data:", error, ", in:", status)
			os.Exit(2)
		}

		fmt.Println("Response:", status)

		fmt.Print("Your Message: ")
		n, errs := fmt.Scan(&msg)
		fmt.Println("n:", n, "| errs:", errs)

		//msg = "Random Number: " + strconv.Itoa(rand.Intn(9999)) + "\n"

		if msg == "quit" {
			fmt.Println("Sending the QUIT command.")
			break
		}
		if !strings.HasSuffix(msg, string(terminator)) {
			fmt.Println("Putting", terminator, "at the end of msg", len(msg))
			msg = msg + string(terminator)
		}
	}

	fmt.Fprintf(con, "quit|")
	con.Close()
	fmt.Println("Goodbye.")
}
