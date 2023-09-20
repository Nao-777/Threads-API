// DBの接続を記述
package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgresqlに接続
func OpenPostgresql() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=threadsAPI_DB port=5434 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("postgresDB connected")
	return db
}

// postgresqlの接続を終了
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}
}
