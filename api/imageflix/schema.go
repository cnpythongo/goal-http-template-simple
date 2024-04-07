package imageflix

type (
	UserCreditResp struct {
		Usable int64 `json:"usable"` // 可用点数
	}

	CreditReduceReq struct {
		Point int64 `json:"point" binding:"required"` // 使用点数
	}

	JobCreateReq struct {
	}

	JobStartReq struct {
		JobId int64 `json:"job_id" binding:"required"` // 任务ID
	}
)
