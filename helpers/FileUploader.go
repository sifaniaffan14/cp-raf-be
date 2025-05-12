package helpers

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint  = "localhost:9000"
	accessKey = "minioadmin"
	secretKey = "minioadmin"
	bucket    = "cp-raf"
	useSSL    = false
)

func UploadToS3(fileHeader *multipart.FileHeader) (string, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	objectName := fmt.Sprintf("banners/%d-%s", time.Now().Unix(), fileHeader.Filename)
	contentType := fileHeader.Header.Get("Content-Type")

	_, err = client.PutObject(context.Background(), bucket, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	imageURL := fmt.Sprintf("http://%s/%s/%s", endpoint, bucket, objectName)
	return imageURL, nil
}