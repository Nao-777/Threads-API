package samplemethod

import (
	"bytes"
	b64 "encoding/base64"
	"image/jpeg"
	"log"
	"os"
)

//送られてくる画像がエンコードされていると仮定
func ImgEndode(path string) string {
	data, err := os.Open(path)
	if err!=nil{
		log.Fatal(err)
	}
	img,err:=jpeg.Decode(data)
	if err!=nil{
		log.Fatal(err)
	}
	data.Close()
	bf:=new(bytes.Buffer)
	if err:=jpeg.Encode(bf,img,nil);err!=nil{
		log.Fatal(err)
	}
	imgBytes:=bf.Bytes()

	sEnc:=b64.StdEncoding.EncodeToString(imgBytes)
	log.Println(sEnc)
	return sEnc
}