package Controller

import (
    "fmt"
    "time"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/Memcache"
)

type MemcacheController struct {
    pgo.Controller
}

// curl -v http://127.0.0.1:8000/memcache/set
func (m *MemcacheController) ActionSet() {
    // 获取memcache的上下文适配对象
    mc := m.GetObject(Memcache.AdapterClass).(*Memcache.Adapter)

    // 设置用户输入值
    key := m.GetContext().ValidateParam("key", "test_key1").Do()
    val := m.GetContext().ValidateParam("val", "test_val1").Do()
    mc.Set(key, val)

    // 设置自定义过期时间
    mc.Set("test_key2", 100, 2*time.Minute)

    // 设置map值，会自动进行json序列化
    data := pgo.Map{"f1": 100, "f2": true, "f3": "hello"}
    mc.Set("test_key3", data)

    // 并行设置多个key
    items := pgo.Map{
        "test_key4": []int{1, 2, 3, 4},
        "test_key5": "test_val5",
        "test_key6": pgo.Map{"f61": 11, "f62": "hello"},
    }
    mc.MSet(items)
}

func (m *MemcacheController) ActionGet() {
    // 获取memcache的上下文适配对象
    mc := m.GetObject(Memcache.AdapterClass).(*Memcache.Adapter)

    // 获取string
    if val := mc.Get("test_key1"); val != nil {
        fmt.Println("value of test_key1:", val.String())
    }

    // 获取int
    if val := mc.Get("test_key2"); val != nil {
        fmt.Println("value of test_key2:", val.Int())
    }

    // 获取序列化的数据
    if val := mc.Get("test_key3"); val != nil {
        var data pgo.Map
        val.Decode(&data)
        fmt.Println("value of test_key3:", data)
    }

    // 获取多个key
    if res := mc.MGet([]string{"test_key4", "test_key5", "test_key6"}); res != nil {
        for key, val := range res {
            fmt.Printf("value of %s: %v\n", key, val.String())
        }
    }
}
