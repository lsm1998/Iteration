/*
* 作者：刘时明
* 时间：2020/6/8-23:17
* 作用：
 */
package handler

import "net"

func ClientStart(send func(net.Conn), receive func(net.Conn)) {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		println("客户端建立连接失败")
		panic(err)
		return
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
