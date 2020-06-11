package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iteration/client/errx"
)

var C Config

type Config struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	JarName   string `json:"jarName"`
	Path      string `json:"path"`
	RemoteDir string `json:"remoteDir"`
	Ip        string `json:"ip"`
	Port      int32  `json:"port"`
}

func init() {
	fmt.Println("开始读取配置文件")
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		errx.ErrorAndExit(err.Error())
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b, &C)
	if err != nil {
		errx.ErrorAndExit(err.Error())
	}
}
