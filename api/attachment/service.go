package attachment

import (
	"errors"
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"goal-app/pkg/storage"
	"gorm.io/gorm"
	"mime/multipart"
	"path"
)

type IAttachmentService interface {
	Add(payload AttachmentAddReq, file *multipart.FileHeader) (AttachmentResp, int, error)
	uploadFile(file *multipart.FileHeader, folder string) (*storage.UploadFile, int, error)
}

type attachmentService struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewAttachmentService() IAttachmentService {
	db := model.GetDB()
	return &attachmentService{
		db: db,
	}
}

func (svc *attachmentService) Add(payload AttachmentAddReq, file *multipart.FileHeader) (AttachmentResp, int, error) {
	res := AttachmentResp{}
	upFile, code, err := svc.uploadFile(file, path.Join("attachments", payload.UserUuid))
	if err != nil {
		log.GetLogger().Errorf("attachmentService.Add.uploadFile err=%v", err)
		return res, code, err
	}
	uid := payload.UserId

	record, err := model.GetAttachmentByUserIdAndMd5(svc.db, uid, upFile.Md5)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.GetLogger().Errorf("attachmentService.Add.GetAttachmentByUserIdAndMd5 err=%v", err)
		return res, render.QueryError, err
	}

	var att model.Attachment
	if record != nil {
		att = *record
		_ = storage.StorageDriver.RemoveFile(upFile.Uri)
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
		err = svc.db.Create(&att).Error
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

func (svc *attachmentService) uploadFile(file *multipart.FileHeader, folder string) (*storage.UploadFile, int, error) {
	res, err := storage.StorageDriver.Upload(file, folder)
	if err != nil {
		return res, render.UploadFileError, err
	}
	return res, render.OK, nil
}
