// 重命名文件 & 移动文件
package main

import (
	"fmt"
	"os"
)

func main() {
	originalPath := "demo.txt"
	newPath := "demo.md"
	if err := os.Rename(originalPath, newPath); err == nil {
		fmt.Println("文件移动/重命名成功")
	} else {
		fmt.Println("文件移动/重命名失败")
	}
}
