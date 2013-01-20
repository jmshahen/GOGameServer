package gameserver

import (
	"fmt"
	"strings"
)

type Job struct {
	Description string
	Header      string
	Run         func(WorkerInfo, UserInfo, string) error
	Data        string
}

func (gs GameServer) Decode(s string) (Job, error) {
	var job Job

	s = strings.TrimRight(s, string(gs.Terminator))
	z := strings.SplitN(s, string(gs.Separator), 2)

	//TODO implement services and a job lookup table
	if z[0] == "ECHO" {
		job = ECHOJob
		job.Data = z[1]
		return job, nil
	} else if z[0] == "QUIT" {
		job = QUITJob
		job.Data = z[1]
		return job, nil
	}

	return job, NewServerError(fmt.Sprintf("Could not find Job %s", z[0]))
}

// Common Jobs
var (
	QUITJob = Job{
		"When the client is about to leave they will send this to logoff.",
		"QUIT",
		func(wi WorkerInfo, ui UserInfo, s string) error {
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
		"",
	}
	ECHOJob = Job{
		"The client wishes for the server to echo something back to them.",
		"ECHO",
		func(wi WorkerInfo, ui UserInfo, s string) error {
			logger.Println("[ ECHO Job ] Echoing back to the user", len(s), s)
			return ui.SendMessage(s)
		},
		"",
	}
)
