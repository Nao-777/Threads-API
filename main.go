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
	// testThread := model.Thread{
	// 	ID:       "test5",
	// 	UserId:   "f87de508-4ae3-45c5-a652-694facd1c1be",
	// 	Title:    "test",
	// 	Contents: "testcontents",
	// }
	userUsecase := usecase.NewUserUsecase(userRepository)
	threadUsecase := usecase.NewThreadUsecase(threadRepository)

	//threadUsecase.CreateThread(&testThread)
	//threads, err := threadUsecase.GetThreadsByUserID("f87de508-4ae3-45c5-a652-694facd1c1be")
	threads, err := threadUsecase.GetThreads(2, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range threads {
		log.Printf("%+v\n", t)
	}

	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
