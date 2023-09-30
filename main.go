package main

import (
	"log"
	"threadsAPI/controller"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/router"
	"threadsAPI/usecase"
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

	dbConnect := db.OpenPostgresql()
	dbConnect.AutoMigrate(model.User{})
	dbConnect.AutoMigrate(model.Thread{})

	userRepository := repository.NewUserRepository(dbConnect)
	threadRepository := repository.NewThreadRpository(dbConnect)
	//データ作成テスト
	testThread := model.Thread{
		ID:       "test4",
		UserId:   "test3",
		Title:    "test",
		Contents: "testcontents",
	}
	threadRepository.CreateThread(&testThread)
	log.Printf("%+v", testThread)

	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
