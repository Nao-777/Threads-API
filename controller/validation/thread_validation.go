package validation

import (
	"threadsAPI/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IThreadValidation interface {
	ThreadValidate(thread model.Thread) error 
}
type threadValidation struct{}

func NewThreadValidation() IThreadValidation {
	return &threadValidation{}
}
func (tv *threadValidation) ThreadValidate(thread model.Thread) error {
	return validation.ValidateStruct(
		&thread,
		validation.Field(
			&thread.Title,
			validation.Required.Error("Title is required"),
			validation.RuneLength(1,100).Error("limited min 1 max 100"),
		),
		validation.Field(
			&thread.Contents,
			validation.Required.Error("Contents is required"),
			validation.RuneLength(1,500).Error("limited min 1 max 500"),
		),
		validation.Field(
			&thread.ImageUrl,
			validation.By(judgeImageSize),
		),
	)
}