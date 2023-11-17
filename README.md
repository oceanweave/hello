# ReadMe

- 参考自
  - https://www.liwenzhou.com/posts/Go/package/

- 发布版本，第一步，打 tag 
  - `git tag -a 版本号 commit-id -m 附加信息`
  - ` git tag -a v1.0 de7cc24 -m "Release v1.0" `
- 发布版本，第二步，发布到 github 上
  - `git push origin v1.0.1`
- 若提交到 main，可以省略分支名称
  - `git push origin v1.0`
- 开发建议
  - 因此建议开发项目时，只创建一个 go.mod, 
  - main 函数对于子目录中的包的引用，采用【项目名（go.mod 名称，也就是 repo 地址，github.com/{username}/{repo-name}】+【包名(一般来说就是所在目录名称）】形式引用，
  - 也就是之前说的解法3，详情查看 demo3 文件夹和 main.go 对其的引用

- 发布版本，版本形式要求 https://go.dev/ref/mod#versions
  - 形式建议 v1.2.3
  - 1 主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
  - 2 次版本号：发布了功能性更新时递增。
  - 3 修订号：发布了bug修复类更新时递增。

- 发布版本 
  - git tag -a v1.0.1 0c74f12d -m "Release v1.0.1"
  - git push origin v1.0.1

- go get 有三种拉取方法
  - 版本号形式拉取： 		go get github.com/oceanweave/hello/demo1@v1.0.1
  - 分支形式拉取：  		go get github.com/oceanweave/hello/demo1@git-tag-test  会拉取分支最新的 commit
  - commit-id形式拉取：	go get github.com/oceanweave/hello/demo1@0c74f12d

一般来说，开发建议，子目录不创建 go.mod
因此来说，本项目有些反常规，所以导致  --> 版本号形式拉取失败
支持【分支形式】和【commit-id】形式 拉取 go 包

因此可以进行如下测试
首先拉取代码，go run main.go 执行，会发现使用的是老版本 demo1 go 包
输出如下：
this is Hello-1 from demo1
this is Hello-2 from demo2 project
this is v1-new-feature from demo2!
this is Hello-3 from demo3 project

接下来利用上面【分支形式】或【commit-id】拉取，再执行 go run main.go
会发现使用的是老版本 demo1 go 包
this is Hello-1 from demo1
this is v1-new-feature from demo1!
this is Hello-2 from demo2 project
this is v1-new-feature from demo2!
this is Hello-3 from demo3 project

使用 【版本号形式】会报错，目前不清楚原因，猜测是因为子目录有 go.mod 文件导致
因此建议开发项目时，只创建一个 go.mod, 对于子目录中的包，采用【项目名】+【包名】形式引用，也就是之前说的解法3，详情查看 demo3 文件夹和 main.go 对其的引用
【版本号形式】会报错，报错信息如下：
go get github.com/oceanweave/hello/demo1@v1.0.1
go: github.com/oceanweave/hello/demo1@v1.0.1: invalid version: unknown revision demo1/v1.0.1
