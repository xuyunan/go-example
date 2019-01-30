// 目录操作
package main

import (
	"fmt"
	"os"
)

// IsDirExist 目录是否存在
func IsDirExist(path string) bool {
	var info os.FileInfo
	var err error
	if info, err = os.Stat(path); err == nil {
		return info.IsDir()
	}
	return false
}

// MakeDir 创建目录
func MakeDir(path string) (bool, error) {
	if IsDirExist(path) {
		return true, nil
	}
	var err error
	if err = os.MkdirAll(path, 0700); err == nil {
		return true, nil
	}
	return false, err
}

func main() {
	result, err := MakeDir("images")
	fmt.Printf("结果: %t %v\n", result, err)
}
