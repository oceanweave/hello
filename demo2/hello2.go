package demo2

import "fmt"

//  2. 解法2： 版本设置  v0.0.0，伪版本，一般用于未发布版本的测试，理解为不会去访问 github
//     那么如何找到正确版本，通过 replace 为使用者指引到该 go 包的正确位置
func Hello2() {
	fmt.Println("this is Hello-2 from demo2 project")
	V1_demo2_new_func()
}

// 此时 go.main 函数后的输出
// this is Hello-1 from demo1
// this is Hello-2 from demo2 project
// this is v1-new-feature from demo2!
// this is Hello-3 from demo3 project
// 可以发现，没有更改 go.mod 文件和 go get -u，但是 demo2 已经是新版本呢，而 demo1 还是旧版本输出
// 因为在 main 项目中的引用是 v0.0.0 和 replace，直接指向了本地的此 go 包，所以不需要拉取，就可以执行新版本程序
func V1_demo2_new_func() {
	fmt.Println("this is v1-new-feature from demo2!")
}
