package Controller

import (
    "fmt"
    "net/http"

    "Service"

    "github.com/pinguo/pgo"
)

type WelcomeController struct {
    pgo.Controller
}

func (w *WelcomeController) Construct() {
    fmt.Println("welcome construct")
}

func (w *WelcomeController) Destruct() {
    fmt.Println("welcome destruct")
}

// 默认动作(index)
func (w *WelcomeController) ActionIndex() {
    w.OutputJson("hello world", http.StatusOK)
}

// 模板渲染
func (w *WelcomeController) ActionView() {
    // 获取并验证参数
    name := w.GetContext().ValidateParam("name", "hitzheng").Do()
    age := w.GetContext().ValidateParam("age", "100").Int().Do()

    data := pgo.Map{
        "name": name,
        "age":  age,
    }

    // 渲染html模板
    w.OutputView("welcome.html", data)
}

// URL路由控制器，根据url自动映射控制器及方法，不需要配置.
// url的最后一段为动作名称，不存在则为index,
// url的其余部分为控制器名称，不存在则为index,
// 例如：/path/to/welcome/say-hello，控制器类名为
// Path/To/WelcomeController 动作方法名为ActionSayHello
func (w *WelcomeController) ActionSayHello() {
    ctx := w.GetContext() // 获取PGO请求上下文件

    // 验证参数，提供参数名和默认值，当不提供默认值时，表明该参数为必选参数。
    // 详细验证方法参见Validate.go
    name := ctx.ValidateParam("name").Min(5).Max(50).Do()          // 验证GET/POST参数(string)，为空或验证失败时panic
    age := ctx.ValidateQuery("age", 20).Int().Min(1).Max(100).Do() // 只验证GET参数(int)，为空或失败时返回20
    ip := ctx.ValidatePost("ip", "").IPv4().Do()                   // 只验证POST参数(string), 为空或失败时返回空字符串

    // 打印日志
    ctx.Info("request from welcome, name:%s, age:%d, ip:%s", name, age, ip)
    ctx.PushLog("clientIp", ctx.GetClientIp()) // 生成clientIp=xxxxx在pushlog中

    // 调用业务逻辑，一个请求生命周期内的对象都要通过GetObject()获取，
    // 这样可自动查找注册的类，并注入请求上下文(Context)到对象中。
    svc := w.GetObject("Service/Welcome").(*Service.Welcome)

    // 添加耗时到profile日志中
    ctx.ProfileStart("Welcome.SayHello")
    svc.SayHello(name, age, ip)
    ctx.ProfileStop("Welcome.SayHello")

    data := pgo.Map{
        "name": name,
        "age":  age,
        "ip":   ip,
    }

    // 输出json数据
    w.OutputJson(data, http.StatusOK)
}

// 正则路由控制器，需要配置Router组件(components.router.rules)
// 规则中捕获的参数通过动作函数参数传递，没有则为空字符串.
// eg. "^/reg/eg/(\\w+)/(\\w+)$ => /welcome/regexp-example"
func (w *WelcomeController) ActionRegexpExample(p1, p2 string) {
    data := pgo.Map{"p1": p1, "p2": p2}
    w.OutputJson(data, http.StatusOK)
}

// RESTFULL动作，url中没有指定动作名，使用请求方法作为动作的名称(需要大写)
// 例如：GET方法请求ActionGET(), POST方法请求ActionPOST()
func (w *WelcomeController) ActionGET() {
    fmt.Println("call in Controller/WelcomeController.GET")
}

func (w *WelcomeController) ActionPOST() {

}
