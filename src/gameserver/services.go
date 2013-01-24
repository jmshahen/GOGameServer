package gameserver

type Service struct {
	Name string
	Jobs []Job
}

func DecodePacket(p Packet) (Job, error) {
	var j Job
	var e error

	return j, e
}

var BasicService Service = Service{"Basic Services", []Job{ECHOJob, QUITJob}}
