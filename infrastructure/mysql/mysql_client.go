package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MConfig struct {
	MaxIdle   int
	MaxActive int
	Configs   []MServerConfig
}

type MServerConfig struct {
	Name     string
	User     string
	Password string
	Url      string
	Db       string
}

var pools map[string]*gorm.DB
var mysqlDefaultDb string

func getConn(db string) *gorm.DB {
	pool, ok := pools[db]
	if ok {
		return pool
	} else {
		return pools[mysqlDefaultDb]
	}
}

func initMysql(mysqlConfig MConfig) {
	pools = make(map[string]*gorm.DB)
	var i int
	var configs = mysqlConfig.Configs
	var flag = true
	for i = 0; i < len(configs); i++ {
		config := configs[i]
		pool := initDb(config, mysqlConfig.MaxIdle, mysqlConfig.MaxActive)
		if pool != nil {
			pools[config.Name] = pool
			if flag {
				mysqlDefaultDb = config.Name
				flag = false
			}
		}
	}
}

func initDb(config MServerConfig, maxIdle int, maxActive int) *gorm.DB {
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Url + ")/" + config.Db + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(`mysql initDb: db(` + config.Url + `): ` + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(`mysql initDb: db(` + config.Url + `): ` + err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(maxActive)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
