package service

import (
	"config"
	"context"
	"fmt"

	r2Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type R2Service struct{}

func (r2Service R2Service) PreSign(obj string, R2Bucket *config.R2Bucket) (string, bool) {
	//捕获异常 返回 ""
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", R2Bucket.AccountID),
		}, nil
	})
	cfg, err := r2Config.LoadDefaultConfig(context.TODO(),
		r2Config.WithEndpointResolverWithOptions(r2Resolver),
		r2Config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(R2Bucket.AccessKey, R2Bucket.SecretKey, "")),
		r2Config.WithRegion("auto"),
	)
	if err != nil {
		return "", false
	}

	client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(client)

	presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(R2Bucket.Bucket),
		Key:    aws.String(obj),
	})

	if err != nil {
		return "", false
	}

	fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
	return presignResult.URL, true
}
