package gofs

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type minioFileService struct {
	ctx   context.Context
	repo  FileRepostory
	minio *minio.Client
}

func (m minioFileService) Upload(data GofsFile) error {
	MinioFilePut(this.ctx, this.minio, data.MerchantId, "", "FILE", data.File)
	attach := GofsFile{
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

func (m minioFileService) Download(data *GofsFile) error {

	fileName := fmt.Sprintf("CHANNEL-%s-%s.txt", data.FileType, data.FileDate)

	if !MinioFileExist(this.ctx, this.minio, data.MerchantId, fileName) {
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

func NewMinioFileService(repo FileRepostory, minio *minio.Client, ctx context.Context) FileService {
	return &minioFileService{
		repo:  repo,
		minio: minio,
		ctx:   ctx,
	}
}
