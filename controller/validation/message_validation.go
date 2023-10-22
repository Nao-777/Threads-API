package validation

import (
	"threadsAPI/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IMessageValidation interface {
	MessageValidate(msg model.Message) error 
}
type messageValidation struct{}

func NewMessageValidation() IMessageValidation {
	return &messageValidation{}
}
func (mv *messageValidation)MessageValidate(msg model.Message) error {
	return validation.ValidateStruct(
		&msg,
		validation.Field(
			&msg.Message,
			validation.Required.Error("message is required"),
			validation.RuneLength(1,500).Error("limited min 1 max 500"),
		),
		validation.Field(
			&msg.ImageUrl,
			validation.By(judgeImageSize),
		),
	)
}