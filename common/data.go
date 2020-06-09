/*
* 作者：刘时明
* 时间：2020/6/8-23:20
* 作用：
 */
package common

import (
	"bytes"
	"encoding/binary"
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
)

type DataMsg struct {
	// 命令号
	Cmd int32
	// 包长度
	Len int32
	// 包序号
	Seq int32
	// 包数量
	Size int32
	// 数据
	Data [MAX_DATE_LEN]byte
}

func ByteToObj(data *[]byte) *DataMsg {
	msg := &DataMsg{}
	buffer := bytes.NewReader(*data)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Cmd)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Len)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Seq)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Size)
	_ = binary.Read(buffer, binary.LittleEndian, &msg.Data)
	return msg
}

func ObjToByte(data *DataMsg) []byte {
	buffer := bytes.NewBuffer([]byte{})
	_ = binary.Write(buffer, binary.LittleEndian, data.Cmd)
	_ = binary.Write(buffer, binary.LittleEndian, data.Len)
	_ = binary.Write(buffer, binary.LittleEndian, data.Seq)
	_ = binary.Write(buffer, binary.LittleEndian, data.Size)
	_ = binary.Write(buffer, binary.LittleEndian, data.Data)
	return buffer.Bytes()
}
