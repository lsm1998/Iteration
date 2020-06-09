/*
* 作者：刘时明
* 时间：2020/6/8-22:58
* 作用：
 */
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"iteration/client/handler"
	"iteration/common"
	"net"
	"os"
	"strings"
)

func main() {
	listeners()
}

var send = func(conn net.Conn) {
	termReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(":> ")
		line, err := termReader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		cmd := string(line)
		switch cmd {
		case "upload":
		case "clear":
			for i := 0; i < 50; i++ {
				fmt.Println()
			}
		default:
			if strings.LastIndex(cmd, "shell") >= 0 {
				sendShell(cmd, conn)
			} else {
				fmt.Println("不能识别的命令")
			}
		}
	}
}

func listeners() {
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
				fmt.Println("读到的字节数=", cnt)
				if cnt < 1024 {
					break read
				}
			}
			fmt.Println("收到消息->" + buffer.String())
		}
	}
	handler.ClientStart(send, receive)
}

func sendShell(cmd string, conn net.Conn) {
	temp := []byte(cmd)
	arr := [common.MAX_DATE_LEN]byte{}
	copy(arr[0:len(temp)], temp)
	msg := &common.DataMsg{Cmd: common.CMD_SHELL, Len: int32(len(cmd)), Size: 1, Seq: 1, Data: arr}
	_, err := conn.Write(common.ObjToByte(msg))
	if err != nil {
		_ = fmt.Errorf("发送失败")
	}
}
