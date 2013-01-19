package gameserver

import (
	"fmt"
)

func (wi WorkerInfo) worker() {
	fmt.Println("Worker", wi.Id, "started")
	var quit = false
	for {
		fmt.Println("[Worker", wi.Id, "]Waitng for work")
		<-wi.doWork
		fmt.Println("There is work to be done!")

		select {
		case <-wi.quit:
			fmt.Println("worker", wi.Id, "quitting")
			wi.sayGoodbye()
			quit = true
			break
		case user := <-wi.addUser:
			fmt.Println("Worker", wi.Id, "Added User", user.Id)
			wi.users = append(wi.users, user)
		default:
			for _, user := range wi.users {
				select {
				case job := <-user.ch:
					fmt.Println("worker", wi.Id, "processing job from user", user.Id)
					fmt.Println("job:", job)
					user.SendMessage(job.msg)
				default:
					fmt.Println("No Work from User", user.Id)
				}
			}
		}
		if quit {
			break
		}
	}
	fmt.Println("Worker", wi.Id, "Shutting down")
}

func (wi WorkerInfo) sayGoodbye() {
	for _, user := range wi.users {
		user.SendMessage("Goodbye")
	}
}
