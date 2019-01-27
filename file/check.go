// 获取文件信息 & 检查文件是否存在
package main

import (
	"fmt"
	"os"
)

func main() {
	if fileInfo, err := os.Stat("demo.txt"); err == nil {
		fmt.Println("文件存在:") // +v输出字段对应的key
		fmt.Println("File name:", fileInfo.Name())
		fmt.Println("Size in bytes:", fileInfo.Size())
		fmt.Println("Permissions:", fileInfo.Mode())
		fmt.Println("Last modified:", fileInfo.ModTime())
		fmt.Println("Is Directory: ", fileInfo.IsDir())
		fmt.Printf("System interface type: %T\n", fileInfo.Sys())
		fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	} else {
		fmt.Printf("文件不存在: %v\n", err)
	}
}
