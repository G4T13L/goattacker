package attacks

import (
	"fmt"
	"log"
)

//workAuth goroutine to send an Auth attack
func workAuth(job login, url, post, proxy string) {
	// fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	html, code, _ := SendAuth(url, post, job.user, job.pass, proxy)
	if code == 200 {
		fmt.Println(green("\t", code, "\t", len(html), "\t", url, "\t", job.user, "\t", job.pass))
	} else {
		fmt.Println(red("\t", code, "\t", len(html), "\t", url, "\t", job.user, "\t", job.pass))
	}
}

//AuthAttack Auth attack for http or https
func AuthAttack(url, post, userFile, passFile, proxy string, nWorkers int) {
	if nWorkers == 0 {
		nWorkers = 9
	}

	log.Println("[0] Starting jobs")
	log.Println(userFile, passFile, url, post, proxy, nWorkers)

	outchan := make(chan login, nWorkers)
	go Read2Files(userFile, passFile, outchan)
	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
	fmt.Println("\tCODE\tlen(html)\tURL\tUSER\tPASS")
	for {
		select {
		case job, ok := <-outchan:
			if !ok {
				wg.Wait()
				fmt.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			go func(job login) {
				workAuth(job, url, post, proxy)
				sem <- 1
			}(job)
		}
	}
}

//workFile goroutine to send an file search attack
func workFile(job login, url, post, proxy string) {
	// fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	html, code, _ := SendAuth(url, post, job.user, job.pass, proxy)
	if code == 200 {
		fmt.Println(green("\t", code, "\t", len(html), "\t", url, "\t", job.user, "\t", job.pass))
	} else {
		fmt.Println(red("\t", code, "\t", len(html), "\t", url, "\t", job.user, "\t", job.pass))
	}
}

//FileAttack Auth attack for http or https
func FileAttack(url, post, userFile, passFile, proxy string, nWorkers int) {
	if nWorkers == 0 {
		nWorkers = 9
	}

	log.Println("[0] Starting jobs")
	log.Println(userFile, passFile, url, post, proxy, nWorkers)

	outchan := make(chan login, nWorkers)
	go Read2Files(userFile, passFile, outchan)
	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
	fmt.Println("\tCODE\tlen(html)\tURL\tUSER\tPASS")
	for {
		select {
		case job, ok := <-outchan:
			if !ok {
				wg.Wait()
				fmt.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			go func(job login) {
				workAuth(job, url, post, proxy)
				sem <- 1
			}(job)
		}
	}
}
