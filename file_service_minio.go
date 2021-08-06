package

import (
	"context"
	"fmt"
)
package gofs


import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"rocket.iosxc.com/gateway/v1/model"
	"rocket.iosxc.com/gateway/v1/repository"
	"rocket.iosxc.com/gateway/v1/util"
)


type minioFileService struct {
	ctx   context.Context
	repo  repository.MerchantAttachRepository
	minio *minio.Client
}

func (this minioFileService) Upload(data model.UploadRequest) model.CommonResult {

	util.FilePut(this.ctx, this.minio, data.MerchantId, "", "FILE", data.File)

	attach := model.MerchantAttach{
		MerchantId: data.MerchantId,
		AttachName: data.FileName,
		AttachType: "FILE",
		AttachPath: data.FileName,
		CreatedBy:  "system",
	}
	this.repo.Save(attach)

	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "upload success",
		Data:    nil,
	}
	return ret
}

func (this minioFileService) Download(data model.DownloadRequest) model.CommonResult {

	fileName := fmt.Sprintf("CHANNEL-%s-%s.txt", data.FileType, data.FileDate)

	if !util.FileExist(this.ctx, this.minio, data.MerchantId, fileName) {
		ret := model.CommonResult{
			Code:    model.RESULT_CODE_FILE_NOT_FOUND,
			Message: "upload success",
			Data:    nil,
		}
		return ret
	}

	token := ""
	downloadUrl := fmt.Sprintf("https://fileserver.rocket.iosxc.com/download/%s/%s?token=", data.FileDate, fileName, token)
	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "Get File URL success",
		Data:    downloadUrl,
	}
	return ret
}

func NewMinioFileService(repo repository.MerchantAttachRepository) FileService {
	return &minioFileService{
		repo: repo,
	}
}
