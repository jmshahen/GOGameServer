package main

import (
	"gameserver"
	"log"
	"os"
)

func main() {
	var gs gameserver.GameServer
	gs.Name = "Test Server"
	gs.Port = 9989
	gs.BufferSize = 2
	gs.NumWorkers = 3
	gs.MaxUsersPerWorker = 2

	gs.Terminator = '|'
	gs.Separator = '&'

	gs.Services = append(gs.Services, gameserver.BasicService)
	gs.Init(log.New(os.Stdout, "[GameServer]", log.LstdFlags|log.Lshortfile))
}
