package attachment

type (
	ReqAttachmentCreate struct {
		UserId int64  `json:"-" form:"-"` // 用户ID
		IP     string `json:"-" form:"-"` // 用户IP
		Biz    string `json:"biz" form:"biz"`
	}

	RespAttachment struct {
		UUID       string `json:"uuid"`        // 文件UUID
		Name       string `json:"name"`        // 文件名
		Size       int64  `json:"size"`        // 文件大小, 单位: 字节
		CreateTime int64  `json:"create_time"` // 上传时间(unix秒时间戳)
	}
)
