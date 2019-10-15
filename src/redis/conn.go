package r

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
	client    redis.Conn
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func (redisConfig *RedisConfig) Dial() error {
	address := fmt.Sprintf("%v:%v", redisConfig.Host, redisConfig.Port)
	dbOption := redis.DialDatabase(redisConfig.Db)
	pwOption := redis.DialPassword(redisConfig.Password)
	// **重要** 设置读写超时
	readTimeout := redis.DialReadTimeout(time.Second * 60)
	writeTimeout := redis.DialWriteTimeout(time.Second * 60)
	conTimeout := redis.DialConnectTimeout(time.Second * 60)

	redisPool = &redis.Pool{
		MaxIdle:   500,
		MaxActive: 500,
		// **重要** 如果空闲列表中没有可用的连接
		// 且当前Active连接数 < MaxActive
		// 则等待
		Wait:        true,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, dbOption, pwOption,
				readTimeout, writeTimeout, conTimeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if nil != err {
				fmt.Println("redis ping error:"+err.Error(), "error")
			}
			return err
		},
	}
	client = redisPool.Get()

	err := redisPool.TestOnBorrow(client, time.Now())

	return err
}

//close redis pool
func Close() {
	err := redisPool.Close()
	if err != nil {
		fmt.Println("Close redis error", err)
	}
}
