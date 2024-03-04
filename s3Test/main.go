package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func S3uploader(w http.ResponseWriter, r *http.Request) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"
	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	bucketName := "don-gunpla-store"

	for _, files := range r.MultipartForm.File {
		fmt.Println(files)
		for _, file := range files {
			objectKey := uuid.NewString() + ".png"
			// Open the uploaded file
			src, err := file.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			// Prepare the S3 upload parameters
		}
	}

	log.Printf("Picture uploaded successfully to S3://%s", bucketName)
}

// CORS middleware function to add CORS headers
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow the GET, POST, and OPTIONS methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// Allow the Content-Type header
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// If it's an OPTIONS request, send an empty response with status 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Attach your handler function to the desired route
	http.HandleFunc("/api/upload-image", S3uploader)

	// Use the CORS middleware
	handler := CORS(http.DefaultServeMux)

	// Start the server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", handler)
}
