package attachment

type (
	AttachmentAddReq struct {
		UserId   uint64 `json:"-" form:"-"` // 用户ID
		UserUuid string `json:"-" form:"-"` // 用户UUID
		IP       string `json:"-" form:"-"` // 用户IP
	}

	AttachmentResp struct {
		UUID       string `json:"uuid"`        // 文件UUID
		Name       string `json:"name"`        // 文件名
		Size       int64  `json:"size"`        // 文件大小, 单位: 字节
		CreateTime int64  `json:"create_time"` // 上传时间(unix秒时间戳)
	}
)
