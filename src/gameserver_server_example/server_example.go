package main

import (
	"gameserver"
)

func main() {
	var si gameserver.ServerInfo
	si.Name = "Test Server"
	si.Port = 9989
	gameserver.StartServer(si)
}
