// 创建文件
package main

import (
	"fmt"
	"os"
)

func main() {
	if file, err := os.Create("demo.txt"); err == nil {
		fmt.Printf("文件创建成功: %s\n", file.Name())
	} else {
		fmt.Printf("文件创建失败: %v\n", err)
	}
}
