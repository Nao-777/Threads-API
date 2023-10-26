package main

import (
	"log"
	"threadsAPI/controller"
	"threadsAPI/controller/validation"
	"threadsAPI/db"
	"threadsAPI/model"
	"threadsAPI/repository"
	"threadsAPI/router"
	"threadsAPI/utility"

	"threadsAPI/usecase"

	"github.com/joho/godotenv"
)

func main() {
	
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
	
	util:=utility.NewUtility()
	userUsecase := usecase.NewUserUsecase(userRepository,util)
	threadUsecase := usecase.NewThreadUsecase(threadRepository,util)
	messageUsecase:=usecase.NewMessageUsecase(messageRepository,util)

	userValidation:=validation.NewUserValidation()
	threadValidation:=validation.NewThreadValidation()
	messageValidation:=validation.NewMessageValidation()
	userController := controller.NewUserController(userUsecase,userValidation)
	threadController := controller.NewThreadController(threadUsecase,threadValidation)
	messageController:=controller.NewMessageController(messageUsecase,messageValidation)

	e := router.NewRouter(userController, threadController,messageController)
	e.Logger.Fatal(e.Start(":8080"))
	defer db.CloseDB(dbConnect)
}
