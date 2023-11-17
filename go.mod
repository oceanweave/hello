module github.com/oceanweave/hello

go 1.21.4

// 1. 解法1： 上传github后，go get 获取
// 由于第一次 commit，已经将 demo1 包上传到了 github
// 所以 go get -u github.com/oceanweave/hello/demo1 便可以下载，同时 main.go 函数也可以正常执行
// v0.0.0-20231117070717-83bd4d1bb2d5 由版本号、commit时间和commit的hash值组成
require github.com/oceanweave/hello/demo1 v0.0.0-20231117070717-83bd4d1bb2d5

// 2. 解法2： 版本设置  v0.0.0，伪版本，一般用于未发布版本的测试，理解为不会去访问 github
//           那么如何找到正确版本，通过 replace 为该 go 包指引到正确的位置
require github.com/oceanweave/hello/demo2 v0.0.0

replace github.com/oceanweave/hello/demo2 => ./demo2
