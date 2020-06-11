/*
* 作者：刘时明
* 时间：2020/6/10-0:57
* 作用：
 */
package common

import (
	"runtime"
)

var (
	// CMD命令，win=cmd.exe，linux=/bin/sh
	CMD_NAME = "cmd.exe"
	// 操作系统类型
	OS_NAME = "windows"
)

const (
	// 执行Shell脚本
	CMD_SHELL = iota
	// 文件传输
	CMD_FILE
)

const (
	// 数据数组长度
	MAX_DATE_LEN = 1024 * 10
	// 每个数据包大小
	MSG_LEN = MAX_DATE_LEN + 4*4
)

func init() {
	if runtime.GOOS == "linux" {
		CMD_NAME = "/bin/sh"
		OS_NAME = "linux"
	}
}
