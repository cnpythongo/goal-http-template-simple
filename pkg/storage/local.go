package storage

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"goal-app/pkg/config"
	"goal-app/pkg/log"
	"goal-app/pkg/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// LocalStorage 本地磁盘存储
type LocalStorage struct{}

// Upload 本地磁盘上传
func (sd *LocalStorage) Upload(file *multipart.FileHeader, folder string) (uf *UploadFile, e error) {
	if e = sd.checkFile(file); e != nil {
		return
	}
	key := sd.buildSaveName(file)
	md5Value := ""
	filePath, err := sd.localUpload(file, key, folder)
	if err != nil {
		return nil, err
	}

	md5Value, err = sd.MakeFileMD5(filePath)
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
		Path: path.Join(config.Cfg.Storage.PublicPrefix, fileRelPath),
		Md5:  md5Value,
	}, nil
}

// localUpload 本地上传
func (sd *LocalStorage) localUpload(file *multipart.FileHeader, key string, folder string) (string, error) {
	// 保存根目录
	directory := config.Cfg.Storage.UploadDirectory
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		log.GetLogger().Errorf("storageDriver.localUpload Open err: err=[%+v]", err)
		return "", errors.New("打开文件失败!")
	}

	defer func(src multipart.File) {
		cErr := src.Close()
		if cErr != nil {
			log.GetLogger().Errorln(cErr)
		}
	}(src)

	savePath := path.Join(directory, folder, path.Dir(key))
	saveFilePath := path.Join(directory, folder, key)
	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		log.GetLogger().Errorf("storageDriver.localUpload MkdirAll err: path=[%s], err=[%+v]", savePath, err)
		return "", errors.New("创建上传目录失败!")
	}

	out, err := os.Create(saveFilePath)
	if err != nil {
		log.GetLogger().Errorf("storageDriver.localUpload Create err: file=[%s], err=[%+v]", saveFilePath, err)
		return "", errors.New("创建文件失败!")
	}

	defer func(out *os.File) {
		cErr := out.Close()
		if cErr != nil {
			log.GetLogger().Errorln(cErr)
		}
	}(out)

	_, err = io.Copy(out, src)
	if err != nil {
		log.GetLogger().Errorf("storageDriver.localUpload Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		return "", errors.New("上传文件失败: " + err.Error())
	}

	return saveFilePath, nil
}

// checkFile 生成文件名称
func (sd *LocalStorage) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060102")
	return path.Join(date, utils.UUID()+ext)
}

// checkFile 文件验证
func (sd *LocalStorage) checkFile(file *multipart.FileHeader) (e error) {
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

func (sd *LocalStorage) RemoveFile(filePath string) (err error) {
	filePath = path.Join(config.Cfg.Storage.UploadDirectory, filePath)
	if sd.IsFileExist(filePath) {
		err = os.Remove(filePath)
		if err != nil {
			log.GetLogger().Errorf("storageDriver RemoveFile err=%v", err)
		}
	}
	return err
}

func (sd *LocalStorage) MakeFileMD5(filePath string) (string, error) {
	f, err := os.Open(filePath)
	defer func(f *os.File) {
		cErr := f.Close()
		if cErr != nil {
			log.GetLogger().Errorln(cErr)
		}
	}(f)

	if err != nil {
		return "", err
	}

	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		log.GetLogger().Errorf("storageDriver.localUpload MakeFileMD5 err: file=[%s], err=[%+v]", filePath, err)
		return "", errors.New("上传文件MD5计算失败: " + err.Error())
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// IsFileExist 判断文件或目录是否存在
func (sd *LocalStorage) IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
