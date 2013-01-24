package gameserver

import (
	"bufio"
)

func (gs GameServer) facilitator() error {
	logger.Println("[ Facilitator ] Started")
	ln := gs.ListenConn
	for {
		logger.Println("[ Facilitator ] Waiting for connection", ln)
		conn, err := ln.Accept()
		if conn == nil {
			logger.Printf("[ Facilitator ] accept error: %s\n", err)
			ln.Close()
			return err
		}
		logger.Println("[ Facilitator ] Connection from", conn.RemoteAddr())

		user := new(UserInfo)
		user.Id = gs.UserCount
		gs.UserCount++
		user.Conn = conn
		user.rw = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		user.ch = make(chan Job, gs.BufferSize)
		user.gs = &gs

		if worker, addError := gs.AddUserToWorker(user); addError != nil {
			logger.Println("[ Facilitator ] Error Server Is Full")
			user.SendMessage("ServerFull")
			continue
		} else {
			logger.Println("[ Facilitator ] User", user.Id, "adding to worker", worker.Id)
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
			return &worker, nil
		}
	}

	return nil, NewServerError("Server Full")
}
