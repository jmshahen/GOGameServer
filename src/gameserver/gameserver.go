package gameserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type ServerInfo struct {
	Name string // the name of the server
	Port int //the port to liosten on
}

var nl byte = 10

func echo(conn *net.TCPConn) {
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


/*

*/
func StartServer(si ServerInfo) (ServerInfo, error) {
	l, err := net.ListenTCP("tcp4", &net.TCPAddr{net.IPv4zero, si.Port})
	if l == nil {
		fmt.Printf("cannot listen: %s\n", err)
		return si, err
	}
	fmt.Printf("listening at %s\n", l.Addr())
	for {
		conn, err := l.AcceptTCP()
		if conn == nil {
			fmt.Printf("accept error: %s\n", err)
			l.Close()
			return si, err
		}
		fmt.Printf("connection from %s\n", conn.RemoteAddr())
		go echo(conn)
	}

	return si, err
}