package gofs

import (
	"rocket.iosxc.com/gateway/v1/model"
)

type FileService interface {
	Upload(data model.UploadRequest) model.CommonResult
	Download(data model.DownloadRequest) model.CommonResult
}