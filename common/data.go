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
