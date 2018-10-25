package Controller

import (
    "github.com/pinguo/pgo"
)

type WelcomeController struct {
    pgo.Controller
}

func (w *WelcomeController) ActionIndex() {
    name := w.GetContext().ValidateParam("name", "hitzheng").Do()
    age := w.GetContext().ValidateParam("age", "100").Int().Do()

    data := pgo.Map{
        "name": name,
        "age":  age,
    }

    w.OutputView("welcome.html", data)
}

func (w *WelcomeController) ActionGET() {

}

func (w *WelcomeController) ActionPOST() {

}
