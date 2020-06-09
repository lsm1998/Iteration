/*
* 作者：刘时明
* 时间：2020/6/8-22:58
* 作用：
 */
package main

import (
	"fmt"
	"iteration/common"
	"iteration/server/handler"
	"net"
)

func main() {
	handler.ServerStart(serHandler)
}

var serHandler = func(conn net.Conn, list []interface{}) {
	first := make([]byte, common.MSG_LEN)
	_, err := conn.Read(first)
	if err != nil {
		fmt.Println("一个客户端退出")
		_ = conn.Close()
		return
	}
	msg := common.ByteToObj(&first)
	switch msg.Cmd {
	case common.CMD_SHELL:
		shellStrategy(msg, conn)
	case common.CMD_FILE:
	}
}

func shellStrategy(msg *common.DataMsg, conn net.Conn) {
	shell := string(msg.Data[0:msg.Len])
	fmt.Println(msg.Cmd)
	fmt.Println(msg.Len)
	fmt.Println("执行脚本=", shell)
	conn.Write([]byte("ok"))
}
