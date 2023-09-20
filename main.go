package main

import (
	"threadsAPI/db"
)

func main() {
	dbConnect := db.OpenPostgresql()
	defer db.CloseDB(dbConnect)
}
