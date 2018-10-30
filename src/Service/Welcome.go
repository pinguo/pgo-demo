package Service

import (
    "fmt"

    "github.com/pinguo/pgo"
)

type Welcome struct {
    pgo.Object
}

// 框架自动调用的构造函数(可选)
func (w *Welcome) Construct() {
    fmt.Printf("call in Service/Welcome.Construct\n")
}

// 框架自动调用的初始函数(可选)
func (w *Welcome) Init() {
    fmt.Printf("call in Service/Welcome.Init\n")
}

func (w *Welcome) SayHello(name string, age int, sex string) {
    fmt.Printf("call in  Service/Welcome.SayHello, name:%s age:%d sex:%s\n", name, age, sex)
}
