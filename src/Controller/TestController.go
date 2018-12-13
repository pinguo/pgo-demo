package Controller

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/Http"
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
    client := t.GetObject("@pgo/Client/Http/Adapter").(*Http.Adapter)

    response := client.Get("http://127.0.0.1:4000/sleep.php", nil, &Http.Option{Timeout: 2 * time.Second})
    content, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    fmt.Println("get response: ", string(content))


    t.str = "9999999"
    t.arr = []int{1, 2, 3}
    t.OutputJson("call /test/index, str:"+t.str, http.StatusOK)
}
