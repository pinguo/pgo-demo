package Controller

import (
    "fmt"
    "net/http"

    "github.com/pinguo/pgo"
)

type TestController struct {
    pgo.Controller

    str string
    arr []int
}

func (t *TestController) Construct() {
    fmt.Printf("test construct, str:%s, arr:%v, addr:%p\n", t.str, t.arr, t)
}

func (t *TestController) ActionIndex() {
    t.str = "9999999"
    t.arr = []int{1, 2, 3}
    t.OutputJson("call /test/index, str:"+t.str, http.StatusOK)
}
