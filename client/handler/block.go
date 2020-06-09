package handler

var blockChannel = make(chan bool, 1)

func BlockWait() {
	<-blockChannel
}

func BlockNotify() {
	blockChannel <- true
}
