package config

import (
	"fmt"

	. "moqikaka.com/goutil/configUtil"
)

type MonitorConfig struct {
	// 监控使用的服务器IP
	MonitorServerIP string

	// 监控使用的服务器名称
	MonitorServerName string

	// 监控的时间间隔（单位：分钟）
	MonitorInterval int
}

func NewEmptyMonitorConfig() *MonitorConfig {
	return &MonitorConfig{}
}

func NewMonitorConfig(monitorServerIP, monitorServerName string, monitorInterval int) *MonitorConfig {
	return &MonitorConfig{
		MonitorServerIP:   monitorServerIP,
		MonitorServerName: monitorServerName,
		MonitorInterval:   monitorInterval,
	}
}

// 加载配置
func LoadMonitorConfig(xmlConfig *XmlConfig) (*MonitorConfig, error) {
	// 解析监控使用的服务器IP
	monitorServerIP, err := xmlConfig.String("Root/Monitor/ServerIP", "")
	if err != nil {
		return nil, err
	}

	// 解析监控使用的服务器名称
	monitorServerName, err := xmlConfig.String("Root/Monitor/ServerName", "")
	if err != nil {
		return nil, err
	}

	// 解析监控的时间间隔
	monitorInterval, err := xmlConfig.Int("Root/Monitor/Interval", "")
	if err != nil {
		return nil, err
	}

	return NewMonitorConfig(monitorServerIP, monitorServerName, monitorInterval), nil
}

//转化为字符串
func (this *MonitorConfig) String() string {
	return fmt.Sprintf("MonitorConfig,MonitorServerIP:%s,MonitorServerName:%s,MonitorInterval:%d",
		this.MonitorServerIP, this.MonitorServerName, this.MonitorInterval)
}
