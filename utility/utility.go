package utility

type IUtility interface {
	ImgEndode(imgBytes []byte) string
	ImgFileEndode(path string) (string, error)
	ImgDecode(imgBase64 string) ([]byte, error)

	CreateUUID() (string, error)
}
type utility struct{}

func NewUtility() IUtility {
	return &utility{}
}