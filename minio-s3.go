package main

import (
	"github.com/minio/minio-go"
	"log"
)

func main() {
	s3Client, err := minio.New("xxxx", "xxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxx", true)
	if err != nil {
		log.Fatalln(err)
	}

	found, err := s3Client.BucketExists("xxxxxx")
	if err != nil {
		log.Fatalln(err)
	}

	if found {
		log.Println("Bucket found.")
	} else {
		log.Println("Bucket not found.")
	}
}