package gofs

import "time"

type GofsModel struct {
	AttachName   string
	AttachType   string
	AttachPath   string
	AttachExpire time.Time
	Memo         string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
}

type GofsResult struct {
	Code    string
	Message string
	Data    interface{}
}


type FileRepostory interface {
	Save(data GofsModel) bool
	Download(data *GofsModel)
}
