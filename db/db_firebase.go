package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func OpenFireStorage()*storage.BucketHandle{
	config:=&firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_STORAGEBUCKET"),
	}
	opt:=option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICEACOUNTKEY"))
	app,err:=firebase.NewApp(context.Background(),config,opt)
	if err!=nil{
		log.Fatal(err)
	}
	client,err:=app.Storage(context.Background())
	if err!=nil{
		log.Fatal(err)
	}
	bucket,err:=client.DefaultBucket()
	if err!=nil{
		log.Fatal(err)
	}
	return bucket
}