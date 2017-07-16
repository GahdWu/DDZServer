package redisMgr

import (
	"moqikaka.com/goutil/debugUtil"
)

var (
	initSuccessMap = make(map[string]chan bool, 8)
)

// 注册初始化成功的通道
// name:模块名称
// ch:通道对象
func RegisterInitSuccess(name string, ch chan bool) {
	initSuccessMap[name] = ch
}

// 通知初始化成功
func notifyInitSuccess() {
	for name, ch := range initSuccessMap {
		debugUtil.Printf("通知:%s,Redis初始化成功\n", name)
		ch <- true
	}
}
