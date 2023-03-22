package main

import (
	"fmt"
	"time"

	"github.com/robinson/gos7"
)

type PlcData struct {
	boolValue    bool
	intValue     uint16
	realValue    float32
	stringValue  string
	wstringValue string
}

func main() {
	const (
		ipAddr = "192.168.123.199" //PLC IP
		rack   = 0                 // PLC机架号
		slot   = 1                 // PLC插槽号
	)
	//PLC tcp连接客户端
	handler := gos7.NewTCPClientHandler(ipAddr, rack, slot)
	//连接及读取超时
	handler.Timeout = 200 * time.Second
	//关闭连接超时
	handler.IdleTimeout = 200 * time.Second
	//打开连接
	handler.Connect()
	//函数退出时关闭连接
	defer handler.Close()

	//获取PLC对象
	client := gos7.NewClient(handler)

	//DB号
	address := 1
	//起始地址
	start := 0
	//读取字节数
	size := 162
	//读写字节缓存区
	buffer := make([]byte, size)

	//读取字节
	client.AGReadDB(address, start, size, buffer)

	//gos7解析数据类
	var helper gos7.Helper

	//gos7内置方法解析数据
	var data PlcData
	helper.GetValueAt(buffer[160:162], 0, &data.intValue)
	// data.intValue = (uint16)helper.GetCounterAt(buffer[160:162], 0)
	// data.stringValue = helper.GetStringAt(buffer[8:264], 0)
	// data.wstringValue = helper.GetWStringAt(buffer[160:], 2)

	//输出数据
	fmt.Println(data)
	fmt.Println(buffer)
}
