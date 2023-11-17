module github.com/oceanweave/hello

go 1.19

// 由于第一次 commit，已经将 demo1 包上传到了 github
// 所以 go get -u github.com/oceanweave/hello/demo1 便可以下载，同时 main.go 函数也可以正常执行
// v0.0.0-20231117070717-83bd4d1bb2d5 由版本号、commit时间和commit的hash值组成
require github.com/oceanweave/hello/demo1 v0.0.0-20231117070717-83bd4d1bb2d5
