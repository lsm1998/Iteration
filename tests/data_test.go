/*
* 作者：刘时明
* 时间：2020/6/8-23:25
* 作用：
 */
package tests

import (
	"fmt"
	"iteration/common"
	"iteration/utils"
	"testing"
)

func TestData(t *testing.T) {
	msg := &common.DataMsg{
		Cmd:  1,
		Len:  5,
		Data: [common.MAX_DATE_LEN]byte{1, 2, 3, 4, 5},
	}
	toByte := common.ObjToByte(msg)

	fmt.Println(len(toByte))

	obj := common.ByteToObj(&toByte)
	fmt.Println(obj.Cmd)
	fmt.Println(obj.Len)
}

func TestFile(t *testing.T) {
	fmt.Println(utils.GetFileSize("1.jpg"))
}
