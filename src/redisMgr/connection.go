package redisMgr

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"moqikaka.com/goutil/logUtil"
	"time"
)

var (
	redisPool *redis.Pool
)

// 创建redis连接池
// redisConnection：Redis服务器连接地址
// password：Redis服务器连接密码
// database：Redis服务器选择的数据库
// maxActive：Redis连接池允许的最大活跃连接数量
// maxIdle：Redis连接池允许的最大空闲数量
// idleTimeout：连接被回收前的空闲时间
// dialConnectTimeout：连接Redis服务器超时时间
// 返回值：无
func createRedisPool(redisConnection, password string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) {
	redisPool = &redis.Pool{
		MaxActive:   maxActive,
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0, 4)
			options = append(options, redis.DialConnectTimeout(dialConnectTimeout))
			if password != "" {
				options = append(options, redis.DialPassword(password))
			}
			options = append(options, redis.DialDatabase(database))
			c, err := redis.Dial("tcp", redisConnection, options...)
			if err != nil {
				logUtil.Log(fmt.Sprintf("Dial failed, err:%s", err), logUtil.Error, true)
				return nil, fmt.Errorf("Dial failed, err:%s", err)
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// 获取Redis连接
// 返回值：
// Redis连接
func getRedisConn() redis.Conn {
	return redisPool.Get()
}
