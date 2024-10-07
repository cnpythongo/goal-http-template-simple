package systemmenu

type (
	// ReqSystemMenuCreate 创建菜单数据
	ReqSystemMenuCreate struct {
		ParentID  uint64 `json:"parent_id" binding:"gte=0"` // 父ID
		Name      string `json:"name" binding:"required"`   // 菜单名称
		Kind      string `json:"kind" binding:"required"`   // 菜单类别
		Icon      string `json:"icon"`                      // 菜单图标
		Sort      uint16 `json:"sort"`                      // 菜单排序
		AuthTag   string `json:"auth_tag"`                  // 权限标识
		Route     string `json:"route"`                     // 路由地址
		Component string `json:"component"`                 // 前端组件
		Params    string `json:"params"`                    // 路由参数
		Status    string `json:"status"`                    // 状态: disable=停用, enable=启用
	}

	// ReqSystemMenuUpdate 更新菜单数据
	ReqSystemMenuUpdate struct {
		ID        uint64 `json:"id" binding:"required,gte=0"` // 菜单ID
		ParentID  uint64 `json:"parent_id" binding:"gte=0"`   // 父ID
		Name      string `json:"name" binding:"required"`     // 菜单名称
		Kind      string `json:"kind" binding:"required"`     // 菜单类别
		Icon      string `json:"icon"`                        // 菜单图标
		Sort      uint16 `json:"sort"`                        // 菜单排序
		AuthTag   string `json:"auth_tag"`                    // 权限标识
		Route     string `json:"route"`                       // 路由地址
		Component string `json:"component"`                   // 前端组件
		Params    string `json:"params"`                      // 路由参数
		Status    string `json:"status"`                      // 状态: disable=停用, enable=启用
	}

	// ReqSystemMenuId 菜单ID
	ReqSystemMenuId struct {
		ID uint64 `json:"id" binding:"required"` // 菜单ID
	}

	ReqSystemMenuIds struct {
		IDs []uint64 `json:"ids" binding:"required"` // 菜单ID
	}

	// RespSystemMenuDetail 菜单详情
	RespSystemMenuDetail struct {
		ID         uint64 `json:"id"`          // 菜单ID
		ParentName string `json:"parent_name"` // 父名称
		ParentID   uint64 `json:"parent_id"`   // 父ID
		Name       string `json:"name"`        // 菜单名称
		Kind       string `json:"kind"`        // 菜单类别
		Icon       string `json:"icon"`        // 菜单图标
		Sort       uint16 `json:"sort"`        // 菜单排序
		AuthTag    string `json:"auth_tag"`    // 权限标识
		Route      string `json:"route"`       // 路由地址
		Component  string `json:"component"`   // 前端组件
		Params     string `json:"params"`      // 路由参数
		Status     string `json:"status"`      // 状态: disable=停用, enable=启用
	}

	// RespSystemMenuTree 菜单树结构数据
	RespSystemMenuTree struct {
		ID         uint64 `json:"id"`          // 菜单ID
		ParentName string `json:"parent_name"` // 父名称
		ParentID   uint64 `json:"parent_id"`   // 父ID
		Name       string `json:"name"`        // 菜单名称
		Kind       string `json:"kind"`        // 菜单类别
		Icon       string `json:"icon"`        // 菜单图标
		Sort       uint16 `json:"sort"`        // 菜单排序
		AuthTag    string `json:"auth_tag"`    // 权限标识
		Route      string `json:"route"`       // 路由地址
		Component  string `json:"component"`   // 前端组件
		Params     string `json:"params"`      // 路由参数
		Status     string `json:"status"`      // 状态: disable=停用, enable=启用

		Children []*RespSystemMenuTree `json:"children"` // 子节点
	}
)
