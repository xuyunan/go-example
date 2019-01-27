// 删除文件
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Remove("demo.txt"); err == nil {
		fmt.Println("Remove success")
	} else {
		fmt.Printf("Remove failed: %v\n", err) // 没有文件会删除失败
	}
}
