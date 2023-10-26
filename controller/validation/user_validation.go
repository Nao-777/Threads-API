package validation

import (
	"fmt"
	"regexp"
	"threadsAPI/model"
	"threadsAPI/utility"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)



type IUserValidation interface {
	UserValidate(user model.User)error
}
type uservalidation struct{}

func NewUserValidation() IUserValidation{
	return &uservalidation{}
}
func (uv *uservalidation) UserValidate(user model.User)error{

	return validation.ValidateStruct(&user,
		validation.Field(
			&user.LoginID,
			validation.Required.Error("loginID is required"),
			validation.RuneLength(6,24).Error("limited min 6 max 24"),
			validation.Match(regexp.MustCompile(`^[a-zA-z0-9]+$`)).Error("does not match half-width alphanumeric characters"),
		),
		validation.Field(
			&user.Name,
			validation.Required.Error("Name is required"),
			validation.RuneLength(1,24).Error("limited max 24"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password id required"),
			validation.RuneLength(8,24).Error("limited min 8 max 24"),
			validation.Match(regexp.MustCompile(`^[a-zA-z0-9]+$`)).Error("does not match half-width alphanumeric characters"),
			validation.By(judgePasswordStrength),
		),
		validation.Field(
			&user.ImageUrl,
			validation.By(judgeImageSize),
		),
	)
}

//半角英数字の大文字、小文字、数字がそれぞれ一つ以上存在するか確認する
func judgePasswordStrength(value interface{})error{
	str,_:=value.(string)
	reg := []*regexp.Regexp{
		regexp.MustCompile(`[a-z]`),
		regexp.MustCompile(`[A-Z]`),
		regexp.MustCompile(`[0-9]`),
	}
	for i, r := range reg {
		if r.FindString(str) == "" {
			switch i {
			case 0:
				return fmt.Errorf("there are no lowercase letters")
			case 1:
				return fmt.Errorf("there are no uppercase letters")
			case 2:
				return fmt.Errorf("there are no numbers")
			}
		}
	}
	return nil
}
func judgeImageSize(value interface{})error{
	imgStr,_:=value.(string)
	imgBytes,err:=utility.NewUtility().ImgDecode(imgStr)
	if err!=nil{
		return err
	}
	if len(imgBytes)>100000{
		return fmt.Errorf("image size is too large")
	}
	return nil
}