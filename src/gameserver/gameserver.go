package gameserver

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

var version string = "0.1.0"

type GameServer struct {
	Name       string // the name of the server
	Port       int    //the port to listen on
	ListenConn net.Listener
	workers    []WorkerInfo
	UserCount  int
	//User Settings
	NumWorkers        int
	BufferSize        int
	Terminator        byte
	MaxUsersPerWorker int
}

type UserInfo struct {
	Id         int
	Name       string
	Conn       net.Conn
	rw         *bufio.ReadWriter
	ch         chan WorkerJob
	doWork     chan bool
	terminator byte
}

type WorkerJob struct {
	msg string
}

type WorkerInfo struct {
	Id      int
	users   []UserInfo
	addUser chan UserInfo
	quit    chan bool
	doWork  chan bool
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (gs GameServer) facilitator() error {
	fmt.Println("[Facilitator] Started")
	ln := gs.ListenConn
	for {
		fmt.Println("[Facilitator] Waiting for connection", ln)
		conn, err := ln.Accept()
		if conn == nil {
			fmt.Printf("accept error: %s\n", err)
			ln.Close()
			return err
		}
		fmt.Println("connection from", conn.RemoteAddr())

		user := new(UserInfo)
		user.Id = gs.UserCount
		gs.UserCount++
		user.Conn = conn
		user.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		user.ch = make(chan WorkerJob, gs.BufferSize)
		user.terminator = gs.Terminator

		if worker, addError := gs.AddUserToWorker(user); addError != nil {
			fmt.Println("Error Server Is Full")
			user.SendMessage("ServerFull")
			continue
		} else {
			fmt.Println("User", user.Id, "adding to worker", worker.Id)
			user.SendMessage("Success")
		}
		go user.listener()
	}

	return nil
}

func (gs GameServer) AddUserToWorker(user *UserInfo) (*WorkerInfo, error) {
	for _, worker := range gs.workers {
		if len(worker.users) < gs.MaxUsersPerWorker {
			user.doWork = worker.doWork
			worker.addUser <- *user
			worker.doWork <- true
			fmt.Println("[AddUserToWorker] DoWork", worker.doWork)
			return &worker, nil
		}
	}

	return nil, &MyError{
		time.Now(),
		"Server Full",
	}
}

/*

*/
func (gs GameServer) Init() error {
	fmt.Println("GO Game Server", version, "by Jonathan Shahen 2013")
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(gs.Port))
	if ln == nil {
		fmt.Println("cannot listen:", err)
		return err
	}
	fmt.Println("listening at", ln.Addr(), ln)
	gs.ListenConn = ln

	gs.workers = make([]WorkerInfo, 0, gs.NumWorkers)
	fmt.Println("Cap", cap(gs.workers))

	for i := 0; i < gs.NumWorkers; i++ {
		w := WorkerInfo{
			i,
			make([]UserInfo, 0, gs.MaxUsersPerWorker),
			make(chan UserInfo, 10),
			make(chan bool),
			make(chan bool)}
		go w.worker()
		fmt.Println("Created worker", i)
		gs.workers = append(gs.workers, w)
	}

	return gs.facilitator()
}
