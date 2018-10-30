package Command

import "github.com/pinguo/pgo"

func init() {
    container := pgo.App.GetContainer()

    // 注册控制器类
    container.Bind(&TestCommand{})
}
