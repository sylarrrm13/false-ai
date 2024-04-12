package service

import (
	"config"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliBucketService struct{}

func (AliBucketService) PreSign(obj string, AliBucket *config.AliBucket) (string, bool) {

	// 创建 OSS 客户端
	client, err := oss.New(AliBucket.EndPoint, AliBucket.AccessKey, AliBucket.SecretKey)
	if err != nil {
		fmt.Println("Error creating OSS client:", err)
		return "", false
	}

	// 获取 bucket
	bucket, err := client.Bucket(AliBucket.Bucket)
	if err != nil {
		fmt.Println("Error getting bucket:", err)
		return "", false
	}

	// 设置预签名 URL 的过期时间

	// 生成预签名 URL
	signedURL, err := bucket.SignURL(obj, oss.HTTPPut, 60*60)
	if err != nil {
		fmt.Println("Error generating signed URL:", err)
		return "", false
	}
	fmt.Printf("%q", signedURL)

	var val = "花间一壶酒，独酌无相亲。 举杯邀明月，对影成三人。"
	err = bucket.PutObjectWithURL(signedURL, strings.NewReader(val))
	if err != nil {
		fmt.Println("Error putting object:", err)
		return "", false

	}
	// 打印预签名 URL
	fmt.Println("Signed URL:", signedURL)
	return signedURL, true
}
