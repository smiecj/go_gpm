package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 创建一个 trace 文件
	file, err := os.Create("trace.out")
	if nil != err {
		panic(err)
	}

	defer file.Close()

	// 开始 跟踪
	err = trace.Start(file)
	if nil != err {
		panic(err)
	}

	fmt.Println("Hello GPM")

	// 结束跟踪
	trace.Stop()
}
