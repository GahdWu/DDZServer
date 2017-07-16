package config

import (
	"fmt"

	"github.com/Gahd/DDZServer/src/model/config"
	"moqikaka.com/Framework/reloadMgr"
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
)

var (
	//基础配置
	baseConfig *config.BaseConfig

	//检测配置
	monitorConfig *config.MonitorConfig

	//Redis数据库配置
	redisConfig *config.RedisConfig
)

func init() {
	logUtil.SetLogPath("Log")

	if e := reload(); e != nil {
		panic(fmt.Errorf("加载配置失败，错误信息为：%s", e))
	}

	// 注册重新加载的方法
	reloadMgr.RegisterReloadFunc("config.reload", reload)
}

func reload() error {

	var err error

	//读取配置文件内容
	xmlConfig := configUtil.NewXmlConfig()
	err = xmlConfig.LoadFromFile("config.xml")
	if err != nil {
		return err
	}

	//加载基础配置
	baseConfig, err = config.LoadBaseConfig(xmlConfig)
	if err != nil {
		return err
	}

	//加载检测配置
	monitorConfig, err = config.LoadMonitorConfig(xmlConfig)
	if err != nil {
		return err
	}

	//加载Redis配置
	redisConfig, err = config.LoadRedisConfig(xmlConfig)
	if err != nil {
		return err
	}

	// 设置debugUtil的状态
	debugUtil.SetDebug(baseConfig.DEBUG)

	//打印配置
	debugUtil.Println(baseConfig.String())
	debugUtil.Println(monitorConfig.String())
	debugUtil.Println(redisConfig.String())

	return nil
}

/*
	获取基础配置副本
	参数：
	返回值：基础配置副本
*/
func GetBaseConfig() config.BaseConfig {
	return *baseConfig
}

/*
	获取Redis配置副本
	参数：
	返回值：Redis配置副本
*/
func GetRedisConfig() config.RedisConfig {
	return *redisConfig
}

/*
	获取检测配置副本
	参数：
	返回值：检测配置副本
*/
func GetMonitorConfig() config.MonitorConfig {
	return *monitorConfig
}
