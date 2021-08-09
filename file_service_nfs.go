package gofs

import (
	"context"
	"fmt"
)

type nfsFileService struct {
	ctx      context.Context
	repo     FilePersistenceCtrl
	basePath string
}

func (this nfsFileService) Upload(data GofsFile) error {

	//util.FilePut(this.ctx, this.minio, data.MerchantId, "", "FILE", data.File)

	attach := GofsFile{
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

func (this nfsFileService) Download(data model.DownloadRequest) model.CommonResult {

	fileName := fmt.Sprintf("CHANNEL-%s-%s.txt", data.FileType, data.FileDate)

	//if !util.FileExist(this.ctx, this.minio, data.MerchantId, fileName) {
	//	ret := model.CommonResult{
	//		Code:    model.RESULT_CODE_FILE_NOT_FOUND,
	//		Message: "upload success",
	//		Data:    nil,
	//	}
	//	return ret
	//}

	token := ""
	downloadUrl := fmt.Sprintf("https://fileserver.rocket.iosxc.com/download/%s/%s?token=", data.FileDate, fileName, token)
	ret := model.CommonResult{
		Code:    model.RESULT_CODE_SUCCESS,
		Message: "Get File URL success",
		Data:    downloadUrl,
	}
	return ret
}

func NewNFSFileService(basePath string, repo FilePersistenceCtrl) FileService {
	return &nfsFileService{
		repo:     repo,
		basePath: basePath,
	}
}
