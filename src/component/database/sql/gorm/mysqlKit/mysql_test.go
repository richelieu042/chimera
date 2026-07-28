package mysqlKit

import (
	"fmt"
	"log"
	"testing"
)

func TestNewGormDB(t *testing.T) {
	user := "myapp_user"
	password := "~Test123"
	addr := "152.53.167.107:3306"
	dbName := "myapp_db"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, addr, dbName)
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

	// User 默认表名users
	type User struct {
		UID  uint `gorm:"primaryKey"`
		Name string
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	/* （1）插入 */
	//{
	//	s := []*User{
	//		{Name: "张三"},
	//		{Name: "李四"},
	//		{Name: "王五"},
	//	}
	//	db.CreateInBatches(s, 2)
	//}

	/* （2）查（第一条、最后一条） */
	{
		f := &User{}
		l := &User{}
		result := db.First(f)
		if result.Error != nil {
			panic(err)
		}
		result = db.Last(l)
		if result.Error != nil {
			panic(err)
		}
		log.Println("first:", f)
		log.Println("last:", l)
		fmt.Println("---")
	}

	//{
	//	u1 := map[string]interface{}{}
	//	rst1 := db.Table("users").Take(&u1)
	//	if rst1.Error != nil {
	//		panic(rst1.Error)
	//	}
	//	log.Println(u1)
	//	fmt.Println("---")
	//}

	//u := &User{}
	//rst := db.Find(u)
	//if rst.Error != nil {
	//	if errors.Is(rst.Error, gorm.ErrRecordNotFound) {
	//		panic("not found")
	//	} else {
	//		panic(err)
	//	}
	//}
	//log.Println(u)

	//db.Order()

}
