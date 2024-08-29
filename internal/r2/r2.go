package r2

import (
	"context"
	"fmt"
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
	println("AKI", accessKey, secretKey, bucket, account)
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

	println(list.Buckets)
}
