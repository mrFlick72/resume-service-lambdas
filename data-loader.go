package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")}))
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(os.Getenv("s3.bucket.file.name"))
	fmt.Println("error", err)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("s3.bucket")),
			Key:    aws.String(os.Getenv("s3.bucket.file.key")),
		})

	fmt.Println("error", err)
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	defer file.Close()
}
