package mysqlKit

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestNewGormDB(t *testing.T) {
	user := "yjs"
	password := "~Test123"
	addr := "localhost:3306"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/test?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr)
	log.Printf("dsn: %s", dsn)

	db, err := NewGormDB(dsn)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	{
		// User 默认表名users
		type User struct {
			UID  uint `gorm:"primaryKey"`
			Name string
		}

		if err := db.AutoMigrate(&User{}); err != nil {
			panic(err)
		}

		u := &User{}
		rst := db.Find(u)
		if rst.Error != nil {
			if errors.Is(rst.Error, gorm.ErrRecordNotFound) {
				panic("not found")
			} else {
				panic(err)
			}
		}
		log.Println(u)

		//u1 := map[string]interface{}{}
		//rst1 := db.Table("users").Take(&u1)
		//if rst1.Error != nil {
		//	panic(rst.Error)
		//}
		//log.Println(u1)
		//
		//db.Order()

		//f := &User{}
		//l := &User{}
		//result := db.First(f)
		//if result.Error != nil {
		//	panic(err)
		//}
		//result = db.Last(l)
		//if result.Error != nil {
		//	panic(err)
		//}
		//log.Println("first:", f)
		//log.Println("last:", l)

		//s := []*User{
		//	{Name: "张三"},
		//	{Name: "李四"},
		//	{Name: "王五"},
		//}
		//db.CreateInBatches(s, 2)
	}
}
