package redis

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type RConfig struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Configs     []RServerConfig
}

type RServerConfig struct {
	Name     string
	Ip       string
	Password string
	Db       int
}

var pools map[string]*redigo.Pool
var redisDefaultDb string

func getConn(db string) redigo.Conn {
	pool, ok := pools[db]
	if ok {
		return pool.Get()
	} else {
		return pools[redisDefaultDb].Get()
	}
}

func releaseConn(conn redigo.Conn) {
	if conn != nil {
		conn.Close()
	}
}

func initRedis(redisConfig RConfig) {
	pools = make(map[string]*redigo.Pool)
	var i int
	var configs = redisConfig.Configs
	for i = 0; i < len(configs); i++ {
		config := configs[i]
		pool := initDb(config, redisConfig.MaxIdle, redisConfig.MaxActive, redisConfig.IdleTimeout)
		pools[config.Name] = pool
		if i == 0 {
			redisDefaultDb = config.Name
		}
	}
}

func initDb(config RServerConfig, maxIdle int, maxActive int, idleTimeout int) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,

		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", config.Ip, redigo.DialDatabase(config.Db))
			if err != nil {
				panic(`redis initDb: tcp(` + config.Ip + `): ` + err.Error())
			}
			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					panic(`redis initDb: AUTH(` + config.Ip + `): ` + err.Error())
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
