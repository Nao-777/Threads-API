package utility

import (
	b64 "encoding/base64"
)

type IUtility interface {
	ImgEndode(imgBytes []byte) string
}
type utility struct {}

func NewUtility() IUtility{
	return &utility{}
}

// 送られてくる画像がエンコードされていると仮定
func (ut *utility)ImgEndode(imgBytes []byte) string {
	imgStr:=b64.StdEncoding.EncodeToString(imgBytes)
	return imgStr
}
// func ImgDecode(imgBase64 string) []byte {
// 	uDec, err := b64.StdEncoding.DecodeString(imgBase64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return uDec
// }