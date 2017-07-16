package config

import (
	"fmt"
	"time"

	. "moqikaka.com/goutil/configUtil"
)

type RedisConfig struct {
	// Redis连接字符串
	RedisConnectionString string

	// Redis密码
	RedisPassword string

	// Redis数据库编号
	RedisDatabase int

	// Redis最大活跃连接数
	RedisMaxActive int

	// Redis最大空闲连接数
	RedisMaxIdle int

	// Redis空闲超时时间
	RedisIdleTimeout time.Duration

	// Redis连接超时时间
	RedisDialConnectTimeout time.Duration

	// Redis中Token过期的秒数
	RedisTokenExpireSeconds int
}

func NewEmptyRedisConfig() *RedisConfig {
	return &RedisConfig{}
}

func NewRedisConfig(redisConnectionString, redisPassword string, redisDatabase, redisMaxActive, redisMaxIdle, redisTokenExpireSeconds int, redisIdleTimeout, redisDialConnectTimeout time.Duration) *RedisConfig {
	return &RedisConfig{
		RedisConnectionString:   redisConnectionString,
		RedisPassword:           redisPassword,
		RedisDatabase:           redisDatabase,
		RedisMaxActive:          redisMaxActive,
		RedisMaxIdle:            redisMaxIdle,
		RedisIdleTimeout:        redisIdleTimeout,
		RedisDialConnectTimeout: redisDialConnectTimeout,
		RedisTokenExpireSeconds: redisTokenExpireSeconds,
	}
}

// 加载配置
func LoadRedisConfig(xmlConfig *XmlConfig) (*RedisConfig, error) {
	// 解析Redis连接字符串
	redisConnectionString, err := xmlConfig.String("Root/Redis/ConnectionString", "")
	if err != nil {
		return nil, err
	}

	// 解析Redis密码
	redisPassword, err := xmlConfig.String("Root/Redis/Password", "")
	if err != nil {
		return nil, err
	}

	// 解析Redis数据库编号
	redisDatabase, err := xmlConfig.Int("Root/Redis/Database", "")
	if err != nil {
		return nil, err
	}

	// 解析Redis最大活跃连接数
	redisMaxActive, err := xmlConfig.Int("Root/Redis/MaxActive", "")
	if err != nil {
		return nil, err
	}

	// 解析Redis最大空闲连接数
	redisMaxIdle, err := xmlConfig.Int("Root/Redis/MaxIdle", "")
	if err != nil {
		return nil, err
	}

	// 解析Redis空闲超时时间
	idleTimeout, err := xmlConfig.Int("Root/Redis/IdleTimeout", "")
	if err != nil {
		return nil, err
	}
	redisIdleTimeout := time.Duration(idleTimeout) * time.Second

	// 解析Redis连接超时时间
	dialConnectTimeout, err := xmlConfig.Int("Root/Redis/DialConnectTimeout", "")
	if err != nil {
		return nil, err
	}
	redisDialConnectTimeout := time.Duration(dialConnectTimeout) * time.Second

	// 解析Redis中Token过期的秒数
	redisTokenExpireSeconds, err := xmlConfig.Int("Root/Redis/TokenExpireSeconds", "")
	if err != nil {
		return nil, err
	}

	return NewRedisConfig(
		redisConnectionString, redisPassword, redisDatabase,
		redisMaxActive, redisMaxIdle, redisTokenExpireSeconds,
		redisIdleTimeout, redisDialConnectTimeout), nil
}

//转化为字符串
func (this *RedisConfig) String() string {
	return fmt.Sprintf("RedisConfig,RedisConnectionString:%s,RedisPassword:%s,RedisDatabase:%d,RedisMaxActive:%d,RedisMaxIdle:%d,RedisIdleTimeout:%v,RedisDialConnectTimeout:%v,RedisTokenExpireSeconds:%d",
		this.RedisConnectionString, this.RedisPassword, this.RedisDatabase, this.RedisMaxActive, this.RedisMaxIdle, this.RedisIdleTimeout, this.RedisDialConnectTimeout, this.RedisTokenExpireSeconds)
}
