package gofs

type FileService interface {
	Upload(data GofsFile) error
	Download(data *GofsFile) error
}
