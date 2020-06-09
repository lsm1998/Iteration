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
	"iteration/utils"
	"net"
	"os"
)

func main() {
	handler.ServerStart(serHandler)
}

var serHandler = func(conn net.Conn, list []interface{}) {
	for {
		first := make([]byte, common.MSG_LEN)
		_, err := conn.Read(first)
		if err != nil {
			fmt.Println("一个客户端退出")
			_ = conn.Close()
			return
		}
		msg := common.ByteToObj(&first)
		fmt.Println("解析到一个msg,cmd=", msg.Cmd, ",seq=", msg.Seq)
		switch msg.Cmd {
		case common.CMD_SHELL:
			shellStrategy(msg, conn)
		case common.CMD_FILE:
			fileStrategy(msg, conn)
		}
	}
}

func fileStrategy(msg *common.DataMsg, conn net.Conn) {
	fmt.Println("收到一个包,seq=", msg.Seq, ",total=", msg.Total, ",Len=", msg.Len)
	// 是否第一个包
	if msg.Seq == 1 {
		_ = utils.MakeFile("copy" + common.JAR_NAME)
	}
	file, err := os.OpenFile("copy"+common.JAR_NAME, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if file == nil || err != nil {
		panic(err)
	}
	_, err = file.Seek(int64((msg.Seq-1)*common.MAX_DATE_LEN), 0)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(msg.Data[0:msg.Len])
	if err != nil {
		panic(err)
	}
	// 是否最后一个包
	if msg.Seq == msg.Total {
		_, _ = conn.Write([]byte("notify最后一个包，文件传输完毕"))
	} else {
		_, _ = conn.Write([]byte(fmt.Sprintf("已完成第%d个包,共%d个", msg.Seq, msg.Total)))
	}
}

func shellStrategy(msg *common.DataMsg, conn net.Conn) {
	shell := string(msg.Data[0:msg.Len])
	result, err := utils.RunCmd(common.CMD_NAME, shell)
	if err != nil {
		_, _ = conn.Write([]byte("执行错误：" + err.Error()))
	} else {
		_, _ = conn.Write([]byte(result))
	}
}
