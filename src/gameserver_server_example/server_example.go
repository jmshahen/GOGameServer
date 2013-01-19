package main

import (
	"gameserver"
)

func main() {
	var gs gameserver.GameServer
	gs.Name = "Test Server"
	gs.Port = 9989
	gs.BufferSize = 2
	gs.NumWorkers = 3
	gs.MaxUsersPerWorker = 2
	gs.Terminator = '|'
	gs.Init()
}
