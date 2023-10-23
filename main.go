package main

import (
	"log"
	"threadsAPI/controller"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/router"
	"threadsAPI/utility"

	"threadsAPI/usecase"

	"github.com/joho/godotenv"
)

func main() {
	//テスト用のuser
	//echoでhttp接続できるようになるまで
	//img:=samplemethod.ImgEndode("./sampleImg/tester2.jpg")
	testUser := model.User{
		ID:       "test2",
		// LoginID:  "testLogin3",
		// Password: "passwordS",
		ImageUrl: "users/098333a4aafd46d78cb4511079c8583c/avator/avaterImg",
	}
	//データ作成テスト
	// testThread := model.Thread{
	// 	ID: "1d4ff667cfd4491b80f3591e8f9acc13",
	// 	UserId:   "098333a4aafd46d78cb4511079c8583c",
	// 	// Title:    "変更1013",
	// 	// Contents: "hennkousitanndasi!",
	// 	ImageUrl: img,
	// }
	// testMgs:=model.Message{
	// 	ThreadId: "0b233ecfd1f746588e10fdf8bbac1743",
	// 	UserId: "438aab35e9d5432d9cb468240d2688a5",
	// 	Message: "テスト1022",
	// 	ImageUrl: img,
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
	threadRepository := repository.NewThreadRpository(dbConnect,fbStorage)
	messageRepository:=repository.NewMessageRepository(dbConnect,fbStorage)

	//userRepository.DeleteUserImg(&testUser)
	
	util:=utility.NewUtility()
	userUsecase := usecase.NewUserUsecase(userRepository,util)
	threadUsecase := usecase.NewThreadUsecase(threadRepository,util)
	messageUsecase:=usecase.NewMessageUsecase(messageRepository,util)

	userUsecase.DeleteUserImg(testUser)

	userController := controller.NewUserController(userUsecase)
	threadController := controller.NewThreadController(threadUsecase)
	messageController:=controller.NewMessageController(messageUsecase)

	e := router.NewRouter(userController, threadController,messageController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
