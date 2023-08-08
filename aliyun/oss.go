package aliyun

import (
  "fmt"
  "github.com/aliyun/aliyun-oss-go-sdk/oss"
  "github.com/go-bineanshi/OssFactory"
  "net/url"
  "os"
  "sync"
)

var lock = &sync.Mutex{}

type OssClient struct {
  Bucket *oss.Bucket
}

var aliYunOssInstance *OssClient

func NewAliyunOssClient(OssAccessKeyId, OssAccessKeySecret, Endpoint, BucketName string) (*OssClient, error) {
  if aliYunOssInstance == nil {
    lock.Lock()
    defer lock.Unlock()
    // 创建OSSClient实例。
    client, err := oss.New(Endpoint, OssAccessKeyId, OssAccessKeySecret)
    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(-1)
    }
    // Bucket 桶
    bucket, err := client.Bucket(BucketName)
    if err != nil {
      panic(err)
    }
    aliYunOssInstance = &OssClient{
      Bucket: bucket,
    }
  }
  return aliYunOssInstance, nil
}

func (c *OssClient) HandleGetSignParams(filename, fileType string) (OssFactory.UploadCred, error) {
  options := []oss.Option{
    oss.ContentType(fileType),
  }
  signedURL, err := c.Bucket.SignURL(filename, oss.HTTPPut, 60, options...)
  if err != nil {
    panic(err)
  }
  enEscapeUrl, err := url.QueryUnescape(signedURL)
  return OssFactory.UploadCred{
    Host: enEscapeUrl,
  }, err
}
