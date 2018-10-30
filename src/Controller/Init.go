package Controller

import "github.com/pinguo/pgo"

// 实践中init流程最好集中放到一个init方法中，
// 并放到包的Init.go文件中。
func init() {
    container := pgo.App.GetContainer()

    // 注册控制器到PGO容器中
    container.Bind(&WelcomeController{})
}
