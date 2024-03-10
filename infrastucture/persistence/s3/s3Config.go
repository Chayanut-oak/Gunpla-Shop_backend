package s3

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func S3uploader(c *gin.Context) {
    awsRegion := "us-east-1"

    err := c.Request.ParseMultipartForm(10 << 20)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    awsCfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(awsRegion),
    )
    if err != nil {
        log.Fatalf("Cannot load the AWS configs: %s", err)
    }

    client := s3.NewFromConfig(awsCfg)

    bucketName := "don-gunpla-store"
    var imageUrls []string

    for , files := range c.Request.MultipartForm.File {
        for , file := range files {
            objectKey := uuid.NewString() + ".png"
            src, err := file.Open()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            defer src.Close()

            _, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
                Bucket: &bucketName,
                Key:    &objectKey,
                Body:   src,
            })
            if err != nil {
                log.Fatalf("Error uploading picture: %v", err)
            }

            imageUrls = append(imageUrls, fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, awsRegion, objectKey))
        }
    }

    c.JSON(http.StatusOK, gin.H{"imageUrls": imageUrls})
}