package storage

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"goal-app/pkg/config"
	"goal-app/pkg/log"
	"goal-app/pkg/utils"
)

// COSStorage 腾讯云COS存储
type COSStorage struct {
	client *cos.Client
	bucket *string
}

// NewCOSStorage 创建新的COS存储实例
func NewCOSStorage() (*COSStorage, error) {
	secretID := config.Cfg.StorageCOS.SecretID
	secretKey := config.Cfg.StorageCOS.SecretKey
	region := config.Cfg.StorageCOS.Region
	bucket := config.Cfg.StorageCOS.Bucket

	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region))
	b := &cos.BaseURL{BucketURL: u}
	auth := &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	}
	client := cos.NewClient(b, auth)

	return &COSStorage{client: client, bucket: &bucket}, nil
}

// Upload COS上传
func (cs *COSStorage) Upload(file *multipart.FileHeader, folder string) (*UploadFile, error) {
	if err := cs.checkFile(file); err != nil {
		return nil, err
	}
	key := cs.buildSaveName(file)
	filePath, err := cs.cosUpload(file, key, folder)
	if err != nil {
		return nil, err
	}

	md5Value, err := cs.MakeFileMD5(filePath)
	if err != nil {
		return nil, err
	}

	fileRelPath := path.Join(folder, key)
	ext := strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1))
	return &UploadFile{
		Name: file.Filename,
		Size: file.Size,
		Ext:  ext,
		Uri:  fileRelPath,
		Path: fmt.Sprintf("https://%s.cos.%s.myqcloud.com/%s", *cs.bucket, config.Cfg.StorageCOS.Region, fileRelPath),
		Md5:  md5Value,
	}, nil
}

// cosUpload COS上传
func (cs *COSStorage) cosUpload(file *multipart.FileHeader, key string, folder string) (string, error) {
	src, err := file.Open()
	if err != nil {
		log.GetLogger().Errorf("COSStorage.cosUpload Open err: err=[%+v]", err)
		return "", errors.New("打开文件失败!")
	}
	defer func(src multipart.File) {
		cErr := src.Close()
		if cErr != nil {
			log.GetLogger().Errorln(err)
		}
	}(src)

	objectKey := path.Join(folder, key)
	_, err = cs.client.Object.Put(context.Background(), objectKey, src, nil)
	if err != nil {
		log.GetLogger().Errorf("COSStorage.cosUpload Put err: key=[%s], err=[%+v]", objectKey, err)
		return "", errors.New("上传文件失败: " + err.Error())
	}
	return objectKey, nil
}

// buildSaveName 生成文件名称
func (cs *COSStorage) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060102")
	return path.Join(date, utils.UUID()+ext)
}

// checkFile 文件验证
func (cs *COSStorage) checkFile(file *multipart.FileHeader) error {
	fileName := file.Filename
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	fileSize := file.Size
	if utils.StrInArrayIndex(fileExt, UploadImageExt) == -1 {
		return errors.New("不支持此图片类型: " + fileExt)
	}
	if fileSize > config.Cfg.Storage.UploadImageSize {
		return errors.New("上传文件应小于: " + strconv.FormatInt(config.Cfg.Storage.UploadImageSize/1024/1024, 10) + "M")
	}
	return nil
}

// RemoveFile 删除文件
func (cs *COSStorage) RemoveFile(filePath string) error {
	_, err := cs.client.Object.Delete(context.Background(), filePath)
	if err != nil {
		log.GetLogger().Errorf("COSStorage RemoveFile err: filePath=[%s], err=[%+v]", filePath, err)
		return errors.New("删除文件失败: " + err.Error())
	}
	return nil
}

// MakeFileMD5 计算文件MD5
func (cs *COSStorage) MakeFileMD5(filePath string) (string, error) {
	resp, err := cs.client.Object.Get(context.Background(), filePath, nil)
	if err != nil {
		log.GetLogger().Errorf("COSStorage MakeFileMD5 Get err: filePath=[%s], err=[%+v]", filePath, err)
		return "", errors.New("获取文件失败: " + err.Error())
	}
	defer func(Body io.ReadCloser) {
		cErr := Body.Close()
		if cErr != nil {
			log.GetLogger().Errorln(cErr)
		}
	}(resp.Body)

	hash := md5.New()
	_, err = io.Copy(hash, resp.Body)
	if err != nil {
		log.GetLogger().Errorf("COSStorage MakeFileMD5 Copy err: filePath=[%s], err=[%+v]", filePath, err)
		return "", errors.New("计算文件MD5失败: " + err.Error())
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
