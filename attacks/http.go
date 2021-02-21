package attacks

import (
	"fmt"
	"log"
	"strings"
)

//workAuth goroutine to send an Auth attack
func workAuth(job login, url, post, proxy string, redirect bool) {
	// fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	html, code := SendAuth(url, post, job.user, job.pass, proxy, redirect)
	if code == 200 {
		log.Println(green("\t ", code, " \t", len(html), "\t", url, "\t ", job.user, ":", job.pass))
		// log.Println(green("\t ", code, " \t", len(html), "\t\t", len(strings.Split(html, " ")), " \t ", url))
	} else {
		log.Println(red("\t ", code, " \t", len(html), "\t", url, "\t ", job.user, ":", job.pass))
	}
}

//AuthAttack Auth attack for http or https
func AuthAttack(url, post, userFile, passFile, proxy string, nWorkers int, redirect bool) {
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
	fmt.Println(yellow("\t\tLOG\tCODE len(html)\t\tURL\t\t\tUSER:PASS"))
	for {
		select {
		case job, ok := <-outchan:
			if !ok {
				wg.Wait()
				log.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			go func(job login) {
				workAuth(job, url, post, proxy, redirect)
				sem <- 1
			}(job)
		}
	}
}

//workFile goroutine to send an file search attack
func workFile(job login, url, word, post, proxy string, redirect bool) {
	defer wg.Done()

	// fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	file := job.user
	ext := job.pass
	html, code, url := FileTry(url, word, file, ext, post, proxy, redirect)
	if code == 200 {
		log.Println(green("\t ", code, " \t", len(html), "\t\t", len(strings.Split(html, " ")), " \t ", url))
	} else if code == 404 {
		log.Println(red("\t ", code, " \t", len(html), "\t\t", len(strings.Split(html, " ")), " \t ", url))
	} else {
		log.Println(magenta("\t ", code, " \t", len(html), "\t\t", len(strings.Split(html, " ")), " \t ", url))
	}
}

//FileAttack File attack for http or https
func FileAttack(url, post, word, file, ext, proxy string, nWorkers int, redirect bool) {
	if nWorkers == 0 {
		nWorkers = 9
	}

	log.Println("[0] Starting jobs")
	log.Println(url, post, word, file, ext, proxy, nWorkers)

	outchan := make(chan login, nWorkers)

	// go Read2Files(userFile, passFile, outchan)
	if ext == "" {
		go ConstWithFile("", file, 1, outchan)
	} else {
		go Read2Files(file, ext, outchan)
	}

	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
	fmt.Println(yellow("\t\tLOG\tCODE\tlen(html)\tw\t\tURL"))
	for {
		select {
		case job, ok := <-outchan:
			if !ok {
				wg.Wait()
				log.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			go func(job login) {
				workFile(job, url, word, post, proxy, redirect)
				sem <- 1
			}(job)
		}
	}
}

//workLogin goroutine to send an form Login attack
func workLogin(job login, url, word, post, proxy string, redirect, show bool) {
	// fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	post = strings.Replace(post, "$$USER$$", job.user, 1)
	post = strings.Replace(post, "$$PASS$$", job.pass, 1)

	html, code, tf := FormLogin(url, post, word, proxy, redirect)
	if show {
		fmt.Println(html)
	}
	if tf == true {
		log.Println(green("\t ", code, " \t", len(html), "\t [OK]", "\t", url, "\t ", post))
		// log.Println(green("\t ", code, " \t", len(html), "\t\t", len(strings.Split(html, " ")), " \t ", url))
	} else {
		log.Println(red("\t ", code, " \t", len(html), "\t [x]", "\t", url, "\t ", post))
		// log.Println(red("\t| ", code, " |\t", len(html), "\t| ", url, "\t| ", job.user, ":", job.pass))
	}
}

//FormAttack Form login attack for http or https
func FormAttack(url, post, userFile, passFile, word, proxy string, nWorkers int, redirect, show bool) {
	if nWorkers == 0 {
		nWorkers = 9
	}

	log.Println("[0] Starting jobs")
	log.Println(url, post, userFile, passFile, word, proxy, nWorkers)

	outchan := make(chan login, nWorkers)

	go Read2Files(userFile, passFile, outchan)

	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
	fmt.Println(yellow("\t\tLOG\tCODE\tlen\tinfo\tURL\t\t\t\tPOST DATA"))
	for {
		select {
		case job, ok := <-outchan:
			if !ok {
				wg.Wait()
				log.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			go func(job login) {
				workLogin(job, url, word, post, proxy, redirect, show)
				sem <- 1
			}(job)
		}
	}
}
