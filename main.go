package main

import (
	"github.com/oceanweave/hello/demo1"
	"github.com/oceanweave/hello/demo2"
)

func main() {
	// 报错
	// main.go:3:8: no required module provides package github.com/oceanweave/hello/demo1; to add it:
	//	go get github.com/oceanweave/hello/demo1
	demo1.Hello1()
	demo2.Hello2()
}
