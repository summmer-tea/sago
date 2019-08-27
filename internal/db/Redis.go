package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sago/internal/comm"
	"sync"
	"time"
)

var (
	once_redis   sync.Once
	redisConn    *redis.Pool
	redisDialect *RedisDialect
)

type RedisDialect struct{}

func (m *RedisDialect) RegisterDbConn() {

	//配置项为空跳过
	if comm.G_config.Redis.Addr == "" {
		return
	}

	redisConn = &redis.Pool{
		//连接池控制
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", comm.G_config.Redis.Addr,
				redis.DialConnectTimeout(time.Second*30),
				redis.DialPassword(comm.G_config.Redis.Password))
			if err != nil {
				panic(err)
			} else {
				return c, nil
			}

		},
	}
	//连接池数量
	redisConn.MaxIdle = 10
	redisConn.MaxActive = 10
	redisConn.IdleTimeout = 60 * time.Second

}

func (m *RedisDialect) GetDbConn() redis.Conn {

	return redisConn.Get()
}

func CreateRedisDialect() redis.Conn {
	once_redis.Do(func() {
		redisDialect.RegisterDbConn()
	})
	return redisDialect.GetDbConn()
}

func (m *RedisDialect) Set(key string, val string) (err error) {

	rconn := redisDialect.GetDbConn()
	defer func() {
		if err := rconn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if _, err := rconn.Do("SET", key, val); err != nil {
		return err
	}

	return nil
}

func (m *RedisDialect) Get(key string) (reply interface{}, err error) {

	rconn := redisDialect.GetDbConn()
	defer func() {
		if err := rconn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if reply, err := redis.String(rconn.Do("GET", key)); err != nil {
		panic(err)
	} else {
		return reply, err
	}
}
