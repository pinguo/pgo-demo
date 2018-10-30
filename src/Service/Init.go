package Service

import "github.com/pinguo/pgo"

func init() {
    container := pgo.App.GetContainer()

    // 注册类
    container.Bind(&Welcome{})

    // 除控制器目录外，其它包的init函数中应该只注册该包的类，
    // 而不应该包含子包。
}
