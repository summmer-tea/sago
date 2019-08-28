package sago

import (
	"github.com/garyburd/redigo/redis"
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
	if GConfig.Redis.Addr == "" {
		return
	}

	redisConn = &redis.Pool{
		//连接池控制
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", GConfig.Redis.Addr,
				redis.DialConnectTimeout(time.Second*30),
				redis.DialPassword(GConfig.Redis.Password))
			if err != nil {
				panic(err)
			} else {
				return c, nil
			}

		},
	}
	//连接池数量
	redisConn.MaxIdle = 15
	redisConn.MaxActive = 20
	redisConn.IdleTimeout = 5 * time.Second

}

//func (m *RedisDialect) getDbConn() redis.Conn {
//
//	return redisConn.Get()
//}

func createRedisDialect() redis.Conn {
	once_redis.Do(func() {
		redisDialect.RegisterDbConn()
	})
	return redisConn.Get()
}

func (m *RedisDialect) Set(key string, val string) (err error) {

	//rconn := redisConn.Get()
	//defer func() {
	//	if err := rconn.Close(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	if _, err := redisConn.Get().Do("SET", key, val); err != nil {
		return err
	}

	return nil
}

func (m *RedisDialect) Get(key string) (reply interface{}, err error) {

	//rconn := redisConn.Get()
	//defer func() {
	//	if err := rconn.Close(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	if reply, err := redis.String(redisConn.Get().Do("GET", key)); err != nil {
		panic(err)
	} else {
		return reply, err
	}
}
