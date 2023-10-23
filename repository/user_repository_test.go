package repository

import (
	"context"
	"log"
	"os"
	"testing"
	"threadsAPI/db"
	"threadsAPI/model"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func TestDeleteUserImg(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	fbStorage:=openFireStorage()
	dbConnect := db.OpenPostgresql()
	ur:=NewUserRepository(dbConnect,fbStorage);
	tests :=[]struct{
		name string
		args model.User
		want string
	}{
		{
			"case-DeleteUserImg",
			model.User{
				ImageUrl: "users/098333a4aafd46d78cb4511079c8583c/avator/avaterImg",
			},
			"",
		},
	}
	for _,test:=range tests{
		t.Run(test.name,func(t *testing.T){
			var got string
			err:=ur.DeleteUserImg(&test.args);
			if err==nil{
				got=""
			}else{
				got=err.Error()
			}
			
			if got!=test.want{
				t.Errorf("want:%s |now:%s\n",test.want,got)
			}
		})
	}
}
func openFireStorage()*storage.BucketHandle{
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