package models

import (
	"bill/modules/log"
	"bill/modules/setting"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-redis/redis_rate/v8"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
	xormLog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

type Record map[string]interface{}

type BaseQuery struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

// MasterDB 数据库对象
var MasterDB *xorm.Engine

// MasterRedis Redis连接池
var MasterRedis *redis.Client

var RedisLimiter *redis_rate.Limiter

func initDB() {
	dns := buildMysqlConnStr()
	db, err := xorm.NewEngine("mysql", dns)
	if err != nil {
		log.Get().Panic(fmt.Sprintf("数据库初始化失败:%s", err))
		return
	}
	err = db.Ping()
	if err != nil {
		log.Get().Panic(fmt.Sprintf("数据库无法访问:%s", dns))
	}
	showSql := setting.GetConfig().Xorm.ShowSql
	logLevel := setting.GetConfig().Xorm.LogLevel
	timezone := setting.GetConfig().Xorm.Timezone

	db.ShowSQL(showSql)
	db.Logger().SetLevel(xormLog.LogLevel(logLevel))
	db.TZLocation, _ = time.LoadLocation(timezone)
	db.SetMapper(names.GonicMapper{})
	MasterDB = db
}

func initRedis() {
	url := fmt.Sprintf("%s:%d", setting.GetConfig().Redis.Host, setting.GetConfig().Redis.Port)
	pwd := setting.GetConfig().Redis.Password
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pwd,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Get().Panic(fmt.Sprintf("Redis连接错误:%s", err))
	}

	MasterRedis = rdb

	RedisLimiter = redis_rate.NewLimiter(MasterRedis)
}

// 从配置文件读取mysql配置并返回连接字符串
func buildMysqlConnStr() string {
	username := setting.GetConfig().Mysql.User

	password := setting.GetConfig().Mysql.Password
	host := setting.GetConfig().Mysql.Host
	port := setting.GetConfig().Mysql.Port
	dbname := setting.GetConfig().Mysql.DBName
	charset := setting.GetConfig().Mysql.Charset

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, host, port, dbname, charset)

	return conn
}

func init() {

	initDB()
	initRedis()
}
