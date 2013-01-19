package gameserver

import (
	"bufio"
	"fmt"
	"io"
)

func (user User) listener() {
	conn := user.Conn
	addr := conn.RemoteAddr()
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for {
		s, err := rw.ReadString(user.terminator)
		if len(s) > 0 {
			fmt.Println("conn", addr, "said", len(s), s)
			select {
			case user.ch <- s:
				rw.WriteString(s)
			default:
				fmt.Println("Channel buffer is currently full")
				rw.WriteString("Buffer Is Full" + string(user.terminator))
			}
			rw.Flush()
		} else if err == io.EOF {
			fmt.Println("conn", addr, "eof")
			conn.Close()
			return
		} else {
			fmt.Println("error reading:", err)
			conn.Close()
			return
		}
		if s == "quit"+string(user.terminator) {
			fmt.Println("client sent quiting command")
			conn.Close()
			return
		}
	}
}
