package main

import (
    "fmt"
    "github.com/go-amz/amz/aws"
    "github.com/go-amz/amz/s3"
    "log"
)

func main() {
    auth := aws.Auth{
        AccessKey: "ASDFASDFASDFASDK",
        SecretKey: "DSFSDFDWESDADSFASDFADFDSFASDF",
    }
    euwest := aws.EUWest

    connection := s3.New(auth, euwest)
    mybucket := connection.Bucket("mytotallysecretbucket")
    res, err := mybucket.List("", "", "", 1000)
    if err != nil {
        log.Fatal(err)
    }
    for _, v := range res.Contents {
        fmt.Println(v.Key)
    }
}