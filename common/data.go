/*
* 作者：刘时明
* 时间：2020/6/8-23:20
* 作用：
 */
package common

import (
	"bytes"
	"encoding/binary"
	"runtime"
)

var (
	// CMD命令，win=cmd.exe，linux=/bin/sh
	CMD_NAME = "cmd.exe"
)

const (
	// 执行Shell脚本
	CMD_SHELL = iota
	// 文件传输
	CMD_FILE
)

const (
	// 数据数组长度
	MAX_DATE_LEN = 1024 * 1024
	// 每个数据包大小
	MSG_LEN = MAX_DATE_LEN + 4*4
	// 针对的Jar包
	JAR_NAME = "hello.exe"
	// 服务器IP地址
	SERVER_ADDR = "118.24.239.74"
	// 服务器端口
	SERVER_PORT = 8848
)

func init() {
	if runtime.GOOS == "linux" {
		CMD_NAME = "/bin/sh"
	}
}

type DataMsg struct {
	// 命令号
	Cmd int32
	// 包长度
	Len int32
	// 包序号
	Seq int32
	// 包数量
	Total int32
	// 数据
	Data [MAX_DATE_LEN]byte
}

func ByteToObj(data *[]byte) *DataMsg {
	msg := &DataMsg{}
	buffer := bytes.NewReader(*data)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Cmd)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Len)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Seq)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Total)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Data)
	return msg
}

func ObjToByte(data *DataMsg) []byte {
	buffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(buffer, binary.LittleEndian, data.Cmd)
	_ = binary.Write(buffer, binary.LittleEndian, data.Len)
	_ = binary.Write(buffer, binary.LittleEndian, data.Seq)
	_ = binary.Write(buffer, binary.LittleEndian, data.Total)
	_ = binary.Write(buffer, binary.LittleEndian, data.Data)
	return buffer.Bytes()
}
