// 打开/关闭文件 & 检查否具有读写权限
package main

import (
	"fmt"
	"os"
)

func main() {
	if file, err := os.Open("demo.txt"); err == nil {
		fmt.Printf("打开成功: %s\n", file.Name()) // 只读方式打开
		file.Close()
	} else {
		fmt.Printf("打开失败: %v %+v\n", err, file)
	}

	/*
		// 打开文件, 可读可写, 没有, 则创建新文件
		// 第一个参数: 文件名
		// 第二个参数: 操作文件的方式, eg: 当参数为O_WRONLY时，出现错误，说明没有写的权限
		// 最后一个参数: 要创建的文件权限
		if file, err := os.OpenFile("demo.txt", os.O_RDWR|os.O_CREATE, 0666); err == nil {
			fmt.Printf("打开成功: %s\n", file.Name())
			file.Close()
		} else {
			fmt.Printf("打开失败: %v\n", err)
		}
	*/
}
