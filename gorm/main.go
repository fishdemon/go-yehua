package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

// gorm 默认的表明是结构体的复数 users
type User struct {
	ID string
	Name string
	Age int8
	Sex string
	Username string
	Password string
}

func testCRUD(db *gorm.DB)  {
	user := User{"12346", "john", 23, "M", "john","123456"}
	res := db.Create(&user)
	if res.Error == nil {
		json1, _ := json.Marshal(user)
		fmt.Println(string(json1))
	}

	var user1 User
	res = db.First(&user1, "12346")
	if res.Error == nil {
		json1, _ := json.Marshal(user1)
		fmt.Println(string(json1))
	}

	db.Model(&user1).Update("age", 49)

	res = db.First(&user1)
	if res.Error == nil {
		json1, _ := json.Marshal(user1)
		fmt.Println(string(json1))
	}

	db.Delete(&user1)
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@/go-study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 禁用表名的复数形式
	db.SingularTable(true)

	return db, nil
}

func main() {
	fmt.Println("Starting")

	//initDB()
	//testCRUD()

	//generateProtoMsg()
	//readProtoMsg()

	Db, err := initDB()
	if err != nil {
		fmt.Println("failed to connect mysql server")
	}

	testCRUD(Db)

}

