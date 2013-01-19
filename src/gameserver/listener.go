package gameserver

import (
	"fmt"
	"io"
)

func (user UserInfo) listener() {
	conn := user.Conn
	addr := conn.RemoteAddr()
	rw := user.rw
	for {
		s, err := rw.ReadString(user.terminator)
		if len(s) > 0 {
			fmt.Println("conn", addr, "said", len(s), s)
			select {
			case user.ch <- WorkerJob{s}:
				fmt.Println("[listener] DoWork", user.doWork)
				user.doWork <- true //TODO put select statement so it doesn't block
			default:
				fmt.Println("Channel buffer is currently full")
				user.SendMessage("Buffer Is Full")
			}
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

func (user UserInfo) SendMessage(s string) error {
	user.rw.WriteString(s + string(user.terminator))
	return user.rw.Flush()
}
