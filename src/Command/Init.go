package Command

import "github.com/pinguo/pgo"

func init() {
    container := pgo.App.GetContainer()

    container.Bind(&TestCommand{})
}
