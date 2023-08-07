package upload

import (
	"fmt"
	"github.com/Godvictory/douyin/internal/conf"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func Aliyun(uploadName string, file io.Reader) (string, error) {
	AliyunAccessKeyId := conf.Conf.Oss.AccessKeyID
	AliyunAccessKeySecret := conf.Conf.Oss.AccessKeySecret
	AliyunEndpoint := conf.Conf.Oss.Endpoint
	AliyunBucketName := conf.Conf.Oss.BucketName

	client, err := oss.New(AliyunEndpoint, AliyunAccessKeyId, AliyunAccessKeySecret)
	if err != nil {
		return "", err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(AliyunBucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(uploadName, file)
	if err != nil {
		return "", err
	}
	// 拼接链接,默认使用https
	return fmt.Sprintf("https://%s.%s/", AliyunBucketName, AliyunEndpoint), nil
}
