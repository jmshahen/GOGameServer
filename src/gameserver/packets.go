package gameserver

import (
	"strings"
)

type Packet struct {
	Header string
	Body   []string
}

func (gs GameServer) SafePacketString(s string) bool {
	return !strings.ContainsAny(s, string(gs.Terminator))
}

func (gs GameServer) DecodeString(s string) (Packet, error) {
	var p Packet

	s = strings.TrimRight(s, string(gs.Terminator))
	z := strings.SplitN(s, string(gs.Separator), 2)

	if len(z) < 2 {
		return p, NewNetworkError("Invalid Packet Format", s)
	}
	p.Header = z[0]
	p.Body = z[1:]

	return p, nil
}
