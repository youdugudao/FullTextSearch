# golang实现的全文搜索
## 设计
- 使用http协议，restful的交互风格
- 分词使用`github.com/huichen/sego`
- 将文章生成的索引，保存在mongodb中
## 准备工作
- 需要安装[mongodb](https://docs.mongodb.com/manual/administration/install-enterprise/)
- 安装[go](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.1.md)
- 安装go依赖包
    - `go get -u -v gopkg.in/gin-gonic/gin.v1`
    - `go get -u -v gopkg.in/mgo.v2`
## 启动
`go run main.go`