package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("sa-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"ID-KEY",
				"SECRET-KEY",
				"TOKEN-KEY",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "go-expert-bucket-exemple"
}

func main() {
	dir, err := os.Open("./tmp/")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 50)
	go func() { // Thread especifica para retentar upload de erros de arquivos.
		for {
			select {
			case fileName := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(fileName, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error while reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	fmt.Printf("Starting to upload %s to bucket %s\n", fileName, s3Bucket)
	file, err := os.Open(fmt.Sprintf("./tmp/%s", fileName))
	if err != nil {
		fmt.Printf("Error opening file %s\n", fileName)
		<-uploadControl             // esvazia o canal
		errorFileUpload <- fileName // preenche o canal
		return
	}
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", fileName)
		<-uploadControl             // esvazia o canal
		errorFileUpload <- fileName // preenche o canal
		return
	}
	fmt.Printf("Uploaded file successfully %s\n", fileName)
	<-uploadControl // esvazia o canal
}
