package storage

import (
	"errors"
	"mime/multipart"
)

var UploadImageExt = []string{"png", "jpeg", "jpg"}

// UploadFile 文件对象
type UploadFile struct {
	Name string // 文件名称
	Size int64  // 文件大小
	Ext  string // 文件扩展
	Uri  string // 文件路径
	Path string // 访问地址
	Md5  string // 文件MD5值
}

// Storage 接口定义了存储的基本方法
type Storage interface {
	Upload(file *multipart.FileHeader, folder string) (*UploadFile, error)
	RemoveFile(filePath string) error
	MakeFileMD5(filePath string) (string, error)
}

// StorageDriver 存储引擎
type StorageDriver struct {
	storage Storage
}

// NewStorageDriver 创建新的存储引擎实例
func NewStorageDriver(storageType string) (*StorageDriver, error) {
	var storage Storage
	switch storageType {
	case "local":
		storage = &LocalStorage{}
	case "cos":
		cosStorage, err := NewCOSStorage()
		if err != nil {
			return nil, err
		}
		storage = cosStorage
	default:
		return nil, errors.New("不支持的存储类型")
	}
	return &StorageDriver{storage: storage}, nil
}

// Upload 上传文件
func (se *StorageDriver) Upload(file *multipart.FileHeader, folder string) (*UploadFile, error) {
	return se.storage.Upload(file, folder)
}

// RemoveFile 删除文件
func (se *StorageDriver) RemoveFile(filePath string) error {
	return se.storage.RemoveFile(filePath)
}

// MakeFileMD5 计算文件MD5
func (se *StorageDriver) MakeFileMD5(filePath string) (string, error) {
	return se.storage.MakeFileMD5(filePath)
}
