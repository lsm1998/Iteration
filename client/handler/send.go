package handler

import (
	"fmt"
	"iteration/client/config"
	"iteration/common"
	"iteration/utils"
	"net"
	"os"
)

func sendShell(cmd string, conn net.Conn) {
	start := len("shell ")
	cmd = cmd[start:]
	temp := []byte(cmd)
	arr := [common.MAX_DATE_LEN]byte{}
	copy(arr[0:len(temp)], temp)
	msg := &common.DataMsg{Cmd: common.CMD_SHELL, Len: int32(len(cmd)), Total: 1, Seq: 1, Data: arr}
	_, err := conn.Write(common.ObjToByte(msg))
	if err != nil {
		_ = fmt.Errorf("发送失败")
	}
}

func sendFile(conn net.Conn) {
	file, err := os.Open(config.C.JarName)
	if file == nil || err != nil {
		panic("找不到文件" + err.Error())
	}
	defer file.Close()
	size := utils.GetFileSize(config.C.JarName)
	var count int32 = 1
	var total int32
	BlockWait()
	for {
		total = int32(size / common.MAX_DATE_LEN)
		if size%common.MAX_DATE_LEN != 0 {
			total++
		}
		msg := &common.DataMsg{
			Cmd:   common.CMD_FILE,
			Seq:   count,
			Total: total,
		}
		data := [common.MAX_DATE_LEN]byte{}
		n, err := file.Read(data[:])
		if err != nil {
			panic(err)
		}
		msg.Len = int32(n)
		msg.Data = data
		_, err = conn.Write(common.ObjToByte(msg))
		if err != nil {
			panic(err)
		}
		if count == total {
			break
		}
		count++
	}
}

// 发送包至服务器
func transferFile() {
	utils.Transfer(&config.C)
}
