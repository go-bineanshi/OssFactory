package qiniu

import (
  "fmt"
  "github.com/go-bineanshi/OssFactory"
  "github.com/qiniu/go-sdk/v7/auth/qbox"
  "github.com/qiniu/go-sdk/v7/storage"
  "sync"
)

type OssClient struct {
  Credentials *qbox.Mac
  Bucket      string
  Endpoint    string
  BucketName  string
}

var lock = &sync.Mutex{}

var qiniuOssInstance *OssClient

func NewQiniu(OssAccessKeyId, OssAccessKeySecret, Endpoint, BucketName string) *OssClient {
  if qiniuOssInstance == nil {
    lock.Lock()
    defer lock.Unlock()
    mac := qbox.NewMac(OssAccessKeyId, OssAccessKeySecret)
    qiniuOssInstance = &OssClient{
      Credentials: mac,
      Endpoint:    Endpoint,
      BucketName:  BucketName,
    }
  }

  return qiniuOssInstance
}

func (oss *OssClient) HandleGetSignParams(filename, fileType string) (OssFactory.UploadCred, error) {
  putPolicy := storage.PutPolicy{
    Scope:      fmt.Sprintf("%s:%s", oss.BucketName, filename),
    ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
  }
  putPolicy.Expires = 7200 //示例2小时有效期
  upToken := putPolicy.UploadToken(oss.Credentials)
  return OssFactory.UploadCred{
    Host: oss.Endpoint,
    Data: struct {
      UpToken string
      Key     string
    }{UpToken: upToken, Key: filename},
  }, nil
}
