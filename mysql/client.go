package mysql

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
)

func init() {
	var err error

	username := "root"
	password := "example"
	protocol := "tcp"
	address := "127.0.0.1:3306"
	dbname := "belf"

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// username:password@protocol(address)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	// "root:example@tcp(127.0.0.1:3306)/belf?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, protocol, address, dbname)
	Client, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return
	}
}

type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

func CreateTest() {
	user := &User{Name: "kyungeun", Age: 22, Birthday: time.Now()}
	// 생성할 데이터의 포인터 넘기기
	result := Client.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
