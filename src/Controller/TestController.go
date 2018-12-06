package Controller

import (
    "fmt"
    "net/http"

    "github.com/pinguo/pgo"
)

type TestController struct {
    pgo.Controller
}

func (t *TestController) Construct() {
    fmt.Println("test construct")
}

func (t *TestController) Destruct() {
    fmt.Println("test destruct")
}

func (t *TestController) ActionIndex() {
    t.OutputJson("call /test/index", http.StatusOK)
}
