/*
* 作者：刘时明
* 时间：2020/6/8-23:45
* 作用：
 */
package utils

import (
	"os"
	"path/filepath"
)

func GetFileSize(filePath string) int64 {
	var result int64
	filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
