package main

import (
	"log"
	"threadsAPI/controller"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/router"
	"threadsAPI/usecase"

	"github.com/joho/godotenv"
)

func main() {
	//テスト用のuser
	//echoでhttp接続できるようになるまで
	// testUser := model.User{
	// 	ID:       "test2",
	// 	LoginID:  "testLogin2",
	// 	Name:     "tester",
	// 	Password: "password",
	// }
	//開発時だけ読み込むようにしたい
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	dbConnect := db.OpenPostgresql()
	dbConnect.AutoMigrate(model.User{})
	dbConnect.AutoMigrate(model.Thread{})
	dbConnect.AutoMigrate(model.Message{})

	userRepository := repository.NewUserRepository(dbConnect)
	threadRepository := repository.NewThreadRpository(dbConnect)
	//データ作成テスト
	// testThread := model.Thread{

	// 	UserId:   "f87de508-4ae3-45c5-a652-694facd1c1be",
	// 	Title:    "test",
	// 	Contents: "testcontents",
	// }
	userUsecase := usecase.NewUserUsecase(userRepository)
	threadUsecase := usecase.NewThreadUsecase(threadRepository)

	userController := controller.NewUserController(userUsecase)
	threadController := controller.NewThreadController(threadUsecase)

	e := router.NewRouter(userController, threadController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
