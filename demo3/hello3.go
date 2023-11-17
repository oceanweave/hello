// demo3 是包名
package demo3

import "fmt"

// 解法3：因为该文件夹内没有 go.mod，所以可以认为该 demo3 go 包 和 main 包在同一个项目中，因此利用项目名+包名 即可引用
func Hello3() {
	fmt.Println("this is Hello-3 from demo3 project")
}
