package main

import (
	"context"

	"gorm.io/gorm"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainServer, err := InitMainServer(ctx)
	if err != nil {
		return
	}
	mainServer.ApiServer.RegisterEndPoint()
	mainServer.ApiServer.G.Run(":1611")
}

type User struct {
	gorm.Model
	Name *string `gorm:"type:varchar(128)"`
}

// func main() {

// 	db, err := gorm.Open(mysql.New(mysql.Config{
// 		DSN: "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local",
// 	}))
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&User{})
// 	a := "minhbeo"
// 	db.Model(&User{}).Create(&User{Name: &a})
// }
