package attachment

import "github.com/cnpythongo/goal-tools/utils"

type (
	AttachmentAddReq struct {
		UserId   int64  `json:"-" form:"-"` // 用户ID
		UserUuid string `json:"-" form:"-"` // 用户UUID
		IP       string `json:"-" form:"-"` // 用户IP
	}

	AttachmentResp struct {
		Uuid      string           `json:"uuid"`       // 文件UUID
		Name      string           `json:"name"`       // 文件名
		Size      int64            `json:"size"`       // 文件大小, 单位: 字节
		CreatedAt *utils.LocalTime `json:"created_at"` // 上传时间
	}
)
