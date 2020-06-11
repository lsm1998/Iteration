/*
* 作者：刘时明
* 时间：2020/6/8-23:45
* 作用：
 */
package utils

import (
	"os"
	"strings"
)

/**
获取文件大小
*/
func GetFileSize(filePath string) int64 {
	stat, err := os.Stat(filePath)
	if err != nil {
		return -1
	}
	return stat.Size()
}

/**
判断文件是否存在
*/
func FileExit(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

/**
创建文件，如果存在则删除
*/
func MakeFile(filePath string) error {
	// 是否带有目录
	if arr := DirPathArr(&filePath); len(arr) > 1 {
		MakeDir(arr[0])
		//MakeFile(arr[0] + "/" + arr[1])
	}
	if FileExit(filePath) {
		_ = os.RemoveAll(filePath)
	}
	file, err := os.Create(filePath)
	defer file.Close()
	return err
}

/**
创建目录
*/
func MakeDir(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModePerm)
}

/**
返回目录+文件
*/
func DirPathArr(path *string) []string {
	index1 := strings.LastIndex(*path, "\\")
	index2 := strings.LastIndex(*path, "/")
	if index1 == -1 && index2 == -1 {
		return []string{*path}
	}
	index := index1
	if index2 > index1 {
		index = index2
	}
	return []string{(*path)[0:index], (*path)[index+1 : len(*path)]}
}

/**
获取文件名称
*/
func GetFileName(filePath *string) string {
	arr := DirPathArr(filePath)
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[1]
}
