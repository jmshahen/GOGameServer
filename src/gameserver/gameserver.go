package gameserver

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

var version string = "0.1.1"
var logger *log.Logger

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
	Separator         byte
	MaxUsersPerWorker int
	Services          []Service
}

type UserInfo struct {
	Id     int
	Name   string
	Conn   net.Conn
	rw     *bufio.ReadWriter
	ch     chan Job
	doWork chan bool
	gs     *GameServer
}

type WorkerInfo struct {
	Id      int
	users   map[int]UserInfo
	addUser chan UserInfo
	quit    chan bool
	doWork  chan bool
}

/*

*/
func (gs GameServer) Init(l *log.Logger) error {
	logger = l
	logger.Println("GO Game Server", version, "by Jonathan Shahen 2013")
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(gs.Port))
	if ln == nil {
		logger.Println("cannot listen:", err)
		return err
	}
	logger.Println("[ Init ] listening at", ln.Addr(), ln)
	gs.ListenConn = ln

	gs.workers = make([]WorkerInfo, 0, gs.NumWorkers)

	for i := 0; i < gs.NumWorkers; i++ {
		w := WorkerInfo{
			i,
			make(map[int]UserInfo),
			make(chan UserInfo, 10),
			make(chan bool),
			make(chan bool)}
		go w.worker()
		logger.Println("[ Init ] Created worker", i)
		gs.workers = append(gs.workers, w)
	}

	return gs.facilitator()
}
