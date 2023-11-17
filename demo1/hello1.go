package demo1

import "fmt"

// 解法1： 上传github后，go get 获取
func Hello1() {
	fmt.Println("this is Hello-1 from demo1")
	V1_demo1_new_func()
}

func V1_demo1_new_func() {
	fmt.Println("this is v1-new-feature from demo1!")
}
