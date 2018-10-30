package Command

import (
    "flag"
    "fmt"

    "github.com/pinguo/pgo"
)

// 命令行控制器
type TestCommand struct {
    pgo.Controller
}

// 命令行控制器的输入参数需要手动解析
func (t *TestCommand) ActionIndex() {
    args := flag.Args()
    fmt.Println("call in Command/Test.Index, args:", args)
}
