package systemorg

type (
	// ReqSystemOrgCreate 创建组织机构数据
	ReqSystemOrgCreate struct {
		ParentID uint64 `json:"parent_id" binding:"required,gte=0"` // 父ID
		Name     string `json:"name" binding:"required"`            // 组织名称
		Manager  string `json:"manager"`                            // 负责人名称
		Phone    string `json:"phone"`                              // 负责人电话
	}

	// ReqSystemOrgUpdate 更新组织机构数据
	ReqSystemOrgUpdate struct {
		ID       uint64 `json:"id" binding:"required"`     // 组织ID
		ParentID uint64 `json:"parent_id" binding:"gte=0"` // 父ID
		Name     string `json:"name" binding:"required"`   // 组织名称
		Manager  string `json:"manager"`                   // 负责人名称
		Phone    string `json:"phone"`                     // 负责人电话
	}

	// ReqSystemOrgId 组织机构ID
	ReqSystemOrgId struct {
		ID uint64 `json:"id" binding:"required"` // 组织ID
	}

	ReqSystemOrgIds struct {
		IDs []uint64 `json:"ids" binding:"required"` // 组织ID
	}

	// RespSystemOrgDetail 组织机构详情
	RespSystemOrgDetail struct {
		ID         uint64 `json:"id"`          // 组织ID
		ParentID   uint64 `json:"parent_id"`   // 父ID
		ParentName string `json:"parent_name"` // 父名称
		Name       string `json:"name"`        // 组织名称
		Manager    string `json:"manager"`     // 负责人名称
		Phone      string `json:"phone"`       // 负责人电话
	}

	// RespSystemOrgTree 组织机构树结构数据
	RespSystemOrgTree struct {
		ID         uint64               `json:"id"`          // 组织ID
		ParentID   uint64               `json:"parent_id"`   // 父ID
		ParentName string               `json:"parent_name"` // 父名称
		Name       string               `json:"name"`        // 组织名称
		Manager    string               `json:"manager"`     // 负责人名称
		Phone      string               `json:"phone"`       // 负责人电话
		Children   []*RespSystemOrgTree `json:"children"`    // 子节点
	}
)
