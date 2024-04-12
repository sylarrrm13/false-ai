package service

import (
	"config"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentBucketService struct{}

func (TencentBucketService) PreSign(obj string, tencentBucket *config.TencentBucket) (string, bool) {
	u, _ := url.Parse(tencentBucket.EndPoint)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})

	// 生成预签名 URL（有效期为一小时）
	presignedURL, err := c.Object.GetPresignedURL(context.Background(), http.MethodPut, obj, tencentBucket.AccessKey, tencentBucket.SecretKey, time.Hour, nil)
	if err != nil {
		fmt.Println("Error generating presigned URL:", err)
		return "", false
	}

	return presignedURL.String(), true
}
