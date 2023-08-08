package test

import (
  "fmt"
  "github.com/go-bineanshi/OssFactory"
  "github.com/go-bineanshi/OssFactory/aliyun"
  "github.com/go-bineanshi/OssFactory/qiniu"
  "os"
  "testing"
)

func TestAliyunOss(t *testing.T) {

  config := OssFactory.OssConfig{
    OssAccessKeyId:     os.Getenv("OSS_ACCESS_KEY_ID"),
    OssAccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
    Endpoint:           os.Getenv("OSS_ENDPOINT"),
    BucketName:         os.Getenv("OSS_BUCKET_NAME"),
  }
  client, err := aliyun.NewAliyunOssClient(config.OssAccessKeyId, config.OssAccessKeySecret, config.Endpoint, config.BucketName)
  if err != nil {
    return
  }
  up := OssFactory.UploadFileInfo{
    Filename: "favicon.png",
    FileType: "image/png",
  }
  accept, err := up.Accept(client)
  if err != nil {
    return
  }
  fmt.Println(accept)
}

func TestQiniuOss(t *testing.T) {
  config := OssFactory.OssConfig{
    OssAccessKeyId:     os.Getenv("QINIU_ACCESS_KEY_ID"),
    OssAccessKeySecret: os.Getenv("QINIU_ACCESS_KEY_SECRET"),
    Endpoint:           os.Getenv("QINIU_ENDPOINT"),
    BucketName:         os.Getenv("QINIU_BUCKET_NAME"),
  }
  client := qiniu.NewQiniu(config.OssAccessKeyId, config.OssAccessKeySecret, config.Endpoint, config.BucketName)

  up := OssFactory.UploadFileInfo{
    Filename: "favicon.png",
    FileType: "image/png",
  }
  accept, err := up.Accept(client)
  if err != nil {
    return
  }
  fmt.Println(accept)
}
