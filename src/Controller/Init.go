package Controller

import "github.com/pinguo/pgo"

// 实践中init流程最好集中放到一个init方法中，
// 并放到包的Init.go文件中。
// 由于controller是程序的逻辑入口，因此controller
// 需要在main.go中手动导入，这里将所有的controller
// (包括子包)在一个地方注册，避免在main.go中逐个导入。
func init() {
    container := pgo.App.GetContainer()

    // 注册控制器到PGO容器中
    container.Bind(&TestController{})
    container.Bind(&WelcomeController{})
    container.Bind(&MysqlController{})

    // 注册其它控制器(包含子包)
}
