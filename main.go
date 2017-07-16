package main

import (
	_ "github.com/Gahd/DDZServer/src/bll"
	_ "moqikaka.com/Framework/linuxMgr"
)

import (
	"sync"

	"github.com/Gahd/DDZServer/src/config"
	"github.com/Gahd/DDZServer/src/rpcServer"
	"github.com/Gahd/DDZServer/src/webServer"
	"moqikaka.com/Framework/monitorMgr"
	"moqikaka.com/Framework/signalMgr"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(1)
}

func main() {

	// 启动信号处理程序
	signalMgr.Start()

	//获取配置
	baseConfig := config.GetBaseConfig()
	monitorConfig := config.GetMonitorConfig()

	// 启动监控处理程序
	monitorMgr.Start(
		monitorConfig.MonitorServerIP,
		monitorConfig.MonitorServerName,
		monitorConfig.MonitorInterval,
	)

	//启动rpc服务器
	go rpcServer.StartServer(&wg)

	// 启动web服务器
	go webServer.Start(&wg, baseConfig.WebServerAddress)

	//阻塞,等待
	wg.Wait()
}
