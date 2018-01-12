package mysql
//https://github.com/go-xorm/xorm
import (
//	"fmt"
	"sync"
	"time"
//	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	once sync.Once
	aa string
)

func init() {
//	if engine == nil {
//		lock.Lock()
//		defer lock.Unlock()
//
//		engine := ""
//	}
	// 单例
	once.Do(func() {

		var err error
		engine, err = xorm.NewEngine("mysql", "root:hello123@tcp(127.0.0.1:3306)/putao_dictcrawl?charset=utf8")

		if err != nil {
			panic(err)
		}

		engine.ShowSQL(true)
		engine.SetMaxIdleConns(30)      // 设置连接池的空闲数大小
		engine.SetMaxOpenConns(30)      // 设置最大打开连接数
		engine.SetConnMaxLifetime(time.Duration(28800) * time.Second)		// 设置连接可以重用的最大时间 空闲时间 28800为Mysql默认时间
	})
}

func TT(key string) interface{} {

	sql := "select * from dict_0 where word like '%"+key+"%'"
	results, _ := engine.QueryString(sql)

	return results
//	return data
}
