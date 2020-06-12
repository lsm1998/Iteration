package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var resultChannel = make(chan string, 1)

var send = func(conn net.Conn) {
	termReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(":>")
		line, err := termReader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		cmd := string(line)
		cmd = strings.ReplaceAll(cmd, "\r\n", "")
		cmd = strings.ReplaceAll(cmd, "\n", "")
		switch cmd {
		case "?":
			printHelp()
		case "pag":
			transferFile()
			sendShell("shell sh build.sh", conn)
			printResult()
		case "exit":
			_ = conn.Close()
			os.Exit(0)
		case "clear":
			for i := 0; i < 50; i++ {
				fmt.Println()
			}
		default:
			if strings.Index(cmd, "upload") >= 0 {
				sendFile(conn)
			} else if strings.Index(cmd, "shell") >= 0 {
				sendShell(cmd, conn)
				printResult()
			} else {
				fmt.Printf("不能识别的命令，你输入的是[%s]，可以输入'?'查看手册\n", cmd)
			}
		}
	}
}

func printHelp() {
	fmt.Println("pag 发包指令")
	fmt.Println("shell [cmd] 执行命令")
	fmt.Println("exit 退出程序")
	fmt.Println("clear 清空命令行")
	fmt.Println("upload [fileName] 上传文件")
}

func printResult() {
	select {
	case str := <-resultChannel:
		fmt.Println("收到回复 ->\n", str)
	case <-time.After(10 * time.Second):
		fmt.Println("回复超时了！")
	}
}

func Listeners() {
	var receive = func(conn net.Conn) {
	the:
		for {
			var buffer bytes.Buffer
		read:
			for {
				// 每次读取1K
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
			result := buffer.String()
			if strings.Index(result, "notify") >= 0 {
				resultChannel <- result[6:]
			} else {
				resultChannel <- result
			}
		}
	}
	ClientStart(send, receive)
}
