package main

import (
	"threadsAPI/db"
	"threadsAPI/model"
)

func main() {
	dbConnect := db.OpenPostgresql()
	dbConnect.AutoMigrate(model.User{})
	defer db.CloseDB(dbConnect)
}
