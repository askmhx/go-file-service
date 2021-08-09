package gofs

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
)

func PathIsExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func MinioFilePut(ctx context.Context, client *minio.Client, bucketName string, memo string, fileType string, file os.File) error {
	fInfo, _ := file.Stat()
	info, err := client.PutObject(ctx, bucketName, file.Name(), &file, fInfo.Size(), minio.PutObjectOptions{ContentType: "application/zip"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", file.Name(), info.Size)
	return nil
}

func MinioFileGet(ctx context.Context, client *minio.Client, bucketName string, fileName string) (os.File, error) {
	_, err := client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return os.File{}, nil
}

func MinioFileExist(ctx context.Context, client *minio.Client, bucketName string, fileName string) bool {
	_, err := client.StatObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	return true
}

func FileSave() {

}

func FileGet() {

}
