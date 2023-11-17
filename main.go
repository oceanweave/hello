package main

import (
	"github.com/oceanweave/hello/demo1"
	"github.com/oceanweave/hello/demo2"
	"github.com/oceanweave/hello/demo3"
)

func main() {
	// 报错
	// main.go:3:8: no required module provides package github.com/oceanweave/hello/demo1; to add it:
	//	go get github.com/oceanweave/hello/demo1

	// 解法1：子文件夹 demo1 创建了 go.mod，上传 github，之后 go get
	demo1.Hello1()
	// 解法2：子文件夹 demo2 创建了 go.mod，利用 v0.0.0 和 replace 操作，在 go.mod 文件指明 go 包位置
	demo2.Hello2()
	// 解法3：子文件夹 demo3 不创建 go.mod (视为和 main 函数同一个项目），利用 项目名（mod 名 github.com/oceanweave/hello） + 包名（demo3) ，进行文件
	demo3.Hello3()
}
