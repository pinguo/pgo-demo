package Controller

import (
    "net/http"

    "github.com/pinguo/pgo"
    "github.com/pinguo/pgo/Client/MaxMind"
)

type MaxMindController struct {
    pgo.Controller
}

// curl -v http://127.0.0.1:8000/max-mind/geo-by-ip
func (m *MaxMindController) ActionGeoByIp() {
    // 获取要解析的IP地址
    ip := m.GetContext().ValidateParam("ip", "182.150.28.13").Do()

    // 获取MaxMind的上下文件适配对象
    mmd := m.GetObject(MaxMind.AdapterClass).(*MaxMind.Adapter)

    // 解析IP的geo信息
    geo := mmd.GeoByIp(ip)

    m.OutputJson(geo, http.StatusOK)
}
