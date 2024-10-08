package generator

import "goal-app/pkg/render"

// ReqGenTableList 待生成代码的数据表列表参数
type ReqGenTableList struct {
	render.Pagination
	Name         string `form:"name"`          // 表名称
	TableComment string `form:"table_comment"` // 表描述
}

// RespGenTableItem 待生成代码的数据表返回信息
type RespGenTableItem struct {
	ID           int64  `json:"id"`            // 主键
	Name         string `json:"name"`          // 表名称
	TableComment string `json:"table_comment"` // 表描述
	GenTpl       string `json:"gen_tpl"`       // 生成模板方式: [crud=单表, tree=树表]
	GenPath      string `json:"gen_path"`      // 生成代码路径: [不填默认项目路径]
	EntityName   string `json:"entity_name"`   // 实体的名称
	ModuleName   string `json:"module_name"`   // '生成模块名
	FunctionName string `json:"function_name"` // 生成功能名
	TreePrimary  string `json:"tree_primary"`  // 树主键字段
	TreeParent   string `json:"tree_parent"`   // 树父级字段
	TreeName     string `json:"tree_name"`     // 树显示字段
	CreateTime   int64  `json:"create_time"`   // 创建时间
	UpdateTime   int64  `json:"update_time"`   // 更新时间
}

// ReqDbTableList 库表列表参数
type ReqDbTableList struct {
	render.Pagination
	TableName    string `form:"table_name"`    // 表名称
	TableComment string `form:"table_comment"` // 表描述
}

// ReqDetailTable 生成详情参数
type ReqDetailTable struct {
	ID int64 `form:"id" binding:"required,gt=0"` // 主键
}

// ReqGenTableCreate 导入表结构参数
type ReqGenTableCreate struct {
	Tables string `json:"tables" binding:"required"` // 导入的表, 用","分隔
}

// ReqSyncTable 同步表结构参数
type ReqSyncTable struct {
	ID int64 `form:"id" binding:"required,gt=0"` // 主键
}

// ReqUpdateGenTable 编辑表结构参数
type ReqUpdateGenTable struct {
	ID           int64  `json:"id" binding:"required,gt=0"`                     // 主键
	Name         string `json:"name" binding:"required,min=1,max=200"`          // 表名称
	TableComment string `json:"table_comment" binding:"required,min=1,max=200"` // 表描述
	EntityName   string `json:"entity_name" binding:"required,min=1,max=200"`   // 实体名称
	ModuleName   string `json:"module_name" binding:"required,min=1,max=60"`    // 生成模块名
	FunctionName string `json:"function_name" binding:"required,min=1,max=60"`  // 生成功能名
	GenTpl       string `json:"gen_tpl" binding:"oneof=crud tree"`              // 生成模板方式: [crud=单表, tree=树表]
	GenPath      string `json:"gen_path,default=/" binding:"required,max=60"`   // 生成路径
	TreePrimary  string `json:"tree_primary"`                                   // 树表主键
	TreeParent   string `json:"tree_parent"`                                    // 树表父键
	TreeName     string `json:"tree_name"`                                      // 树表名称
	SubTableName string `json:"sub_table_name"`                                 // 子表名称
	SubTableFk   string `json:"sub_table_fk"`                                   // 子表外键
	// AuthorName   string `json:"author_name" binding:"required,min=1,max=100"`   // 作者名称
	// Remarks      string `json:"remarks" binding:"max=60"`                       // 备注信息
	// GenType      int    `json:"gen_type" binding:"oneof=0 1"`                   // 生成代码方式: [0=zip压缩包, 1=自定义路径]
}

// ReqUpdateGenColumn 表编辑列
type ReqUpdateGenColumn struct {
	ID            int64  `json:"id" binding:"required,gt=0"`                // 主键
	ColumnName    string `json:"column_name" binding:"required,max=200"`    // 列描述
	ColumnComment string `json:"column_comment" binding:"required,max=200"` // 列描述
	GoType        string `json:"go_type" binding:"max=100"`                 // 字段
	GoField       string `json:"go_field" binding:"max=100"`                // 字段
	IsRequired    uint8  `json:"is_required" binding:"oneof=0 1"`           // 是否必填: [0=否, 1=是]
	IsInsert      uint8  `json:"is_insert" binding:"oneof=0 1"`             // 是否新增字段: [0=否, 1=是]
	IsEdit        uint8  `json:"is_edit" binding:"oneof=0 1"`               // 是否编辑字段: [0=否, 1=是]
	IsList        uint8  `json:"is_list" binding:"oneof=0 1"`               // 是否列表字段: [0=否, 1=是]
	IsQuery       uint8  `json:"is_query" binding:"oneof=0 1"`              // 是否查询字段: [0=否, 1=是]
	QueryType     string `json:"query_type" binding:"required,max=30"`      // 查询方式
	HtmlType      string `json:"html_type" binding:"required,max=30"`       // 表单类型
	DictType      string `json:"dict_type" binding:"max=200"`               // 字典类型
}

// ReqDelTable 删除表结构参数
type ReqDelTable struct {
	Ids []int64 `json:"ids" binding:"required"` // 主键
}

type ReqDelGenTableColumn ReqDelTable

// ReqPreview 预览代码参数
type ReqPreview struct {
	ID int64 `form:"id" uri:"id" binding:"required,gt=0"` // 主键
}

// RespPreviewItem 预览代码返回值
type RespPreviewItem struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Content  string `json:"content"`
}

// ReqGenCode 生成代码参数
type ReqGenCode struct {
	Tables string `json:"tables" binding:"required"` // 生成的表, 用","分隔
}

// RespDbTable 数据表单条返回信息
type RespDbTable struct {
	Name         string `json:"name"`          // 表名称
	TableComment string `json:"table_comment"` // 表描述
	CreateTime   int64  `json:"create_time"`   // 创建时间
	UpdateTime   int64  `json:"update_time"`   // 更新时间
}

// RespGenColumn 生成列返回信息
type RespGenColumn struct {
	ID            int64  `json:"id"`             // 字段主键
	ColumnName    string `json:"column_name"`    // 字段名称
	ColumnComment string `json:"column_comment"` // 字段描述
	ColumnLength  int    `json:"column_length"`  // 字段长度
	ColumnType    string `json:"column_type"`    // 字段类型
	GoType        string `json:"go_type"`        // Go类型
	GoField       string `json:"go_field"`       // Go字段
	IsRequired    uint8  `json:"is_required"`    // 是否必填
	IsInsert      uint8  `json:"is_insert"`      // 是否为插入字段
	IsEdit        uint8  `json:"is_edit"`        // 是否编辑字段
	IsList        uint8  `json:"is_list"`        // 是否列表字段
	IsQuery       uint8  `json:"is_query"`       // 是否查询字段
	QueryType     string `json:"query_type"`     // 查询方式: [等于、不等于、大于、小于、范围]
	HtmlType      string `json:"html_type"`      // 显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType      string `json:"dict_type"`      // 字典类型
	CreateTime    int64  `json:"create_time"`    // 创建时间
	UpdateTime    int64  `json:"update_time"`    // 更新时间
}
