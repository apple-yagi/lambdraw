package uploader

type Uploader interface {
	Execute() error
}