package gofs

import (
	"os"
	"time"
)

type GofsFile struct {
	FileAttach os.File
	FileName   string
	FileType   string
	BucketName string
	Memo       string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
}

type FilePersistenceCtrl interface {
	Save(data GofsFile) error
	Query(data *GofsFile) GofsFile
}
