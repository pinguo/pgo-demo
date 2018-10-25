package Controller

import (
    "net/http"
    "time"

    "github.com/pinguo/pgo"
)

type WelcomeController struct {
    pgo.Controller
}

func (w *WelcomeController) ActionIndex() {
    data := pgo.Map{"text": "welcome to pgo-demo", "now": time.Now()}
    w.OutputJson(data, http.StatusOK)
}
