package main

import (
	"fmt"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
)

func main() {
	//テスト用のuser
	//echoでhttp接続できるようになるまで
	testUser := model.User{
		ID:       "test2",
		LoginID:  "testLogin2",
		Name:     "tester",
		Password: "password",
	}
	dbConnect := db.OpenPostgresql()
	dbConnect.AutoMigrate(model.User{})
	userRepository := repository.NewUserRepository(dbConnect)
	//サンプルユーザデータの追加
	userRepository.InsertUser(&testUser)
	//サンプルユーザデータの取得
	testUserRes := model.User{}
	userRepository.GetUserByLoginId(&testUserRes, "testLogin1")
	fmt.Printf("%+v", testUserRes)
	defer db.CloseDB(dbConnect)
}
