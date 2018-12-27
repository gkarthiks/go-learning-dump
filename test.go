package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"context"
	"time"
)

func main() {
	fmt.Println("HAI")	
	s3Client, err := minio.New("XXXX", "XXXX", "XXXX", true)
	if err != nil {
		log.Fatalln(err)
	}

	found, err := s3Client.BucketExists("dsp")
	if err != nil {
		log.Fatalln(err)
	}

	if found {
		log.Println("Bucket found.")
	} else {
		log.Println("Bucket not found.")
	}


	// Trying to read the data
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	if err := s3Client.FGetObjectWithContext(ctx, "dsp", "glassdoor", "my-filename.csv", minio.GetObjectOptions{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully saved my-filename.csv")
}
