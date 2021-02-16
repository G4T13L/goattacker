package attacks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var magenta = color.New(color.FgHiMagenta).SprintFunc()

type login struct {
	user string
	pass string
	pos  string
}

var wg sync.WaitGroup

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Read2Files read 2 files to itter between them
func Read2Files(userFile, passFile string, loginChan chan login) {

	f1, err := os.Open(userFile)
	check(err)
	f2, err := os.Open(passFile)
	check(err)

	defer func() {
		f1.Close()
		f2.Close()
		close(loginChan)
	}()

	ureader := bufio.NewScanner(f1)
	preader := bufio.NewScanner(f2)

	var userList []string
	for i := 0; preader.Scan(); i++ {
		if i == 0 {
			for j := 0; ureader.Scan(); j++ {
				userList = append(userList, ureader.Text())
				loginChan <- login{user: userList[len(userList)-1], pass: preader.Text(), pos: strconv.Itoa(i) + "-" + strconv.Itoa(j)}
			}

			if err := ureader.Err(); err != nil {
				check(err)
			}
		} else {
			for j, u := range userList {
				loginChan <- login{user: u, pass: preader.Text(), pos: strconv.Itoa(i) + "-" + strconv.Itoa(j)}
			}
		}
	}
	if err := preader.Err(); err != nil {
		check(err)
	}

}

//ConstWithFile read one file and other argument that it could be user or password and it is specified in posArg
func ConstWithFile(argument, file string, posArg int, loginChan chan login) {
	f1, err := os.Open(file)
	check(err)

	defer func() {
		f1.Close()
		close(loginChan)
	}()

	freader := bufio.NewScanner(f1)
	for i := 0; freader.Scan(); i++ {
		if posArg == 0 {
			fmt.Println(freader.Text())
			loginChan <- login{user: argument, pass: freader.Text(), pos: strconv.Itoa(i)}
		} else if posArg == 1 {
			loginChan <- login{user: freader.Text(), pass: argument, pos: strconv.Itoa(i)}
		}
		if err := freader.Err(); err != nil {
			check(err)
		}
	}
	if err := freader.Err(); err != nil {
		check(err)
	}
}
