package errx

import (
	"fmt"
	"time"
)

/**
延迟退出，方便查看错误信息
*/
func ErrorAndExit(err string) {
	_ = fmt.Errorf(err + "\n")
	time.Sleep(2 * time.Second)
	panic(err)
}
