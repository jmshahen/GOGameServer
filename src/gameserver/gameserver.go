package gameserver

import (
	"fmt"
	"net"
	"strconv"
)

var version string = "0.0.4"

type GameServer struct {
	Name       string // the name of the server
	Port       int    //the port to listen on
	ListenConn net.Listener
	NumWorkers int
	BufferSize int
	Terminator byte
}

type User struct {
	Id         int
	Name       string
	Conn       net.Conn
	ch         chan string
	terminator byte
}

func (gs GameServer) facilitator() error {
	ln := gs.ListenConn
	for {
		conn, err := ln.Accept()
		if conn == nil {
			fmt.Printf("accept error: %s\n", err)
			ln.Close()
			return err
		}
		fmt.Printf("connection from %s\n", conn.RemoteAddr())
		user := new(User)
		user.Conn = conn
		user.ch = make(chan string, gs.BufferSize)
		user.terminator = gs.Terminator
		go user.listener()
	}

	return nil
}

/*

*/
func (gs GameServer) Init() error {
	fmt.Println("GO Game Server", version, "by Jonathan Shahen 2013")
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(gs.Port))
	if ln == nil {
		fmt.Printf("cannot listen: %s\n", err)
		return err
	}
	fmt.Printf("listening at %s\n", ln.Addr())

	gs.ListenConn = ln

	return gs.facilitator()
}
