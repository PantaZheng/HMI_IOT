package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pantazheng/bci/config"
	"log"
)

/**
*设置数据库连接
*@param diver string
 */

var (
	DB = New()
)

func New() *gorm.DB {
	mysqlConfig := config.Conf.MySql
	connect := mysqlConfig.UserName + ":" + mysqlConfig.Password + "@/" + mysqlConfig.DBName + "?charset" + mysqlConfig.Charset + "&collate=" + mysqlConfig.Collation + "&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysqlConfig.DBDriver, connect)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
	log.Printf("建立数据库连接\n")
	return DB
}
