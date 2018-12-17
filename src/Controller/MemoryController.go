package Controller

import (
    "fmt"
    "time"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/Memory"
)

type MemoryController struct {
    pgo.Controller
}

// curl -v http://127.0.0.1:8000/memory/set
func (m *MemoryController) ActionSet() {
    // 获取memory的上下文适配对象
    mm := m.GetObject(Memory.AdapterClass).(*Memory.Adapter)

    // 设置用户输入值
    key := m.GetContext().ValidateParam("key", "test_key1").Do()
    val := m.GetContext().ValidateParam("val", "test_val1").Do()
    mm.Set(key, val)

    // 设置自定义过期时间
    mm.Set("test_key2", 100, 2*time.Minute)

    // 设置map值，会自动进行json序列化
    data := pgo.Map{"f1": 100, "f2": true, "f3": "hello"}
    mm.Set("test_key3", data)

    // 并行设置多个key
    items := pgo.Map{
        "test_key4": []int{1, 2, 3, 4},
        "test_key5": "test_val5",
        "test_key6": pgo.Map{"f61": 11, "f62": "hello"},
    }
    mm.MSet(items)
}

// curl -v http://127.0.0.1:8000/memory/get
func (m *MemoryController) ActionGet() {
    // 获取memory的上下文适配对象
    mm := m.GetObject(Memory.AdapterClass).(*Memory.Adapter)

    // 获取string
    if val := mm.Get("test_key1"); val != nil {
        fmt.Println("value of test_key1:", val.String())
    }

    // 获取int
    if val := mm.Get("test_key2"); val != nil {
        fmt.Println("value of test_key2:", val.Int())
    }

    // 获取序列化的数据
    if val := mm.Get("test_key3"); val != nil {
        var data pgo.Map
        val.Decode(&data)
        fmt.Println("value of test_key3:", data)
    }

    // 获取多个key
    if res := mm.MGet([]string{"test_key4", "test_key5", "test_key6"}); res != nil {
        for key, val := range res {
            fmt.Printf("value of %s: %v\n", key, val.String())
        }
    }
}
