package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

var send = func(conn net.Conn) {
	termReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(":>")
		line, err := termReader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		cmd := string(line)
		cmd = cmd[0 : len(cmd)-1]
		switch cmd {
		case "exit":
			_ = conn.Close()
			os.Exit(0)
		case "clear":
			for i := 0; i < 50; i++ {
				fmt.Println()
			}
		default:
			if strings.Index(cmd, "upload") >= 0 {
				sendFile(cmd, conn)
			} else if strings.Index(cmd, "shell") >= 0 {
				sendShell(cmd, conn)
			} else {
				fmt.Println("不能识别的命令")
			}
		}
	}
}

func Listeners() {
	var receive = func(conn net.Conn) {
	the:
		for {
			var buffer bytes.Buffer
		read:
			for {
				// 每次读取2个字节
				temp := make([]byte, 1024)
				cnt, err := conn.Read(temp)
				if err != nil {
					println("服务器异常")
					_ = conn.Close()
					break the
				}
				buffer.Write(temp[0:cnt])
				if cnt < 1024 {
					break read
				}
			}
			fmt.Println("收到消息->" + buffer.String())
			BlockNotify()
		}
	}
	ClientStart(send, receive)
}
