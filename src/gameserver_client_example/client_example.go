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

var version string = "0.0.6"
var terminator byte = '|'
var separator byte = '&'

func main() {
	fmt.Println("GO Game Server Client Example", version, "by Jonathan Shahen 2013")
	rand.Seed(time.Now().UnixNano())
	var (
		host   = "127.0.0.1"
		port   = "9989"
		remote = host + ":" + port
		msg    = "Random Number: " + strconv.Itoa(rand.Intn(9999)) + "|"
		err    error
	)

	con, err := net.Dial("tcp", remote)

	if err != nil {
		fmt.Printf("Host not found: %s\n", err)
		os.Exit(1)
	}

	var bconn = bufio.NewReader(con)
	status, err := bconn.ReadString(terminator)
	if err != nil {
		fmt.Println("Error reading data:", err, ", in:", status)
		os.Exit(2)
	}
	if status == "Success"+string(terminator) {
		fmt.Println("Successfully connected to the server!")
	} else {
		fmt.Println("An internal server err occured:", status)
		os.Exit(500)
	}

	var stdinR = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Message:", msg)
		in, err := con.Write([]byte(MakeEchoPacket(msg)))
		if err != nil {
			fmt.Println("Error sending data:", err, ", in:", in)
			os.Exit(2)
		}

		status, err := bconn.ReadString(terminator)
		if err != nil {
			fmt.Println("Error reading data:", err, ", in:", status)
			os.Exit(2)
		}

		fmt.Println("Response:", status)

		fmt.Print("Your Message: ")
		// n, err := fmt.Scanln(os.Stdin, &msg)
		// fmt.Println("n:", n, "| err:", err)
		msg, err = stdinR.ReadString('\n')
		fmt.Println("| err:", err)

		//msg = "Random Number: " + strconv.Itoa(rand.Intn(9999)) + "\n"

		if msg == "quit" {
			fmt.Println("Sending the QUIT command.")
			break
		}
	}

	fmt.Fprintf(con, MakeQuitPacket(""))
	con.Close()
	fmt.Println("Goodbye.")
}

func MakeQuitPacket(s string) string {
	//strips off the new line at the end of the string
	s = strings.TrimRight(s, "\r\n ")

	if !strings.HasSuffix(s, string(terminator)) {
		s = s + string(terminator)
	}
	return "QUIT" + string(separator) + s
}

func MakeEchoPacket(s string) string {
	//strips off the new line at the end of the string
	s = strings.TrimRight(s, "\r\n ")

	if !strings.HasSuffix(s, string(terminator)) {
		s = s + string(terminator)
	}
	return "ECHO" + string(separator) + s
}
