package db

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3UploadClient contains the necessary fields for interacting with S3
type S3UploadClient struct {
    Client   *s3.Client
    Bucket   string
}

// NewS3UploadClient initializes the S3 client
func NewS3UploadClient() (*S3UploadClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel() 
	bucket := os.Getenv("S3_BUCKET_NAME")
    region := os.Getenv("AWS_REGION")
    accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")


	// Load AWS config with credentials from .env
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(region),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
    )
    if err != nil {
        // is fatal
		log.Fatalf("unable to load SDK config, %v", err)
    }

    s3Client := s3.NewFromConfig(cfg)
    return &S3UploadClient{
        Client: s3Client,
        Bucket: bucket,
    }
}

// UploadFileToS3 uploads a PDF to the S3 bucket
func (s *S3UploadClient) UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
    defer file.Close()

    // Ensure the file is a PDF
    if filepath.Ext(fileHeader.Filename) != ".pdf" {
        return "", fmt.Errorf("file type not supported")
    }

    key := fmt.Sprintf("uploads/%s-%s", time.Now().Format("20060102150405"), fileHeader.Filename)

    _, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(s.Bucket),
        Key:    aws.String(key),
        Body:   file,
        ContentType: aws.String("application/pdf"),
    })

    if err != nil {
        return "", err
    }

    return key, nil
}