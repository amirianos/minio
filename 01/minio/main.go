package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func main() {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "minio_user"
	secretAccessKey := "minio_password"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	buckets, _ := minioClient.ListBuckets(context.Background())
	//show list of buckets
	fmt.Println(buckets)

	// Upload the zip file
	objectName := "tech.pdf"
	filePath := "/usr/local/go/src/lab/test1/minio/tech.pdf"
	//contentType := "application/zip"
	contentType := "pdf"
	bucketName := "testbucket"
	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	//log.Printf("%#v\n", minioClient) // minioClient is now set up
}
