package utility

import (
	"bytes"
	b64 "encoding/base64"
	"image/jpeg"
	"os"
)

type IUtility interface {
	ImgEndode(imgBytes []byte) string
	ImgFileEndode(path string) (string,error)
	ImgDecode(imgBase64 string) ([]byte,error)
}
type utility struct {}

func NewUtility() IUtility{
	return &utility{}
}

func (ut *utility)ImgEndode(imgBytes []byte) string {
	imgStr:=b64.StdEncoding.EncodeToString(imgBytes)
	return imgStr
}
func (ut *utility)ImgFileEndode(path string) (string,error) {
	data, err := os.Open(path)
	if err!=nil{
		return "",err
	}
	img,err:=jpeg.Decode(data)
	if err!=nil{
		return "",err
	}
	data.Close()
	bf:=new(bytes.Buffer)
	if err:=jpeg.Encode(bf,img,nil);err!=nil{
		return "",err
	}
	imgBytes:=bf.Bytes()
	sEnc:=b64.StdEncoding.EncodeToString(imgBytes)
	return sEnc,nil
}
func (ut *utility)ImgDecode(imgBase64 string) ([]byte,error) {
	uDec, err := b64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return []byte{},err
	}
	return uDec,nil
}