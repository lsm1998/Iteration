/*
* 作者：刘时明
* 时间：2020/6/8-23:17
* 作用：
 */
package handler

import (
	"fmt"
	"iteration/client/config"
	"iteration/client/errx"
	"net"
)

func ClientStart(send func(net.Conn), receive func(net.Conn)) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.C.Ip, config.C.Port))
	if err != nil {
		println("客户端建立连接失败")
		errx.ErrorAndExit(err.Error())
	}
	go getHandler(receive, conn)
	clientHandler(send, conn)
}

func clientHandler(send func(net.Conn), conn net.Conn) {
	send(conn)
}

func getHandler(receive func(net.Conn), conn net.Conn) {
	receive(conn)
}
