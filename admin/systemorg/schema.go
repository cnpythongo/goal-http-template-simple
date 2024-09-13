package systemorg

type (
	// ReqSystemOrgCreate 创建组织机构数据
	ReqSystemOrgCreate struct {
		ParentID uint64 `json:"parent_id"` // 父ID
		Name     string `json:"name"`      // 组织名称
	}

	// ReqSystemOrgUpdate 更新组织机构数据
	ReqSystemOrgUpdate struct {
		ID       uint64 `json:"id"`        // 组织ID
		ParentID uint64 `json:"parent_id"` // 父ID
		Name     string `json:"name"`      // 组织名称
	}

	// ReqSystemOrgId 组织机构ID
	ReqSystemOrgId struct {
		ID uint64 `json:"id"` // 组织ID
	}

	// RespSystemOrgTree 组织机构树结构数据
	RespSystemOrgTree struct {
		ID         uint64               `json:"id"`          // 组织ID
		ParentID   uint64               `json:"parent_id"`   // 父ID
		ParentName string               `json:"parent_name"` // 父名称
		Name       string               `json:"name"`        // 组织名称
		Children   []*RespSystemOrgTree `json:"children"`    // 子节点
	}
)
