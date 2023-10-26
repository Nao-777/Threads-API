package utility

import (
	"strings"

	"github.com/google/uuid"
)

func (ut *utility) CreateUUID() (string, error) {
	//userIDを生成
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	//ハイフンを除去、文字列にした値をユーザIDとして登録
	userId := strings.Replace(uuid.String(), "-", "", -1)
	return userId,nil
}