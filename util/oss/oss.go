package oss

import (
	"bytes"
	"os"
	"tfserver/config"
	"tfserver/util/log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

/*
上传文件到oss
file：文件流
path：文件相对地址
*/
func Upload(file []byte, path string) error {
	// 创建OSSClient实例。
	client, err := oss.New(config.OssEndpoint, config.OssAccessKeyId, config.OssAccessKeySecret)
	if err != nil {
		log.ErrorLog("Create oss client failed!", err.Error())
		os.Exit(-1)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.OssBucketName)
	if err != nil {
		log.ErrorLog("Get oss bucket failed!", err.Error())
		os.Exit(-1)
		return err
	}

	// 上传文件流。
	err = bucket.PutObject(path, bytes.NewReader([]byte(file)))
	if err != nil {
		log.ErrorLog("Put file to oss failed!", err.Error())
		os.Exit(-1)
		return err
	}

	return err
}
