package helper

import (
	"career/config"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	firebase "firebase.google.com/go"
	"github.com/thanhpk/randstr"
	"google.golang.org/api/option"
)

func FirebaseImageUploader(ctx context.Context, image *multipart.FileHeader, folder string) (string, error) {
	// Initialize the Firebase app
	bucket_name := config.FirebaseBucketName()
	config := &firebase.Config{
		ProjectID:     config.FirebaseProjectId(),
		StorageBucket: config.FirebaseBucketName(),
	}
	opt := option.WithCredentialsFile("firebase-storage-sa.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return "", err
	}
	// Get a reference to the Firebase Storage bucket
	storageClient, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}
	bucket, err := storageClient.Bucket(bucket_name)
	if err != nil {
		return "", err
	}
	// Upload the file to the specified path in the bucket
	fileName := randstr.Hex(16)
	obj := bucket.Object(fmt.Sprintf("%s/%s", folder, fileName))
	writer := obj.NewWriter(ctx)
	openedImage, err := image.Open()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(writer, openedImage)
	if err != nil {
		log.Printf("error uploading image file: %v\n", err)
		return "", err
	}
	if err := writer.Close(); err != nil {
		log.Printf("error closing object writer: %v\n", err)
		return "", err
	}
	// Get a signed URL for the uploaded file that will be valid for 1 hour
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", string(bucket_name), folder+"%2F"+string(fileName))
	return url, nil
}

func FirebaseMultipleImageUploader(ctx context.Context, images []*multipart.FileHeader, folder string) ([]string, error) {
	var urls []string

	// Initialize the Firebase app
	bucket_name := config.FirebaseBucketName()
	config := &firebase.Config{
		ProjectID:     config.FirebaseProjectId(),
		StorageBucket: config.FirebaseBucketName(),
	}
	opt := option.WithCredentialsFile("firebase-storage-sa.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return urls, err
	}
	// Get a reference to the Firebase Storage bucket
	storageClient, err := app.Storage(ctx)
	if err != nil {
		return urls, err
	}
	bucket, err := storageClient.Bucket(bucket_name)
	if err != nil {
		return urls, err
	}

	for _, image := range images {
		fileName := randstr.Hex(16)
		object := bucket.Object(fmt.Sprintf("%s/%s", folder, fileName))
		writer := object.NewWriter(ctx)
		openedImg, err := image.Open()
		if err != nil {
			return urls, err
		}
		_, err = io.Copy(writer, openedImg)
		if err != nil {
			log.Printf("error uploading image file: %v\n", err)
			return urls, err
		}
		if err := writer.Close(); err != nil {
			log.Printf("error closing object writer: %v\n", err)
			return urls, err
		}
		urls = append(urls, fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", string(bucket_name), folder+"%2F"+string(fileName)))
	}
	return urls, nil
}
