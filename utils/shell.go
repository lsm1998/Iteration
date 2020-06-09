/*
* 作者：刘时明
* 时间：2020/6/9-0:19
* 作用：
 */
package utils

import (
	"bytes"
	"os/exec"
)

func RunCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name)
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	go func() {
		for _, v := range args {
			in.WriteString(v + "\n")
		}
	}()
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
