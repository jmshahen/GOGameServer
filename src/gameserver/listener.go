package gameserver

import (
	"strings"
)

func (user UserInfo) listener() {
	conn := user.Conn
	addr := conn.RemoteAddr()
	rw := user.rw
	for {
		s, err := rw.ReadString(user.gs.Terminator)
		if len(s) > 0 {
			logger.Println("[ Listener ] conn", addr, "said", len(s), s)

			packet, perr := user.gs.DecodeString(s)
			if perr != nil {
				logger.Println("[ Listener ] Error decoding packet: ", perr)
				continue
			}

			job, jerr := user.gs.DecodePacket(packet)
			if jerr != nil {
				logger.Println("[ Listener ] Error decoding job: ", jerr)
				continue
			}
			select {
			case user.ch <- job:
				//logger.Println("[listener] DoWork", user.doWork)
				user.doWork <- true //TODO put select statement so it doesn't block
			default:
				logger.Println("[ Listener ] Channel buffer is currently full")
				user.SendMessage("Buffer Is Full")
			}
		} else {
			logger.Printf("[ Listener ] Connection with user lost %#v\n", err)

			user.doWork <- true //is allowed to block
			user.ch <- QUITJob

			conn.Close()
			return
		}
	}
}

func (user UserInfo) SendMessage(s string) error {
	if !strings.HasSuffix(s, string(user.gs.Terminator)) {
		s = s + string(user.gs.Terminator)
	}
	user.rw.WriteString(s)
	return user.rw.Flush()
}
