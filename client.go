// Copyright 2009 ER Technology Ltda. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt";
	"net";
	"os";
	"bufio";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		msg string = "Doing some stuff\n";
	)

	con, error := net.Dial("tcp", remote);
	defer con.Close();
	if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }

	in, error := con.Write([]byte(msg));
	if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }

	status, error := bufio.NewReader(con).ReadString('\n')
	if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }

	fmt.Printf("Response: %s", status)


	fmt.Fprintf(con, "quit\r\n")
}


