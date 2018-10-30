package main

import (
    // 由于controller/command是逻辑入口，
    // 所以这两个包需要手动导入以执行注册流程
    _ "Command"    // 导入Command
    _ "Controller" // 导入Controller

    "github.com/pinguo/pgo"
)

func main() {
    // 启动服务，默认以web模式运行，接入并处理http请求，
    // 以命令行模式运行：bin/pgo-demo --cmd /test/index 执行命令行控制器
    // 设置配置环境：bin/pgo-demo --env production
    pgo.Run()
}
