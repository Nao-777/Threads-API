package utility

import (
	b64 "encoding/base64"
)

type IUtility interface {
	ImgEndode(imgBytes []byte) string
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
func (ut *utility)ImgDecode(imgBase64 string) ([]byte,error) {
	uDec, err := b64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return []byte{},err
	}
	return uDec,nil
}