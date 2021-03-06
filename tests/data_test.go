/*
* 作者：刘时明
* 时间：2020/6/8-23:25
* 作用：
 */
package tests

import (
	"fmt"
	"iteration/client/config"
	"iteration/common"
	"iteration/utils"
	"testing"
	"time"
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
	path := "C:/Users\\Administrator/GolandProjects/Iteration/tests/data_test.go"
	fmt.Println(utils.GetFileName(&path))
	fmt.Println(utils.GetFileName(&path))
	// utils.MakeFile("www/ast/1.txt")
}

func TestCmd(t *testing.T) {
	fmt.Println(utils.RunCmd("cmd.exe", "dir"))
}

func TestTransfer(t *testing.T) {
	start := time.Now().Unix()
	utils.Transfer(&config.C)
	end := time.Now().Unix()
	fmt.Println(end - start)
}
