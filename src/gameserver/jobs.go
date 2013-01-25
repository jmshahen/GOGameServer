package gameserver

import (
	"fmt"
	"strings"
)

type Job struct {
	Description string
	Header      string
	Run         func(WorkerInfo, UserInfo, []string) error
	Data        []string
}

func (gs GameServer) DecodePacket(p Packet) (Job, error) {
	var job Job

	for _, service := range gs.Services {
		for _, job := range service.Jobs {
			if job.Header == p.Header {
				var jnew Job = job
				jnew.Data = p.Data
				return jnew, nil
			}
		}

	}

	return job, NewServerError(fmt.Sprintf("Could not find Job %s", p.Header))
}

// Common Jobs
var (
	QUITJob = Job{
		"When the client is about to leave they will send this to logoff.",
		"QUIT",
		func(wi WorkerInfo, ui UserInfo, s []string) error {
			for i, user := range wi.users {
				if user.Id == ui.Id {
					logger.Println("Deleting user", user.Id,
						"from worker", wi.Id)
					delete(wi.users, i)
					return nil
				}
			}

			return NewServerError(
				fmt.Sprintf("User %d Not Found In Worker %d's Queue",
					ui.Id, wi.Id))
		},
		[]string{},
	}
	ECHOJob = Job{
		"The client wishes for the server to echo something back to them.",
		"ECHO",
		func(wi WorkerInfo, ui UserInfo, s []string) error {
			logger.Println("[ ECHO Job ] Echoing back to the user", len(s), s)
			return ui.SendMessage(strings.Join(s, "\n"))
		},
		[]string{},
	}
)
