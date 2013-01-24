package gameserver

type Request interface {
	Encode(interface{}) Packet
	Decode(interface{}) Job
	Responce(interface{}) Packet
}

type EchoRequest struct {
	msg string
}

func (e EchoRequest) Encode(a interface{}) Packet {
	var p Packet

	return p
}

func (e EchoRequest) Decode(a interface{}) Job {
	var j Job

	return j
}

func (e EchoRequest) Responce(a interface{}) Packet {
	var p Packet

	return p
}
