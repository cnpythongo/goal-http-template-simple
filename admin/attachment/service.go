package attachment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/config"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"goal-app/pkg/storage"
	"goal-app/pkg/utils"
	"gorm.io/gorm"
	"mime/multipart"
	"path"
)

type IAttachmentService interface {
	Create(c *gin.Context, payload *ReqAttachmentCreate, file *multipart.FileHeader) (*RespAttachment, int, error)
}

type attachmentService struct {
	ctx *gin.Context
}

func NewAttachmentService() IAttachmentService {
	return &attachmentService{}
}

func (svc *attachmentService) Create(c *gin.Context, payload *ReqAttachmentCreate, file *multipart.FileHeader) (*RespAttachment, int, error) {
	res := &RespAttachment{}
	driver, err := storage.NewStorageDriver(config.Cfg.Storage.Driver)
	if err != nil {
		return nil, render.CustomerError, err
	}

	folder := path.Join("attachments", payload.Biz)
	upFile, err := driver.Upload(file, folder)
	if err != nil {
		log.GetLogger().Errorf("attachmentService.Create driver.Upload err=%v", err)
		return nil, render.CustomerError, err
	}

	uid := payload.UserId
	record, err := model.GetAttachmentByUserIdAndMd5(model.GetDB(), uid, upFile.Md5)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.GetLogger().Errorf("attachmentService.Add.GetAttachmentByUserIdAndMd5 err=%v", err)
		return res, render.QueryError, err
	}

	var att model.Attachment
	if record != nil {
		att = *record
		_ = driver.RemoveFile(upFile.Uri)
	} else {
		att = model.Attachment{
			UserId: uid,
			UUID:   utils.UUID(),
			Name:   upFile.Name,
			Ext:    upFile.Ext,
			Size:   upFile.Size,
			Md5:    upFile.Md5,
			Path:   upFile.Uri,
			IP:     payload.IP,
		}
		err = model.GetDB().Create(&att).Error
		if err != nil {
			return res, render.CreateError, errors.New("创建上传附件记录失败")
		}
	}

	err = copier.Copy(&res, &att)
	if err != nil {
		return res, render.DBAttributesCopyError, err
	}

	return res, render.OK, nil
}
