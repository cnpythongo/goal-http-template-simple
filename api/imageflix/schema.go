package imageflix

type (
	UserCreditResp struct {
		Usable int64 `json:"usable"` // 可用点数
	}

	CreditReduceReq struct {
		Point int64 `json:"point" binding:"required"` // 使用点数
	}
)
