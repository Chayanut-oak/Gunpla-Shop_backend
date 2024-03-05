package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/google/uuid"
// )

// func S3uploader(w http.ResponseWriter, r *http.Request) {
// 	awsEndpoint := "http://localhost:4566"
// 	awsRegion := "us-east-1"
// 	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 		if awsEndpoint != "" {
// 			return aws.Endpoint{
// 				PartitionID:   "aws",
// 				URL:           awsEndpoint,
// 				SigningRegion: awsRegion,
// 			}, nil
// 		}

// 		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
// 	})

// 	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithRegion(awsRegion),
// 		config.WithEndpointResolverWithOptions(customResolver),
// 	)
// 	if err != nil {
// 		log.Fatalf("Cannot load the AWS configs: %s", err)
// 	}

// 	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
// 		o.UsePathStyle = true
// 	})

// 	bucketName := "don-gunpla-store"
// 	var imageUrls []string
// 	for _, files := range r.MultipartForm.File {
// 		fmt.Println(files)
// 		for _, file := range files {
// 			objectKey := uuid.NewString() + ".png"
// 			// Open the uploaded file
// 			src, err := file.Open()
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 			defer src.Close()
// 			_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
// 				Bucket: &bucketName,
// 				Key:    &objectKey,
// 				Body:   src,
// 			})
// 			if err != nil {
// 				log.Fatalf("Error uploading picture: %v", err)
// 			}
// 			imageUrls = append(imageUrls, fmt.Sprintf("%s-%s-%s", awsEndpoint, bucketName, objectKey))
// 			// Prepare the S3 upload parameters
// 		}
// 	}

// 	log.Printf("Picture uploaded successfully to S3://%s", bucketName)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(imageUrls)
// }

// // CORS middleware function to add CORS headers
// func CORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Allow requests from any origin
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		// Allow the GET, POST, and OPTIONS methods
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		// Allow the Content-Type header
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		// If it's an OPTIONS request, send an empty response with status 200
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		// Call the next handler in the chain
// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {
// 	// Attach your handler function to the desired route
// 	http.HandleFunc("/api/upload-image", S3uploader)

// 	// Use the CORS middleware
// 	handler := CORS(http.DefaultServeMux)

// 	// Start the server
// 	fmt.Println("Server listening on :8080")
// 	http.ListenAndServe(":8080", handler)
// }
import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func S3uploader(c *gin.Context) {
	awsEndpoint := "http://localhost:4566"
	awsRegion := "us-east-1"

	err := c.Request.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	var imageUrls []string // Slice to hold image URLs

	fmt.Println(c.Request.MultipartForm.File)
	for _, files := range c.Request.MultipartForm.File {

		for _, file := range files {
			objectKey := uuid.NewString() + ".png"
			// Open the uploaded file
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
			// Append the URL of the uploaded image to imageUrls

			imageUrls = append(imageUrls, fmt.Sprintf("%s/%s/%s", awsEndpoint, bucketName, objectKey))
		}
	}

	c.JSON(http.StatusOK, gin.H{"imageUrls": imageUrls})
}

func main() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}

	r.Use(cors.New(config))

	r.POST("/s3/upload-image", S3uploader)
	r.Run(":8080")
}
