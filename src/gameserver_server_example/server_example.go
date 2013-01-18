package main

import (
	"gameserver"
)

func main() {
	var gs gameserver.GameServer
	gs.Name = "Test Server"
	gs.Port = 9989
	gs.Init()
}
