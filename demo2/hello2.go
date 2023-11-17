package demo2

import "fmt"

//  2. 解法2： 版本设置  v0.0.0，伪版本，一般用于未发布版本的测试，理解为不会去访问 github
//     那么如何找到正确版本，通过 replace 为使用者指引到该 go 包的正确位置
func Hello2() {
	fmt.Println("this is Hello-2 from demo2 project")
	V1_demo2_new_func()
}

func V1_demo2_new_func() {
	fmt.Println("this is v1-new-feature from demo2!")
}
