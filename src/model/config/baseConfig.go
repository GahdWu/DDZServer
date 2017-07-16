package config

import (
	"fmt"

	. "moqikaka.com/goutil/configUtil"
)

type BaseConfig struct {

	//WEB服务地址
	WebServerAddress string

	// 是否是DEBUG模式
	DEBUG bool
}

func NewEmptyBaseConfig() *BaseConfig {
	return &BaseConfig{}
}

func NewBaseConfig(webServerAddress string, debug bool) *BaseConfig {
	return &BaseConfig{
		WebServerAddress: webServerAddress,
		DEBUG:            debug,
	}
}

// 加载配置
func LoadBaseConfig(xmlConfig *XmlConfig) (*BaseConfig, error) {
	// 解析WebServerAddress
	webServerAddress, err := xmlConfig.String("Root/WebServerAddress", "")
	if err != nil {
		return nil, err
	}

	// 解析DEBUG配置
	debug, err := xmlConfig.Bool("Root/DEBUG", "")
	if err != nil {
		return nil, err
	}

	return NewBaseConfig(webServerAddress, debug), nil
}

//转化为字符串
func (this *BaseConfig) String() string {
	return fmt.Sprintf("BaseConfig,WebServerAddress:%s,DEBUG:%t", this.WebServerAddress, this.DEBUG)
}
