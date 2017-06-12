package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	region = kingpin.Flag("region", "The region to use.").Required().String()
	bucket = kingpin.Flag("bucket", "The bucket to upload into.").Required().String()
	source = kingpin.Arg("source", "The file to upload.").Required().File()
	target = kingpin.Arg("target", "The name of the file.").Required().String()
)

func main() {
	kingpin.Parse()

	// Create the AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(*region),
	}))

	// Create an S3 uploader, this uploader allows multipart uploads
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(*target),
		Body:   *source,
	})

	// Check for errors
	if err != nil {
		fmt.Printf("failed to upload file, %v\n", err)
		return
	}

	fmt.Printf("Uploaded to %v\n", result.Location)
}
