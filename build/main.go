package main

import (
	"iteration/utils"
	"time"
)

func main() {
	_, _ = utils.RunCmd("cmd.exe", "git add 1.txt")
	_, _ = utils.RunCmd("cmd.exe", "git commit -m 'ok'")
	_, _ = utils.RunCmd("cmd.exe", "git push  -u origin master")
	time.Sleep(time.Second * 10)
}
