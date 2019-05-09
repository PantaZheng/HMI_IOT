package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pantazheng/bci/config"
	"github.com/pelletier/go-toml"
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
	driver := config.Conf.Get("database.driver").(string)
	configTree := config.Conf.Get(driver).(*toml.Tree)
	userName := configTree.Get("databaseUserName").(string)
	password := configTree.Get("databasePassword").(string)
	databaseName := configTree.Get("databaseName").(string)
	connect := userName + ":" + password + "@/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(driver, connect)

	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
	log.Printf("建立数据库连接\n")
	return DB
}
