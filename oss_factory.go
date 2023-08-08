package OssFactory

// OssConfig 定义 OSS 配置结构体
type OssConfig struct {
  OssAccessKeyId     string
  OssAccessKeySecret string
  Endpoint           string
  BucketName         string
  Mold               string // aliyunoss  qiniuoss
}

// OSSClient 定义通用的 OSS 客户端接口
type OSSClient interface {
  HandleGetSignParams(filename, filetype string) (UploadCred, error)
}

type UploadFileInfo struct {
  Filename string
  FileType string
  Mold     string
}

type UploadCred struct {
  Host string
  Data struct {
    UpToken string
    Key     string
  }
}

func (up *UploadFileInfo) Accept(ossClient OSSClient) (UploadCred, error) {
  uploadCred, err := ossClient.HandleGetSignParams(up.Filename, up.FileType)
  if err != nil {
    return UploadCred{}, err
  }
  return uploadCred, nil
}
