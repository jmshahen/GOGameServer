package gameserver

import (
// "fmt"
)

func (wi WorkerInfo) worker() {
	defer wi.sayGoodbye()
	var quit = false
	for {
		wi.logMsg("Waitng for work")
		<-wi.doWork
		wi.logMsg("There is work to be done!")

		select {
		case <-wi.quit:
			wi.logMsg("Quitting")
			quit = true
			break
		case user := <-wi.addUser:
			wi.logMsg("Added User", user.Id)
			wi.users[user.Id] = user
		default:
			for _, user := range wi.users {
				select {
				case job := <-user.ch:
					wi.logMsg("Processing job from user", user.Id)
					wi.logMsg("job:", job)
					job.Run(wi, user, job.Data)
				default:
					//wi.logMsg("No Work from User", user.Id)
				}
			}
		}
		if quit {
			break
		}
	}
	wi.logMsg("Is Shutdown")
}

func (wi WorkerInfo) sayGoodbye() {
	wi.logMsg("Saying goodbye to all my users")
	for _, user := range wi.users {
		user.SendMessage("Goodbye")
	}
}

func (wi WorkerInfo) logMsg(s ...interface{}) {
	logger.Println("[ Worker", wi.Id, "]", s)
}
