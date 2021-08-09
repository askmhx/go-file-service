package gofs

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type minioFileService struct {
	ctx   context.Context
	repo  FilePersistenceCtrl
	minio *minio.Client
}

func (this minioFileService) Upload(data GofsFile) error {
	MinioFilePut(this.ctx, this.minio, data.BucketName, data.Memo, data.FileType, data.FileAttach)
	if this.repo != nil {
		return this.repo.Save(data)
	}
	return nil
}

func (this minioFileService) Download(data *GofsFile) error {

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

func NewMinioFileService(ctx context.Context, minio *minio.Client, repo FilePersistenceCtrl) FileService {
	return &minioFileService{
		repo:  repo,
		minio: minio,
		ctx:   ctx,
	}
}
