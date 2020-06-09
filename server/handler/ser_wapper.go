/*
* 作者：刘时明
* 时间：2020/6/8-23:03
* 作用：
 */
package handler

import (
	"fmt"
	"net"
)

var dataList = make([]interface{}, 0, 10)

// 服务端启动
func ServerStart(handler func(net.Conn, []interface{})) {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", 8848))
	if err != nil {
		fmt.Println("TCP服务端启动失败")
		panic(err)
	}
	fmt.Println("TCP服务端启动完毕...")
	for {
		conn, err := server.Accept()
		fmt.Println("一个客户端连入...")
		if err != nil {
			fmt.Println("连接出错")
			continue
		}
		go serverHandler(handler, conn, dataList)
	}
}

func serverHandler(handler func(net.Conn, []interface{}), conn net.Conn, list []interface{}) {
	handler(conn, list)
}
