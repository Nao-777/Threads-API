package main

import (
	"log"
	"threadsAPI/controller"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/router"
	samplemethod "threadsAPI/sampleMethod"
	"threadsAPI/usecase"

	"github.com/joho/godotenv"
)

func main() {
	//テスト用のuser
	//echoでhttp接続できるようになるまで
	img:=samplemethod.ImgEndode("./sampleImg/firebasetest.jpg")
	testUser := model.User{
		ID:       "test2",
		// LoginID:  "testLogin3",
		// Password: "passwordS",
		ImageUrl: img,
	}
	//データ作成テスト
	// testThread := model.Thread{
	// 	ID: "0b233ecfd1f746588e10fdf8bbac1743",
	// 	// UserId:   "f87de508-4ae3-45c5-a652-694facd1c1be",
	// 	Title:    "変更1013",
	// 	Contents: "hennkousitanndasi!",
	// }
	//開発時だけ読み込むようにしたい
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	fbStorage:=db.OpenFireStorage()
	dbConnect := db.OpenPostgresql()
	dbConnect.AutoMigrate(model.User{})
	dbConnect.AutoMigrate(model.Thread{})
	dbConnect.AutoMigrate(model.Message{})

	userRepository := repository.NewUserRepository(dbConnect,fbStorage)
	threadRepository := repository.NewThreadRpository(dbConnect)
	messageRepository:=repository.NewMessageRepository(dbConnect)

	//userRepository.PostUserImg(&testUser)
	//userRepository.GetUserImg(&testUser)
	
	
	userUsecase := usecase.NewUserUsecase(userRepository)
	threadUsecase := usecase.NewThreadUsecase(threadRepository)
	messageUsecase:=usecase.NewMessageUsecase(messageRepository)
	
	if err:=userUsecase.PostUserImg(testUser);err!=nil{
		log.Fatal(err)
	}

	userController := controller.NewUserController(userUsecase)
	threadController := controller.NewThreadController(threadUsecase)
	messageController:=controller.NewMessageController(messageUsecase)

	e := router.NewRouter(userController, threadController,messageController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
