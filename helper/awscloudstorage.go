package helper

import (
	"context"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/internal"
	"github.com/aws/aws-sdk-go-v2/aws/option"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
) 

type StorageS3Config struct {
	S3Client *s3.Client 
	BucketName string
	FolderName string
}

func UploadImage(c echo.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}

	image, err := file.Open()
	if err != nil {
		return "", err
	}
	defer image.Close() 

	s3c := StorageS3Config{
		S3Client: InitS3Client(),
		BucketName: "<your-bucket-name>",
		FolderName: "<your-folder-name>",
	}

	fileName := file.Filename
	// setelah mendapatkan ekstensi file, gunakan ekstensi tersebut sebagai ContentType
	fileExtension := filepath.Ext(fileName)
	contentType := "image/" + fileExtension [1:] 
	fileNameOnS3 := s3c.FolderName + "/" + fileName 
	// Membuat PutObjectInput dengan metadata file 
	input := &s3c.PutObjectInput{
		Bucket: 		aws.String(s3c.BucketName),
		Key: 			aws.String(fileNameOnS3),
		Body:           image, 
		contentType:    aws.String(contentType),
	}

	// Mengunggah file ke S3 
	_, err = s3c.S3Client.PutObject(context.Background(), input)
	if err != nil {
		return "", err
	}

	// Mengembalikan URL file yang diunggah 
	return "https://" + s3c.BucketName + ".s3.amazonaws.com/" + fileNameOnS3, nil
} 

func InitS3Client() *s3.Client {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Println("Tidak dapat memuat konfigurasi AWS:", err)
		os.Exit(1)
	}
	// konfigurasi tambahan seperti mengubah region dan lainnya 
	cfg.Region = "<your-region>"
	opts := []func (*s3.Options){}
	// untuk mengkonfigurasi kredensial dari enviroment variable, hilangkan opsi ini 
	opts = append(opts, func(o *s3.Options) {
		o.Credential = aws.NewEnvCredentials()
	})

	// Membuat klien s3 
	s3Client := s3.NewFromConfig(cfg, opts...)
	return s3Client
}

