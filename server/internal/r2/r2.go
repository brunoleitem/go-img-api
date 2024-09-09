package r2

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2service struct {
	client *s3.Client
	bucket string
}

func NewR2Service() (*R2service, error) {
	account := os.Getenv("R2_ACCOUNT_ID")
	accessKey := os.Getenv("R2_ACCESS_KEY")
	secretKey := os.Getenv("R2_SECRET_KEY")
	bucket := os.Getenv("R2_BUCKET")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("auto"),
	)

	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", account))
	})
	fmt.Println("INICIOU R2")

	return &R2service{
		client: client,
		bucket: bucket,
	}, nil
}

func (r *R2service) ListBuckets(ctx context.Context) {
	list, err := r.client.ListBuckets(ctx, nil)
	if err != nil {
		panic(err)
	}
	nome := *list.Buckets[0].Name
	println("Listando buckets: ")
	println(nome)
}

func (r *R2service) UploadImage(ctx context.Context, key *string, img io.Reader, contentType string) error {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &r.bucket,
		Key:         key,
		Body:        img,
		ContentType: &contentType,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *R2service) DeleteImage(ctx context.Context, imageId *string) error {
	fmt.Println("version")
	fmt.Println(imageId)
	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucket,
		Key:    imageId,
	})
	if err != nil {
		return err
	}

	return nil
}
