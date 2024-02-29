package types

type (
	// Pagination 分页参数基础结构
	Pagination struct {
		Page      int      `form:"page" default:"1" example:"1"`    // 页码
		Limit     int      `form:"limit" default:"10" example:"10"` // 每页数量
		CreatedAt []string `form:"created_at[]"`                    // 数据创建时间起止区间
	}

	RespPageJson struct {
		Page   int         `json:"page"`
		Limit  int         `json:"limit"`
		Total  int64       `json:"total"`
		Result interface{} `json:"result"`
	}
)
