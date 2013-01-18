package gameserver

import (
	"bufio"
	"fmt"
	"io"
)

var nl byte = 10

func (user User) listener() {
	conn := user.Conn
	addr := conn.RemoteAddr()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		s, err := rw.ReadString(nl)
		if len(s) > 0 {
			fmt.Printf("conn %s said %d %s\n", addr, len(s), s)
			rw.WriteString(s)
			rw.Flush()
		} else if err == io.EOF {
			fmt.Printf("conn %s eof\n", addr)
			conn.Close()
			return
		} else {
			fmt.Printf("error reading: %s\n", err)
			conn.Close()
			return
		}
		if s == "quit\r\n" {
			conn.Close()
			return
		}
	}
}
