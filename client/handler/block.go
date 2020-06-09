/*
* 作者：刘时明
* 时间：2020/6/10-0:58
* 作用：模拟阻塞
 */
package handler

var blockChannel = make(chan bool, 1)

func BlockWait() {
	if len(blockChannel) == 0 {
		blockChannel <- true
	}
}

func BlockNotify() {
	if len(blockChannel) > 0 {
		<-blockChannel
	}
}
