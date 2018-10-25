package Command

import (
    "fmt"

    "github.com/pinguo/pgo"
)

type TestCommand struct {
    pgo.Controller
}

func (t *TestCommand) ActionIndex() {
    fmt.Println("command test.index run")
}
