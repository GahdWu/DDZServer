package webServer

import (
	"fmt"
	"net/http"
	"sync"

	"moqikaka.com/goutil/logUtil"
)

// 启动Web服务器
// wg:WaitGroup对象
// address:服务器地址
func Start(wg *sync.WaitGroup, webAddress string) {
	defer func() {
		wg.Done()
	}()

	logUtil.NormalLog(fmt.Sprintf("Web服务器开始监听:%s...", webAddress), logUtil.Info)

	// 启动Web服务器监听
	err := http.ListenAndServe(webAddress, new(ServerHandler))
	if err != nil {
		panic(fmt.Errorf("ListenAndServe失败，错误信息为：%s", err))
	}
}
