package oss

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"tfserver/config"
	"tfserver/util/log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//上传文件到oss
//file：文件流
//path：文件相对地址
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

//从oss下载文件
func Download(path string) ([]byte, error) {
	var file []byte
	// 创建OSSClient实例。
	client, err := oss.New(config.OssEndpoint, config.OssAccessKeyId, config.OssAccessKeySecret)
	if err != nil {
		log.ErrorLog("Create oss client failed!", err.Error())
		os.Exit(-1)
		return file, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.OssBucketName)
	if err != nil {
		log.ErrorLog("Get oss bucket failed!", err.Error())
		os.Exit(-1)
		return file, err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(path)
	if err != nil {
		log.ErrorLog("Get file failed!", err.Error())
		os.Exit(-1)
		return file, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	file, err = ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
		return file, err
	}

	return file, err
}
