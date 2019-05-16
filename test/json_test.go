package test

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"testing"
)

/**
*@Author: PantaZheng
*@CreateAt: 2019/5/9 10:56
*@Title: json_test.go
*@Package: main
*@Description: 测试JSON
@Software: GoLand
*/

func TestJson(t *testing.T) {
	userJSON := UserJSON{ID: 1, IDCard: "1111"}
	log.Println(json.Marshal(userJSON))
	log.Println(reflect.TypeOf(userJSON.Telephone))
}

type User struct {
	gorm.Model
	OpenID    string `gorm:"unique"`
	Code      string
	Name      string
	IDCard    string
	Level     int
	Telephone string
}

//UserJSON 用户Json原型
type UserJSON struct {
	/**
	@Author: PantaZheng
	@Description:用户JSON
	@Date: 2019/5/9 10:42
	*/
	ID        uint   `json:"id,omitempty"`
	OpenID    string `json:"openid,omitempty"`
	Code      string `json:"code,omitempty"`
	Name      string `json:"name,omitempty"`
	IDCard    string `json:"idCard,omitempty"`
	Level     int    `json:"level"`
	Telephone string `json:"telephone,omitempty"`
}
