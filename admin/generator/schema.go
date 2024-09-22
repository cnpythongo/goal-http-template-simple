package generator

import "goal-app/pkg/render"

// ReqDbTables 库表列表参数
type ReqDbTables struct {
	render.Pagination
	TableName    string `form:"table_name"`    // 表名称
	TableComment string `form:"table_comment"` // 表描述
}

// ReqListTable 生成列表参数
type ReqListTable struct {
	TableName    string `form:"table_name"`    // 表名称
	TableComment string `form:"table_comment"` // 表描述
	StartTime    uint64 `form:"start_time"`    // 开始时间
	EndTime      uint64 `form:"end_time"`      // 结束时间
}

// ReqDetailTable 生成详情参数
type ReqDetailTable struct {
	ID uint64 `form:"id" binding:"required,gt=0"` // 主键
}

// ReqImportTable 导入表结构参数
type ReqImportTable struct {
	Tables string `json:"tables" binding:"required"` // 导入的表, 用","分隔
}

// ReqSyncTable 同步表结构参数
type ReqSyncTable struct {
	ID uint64 `form:"id" binding:"required,gt=0"` // 主键
}

// ReqEditColumn 表编辑列
type ReqEditColumn struct {
	ID            uint64 `form:"id" binding:"required,gt=0"`                // 主键
	ColumnComment string `form:"column_comment" binding:"required,max=200"` // 列描述
	JavaField     string `form:"go_field" binding:"required,max=100"`       // 字段
	IsRequired    uint8  `form:"is_stop" binding:"oneof=0 1"`               // 是否必填: [0=否, 1=是]
	IsInsert      uint8  `form:"is_insert" binding:"oneof=0 1"`             // 是否新增字段: [0=否, 1=是]
	IsEdit        uint8  `form:"is_edit" binding:"oneof=0 1"`               // 是否编辑字段: [0=否, 1=是]
	IsList        uint8  `form:"is_list" binding:"oneof=0 1"`               // 是否列表字段: [0=否, 1=是]
	IsQuery       uint8  `form:"is_query" binding:"oneof=0 1"`              // 是否查询字段: [0=否, 1=是]
	QueryType     string `form:"query_type" binding:"required,max=30"`      // 查询方式
	HtmlType      string `form:"html_type" binding:"required,max=30"`       // 表单类型
	DictType      string `form:"dict_type" binding:"required,max=200"`      // 字典类型
}

// ReqEditTable 编辑表结构参数
type ReqEditTable struct {
	ID           uint64          `form:"id" binding:"required,gt=0"`                     // 主键
	TableName    string          `form:"table_name" binding:"required,min=1,max=200"`    // 表名称
	EntityName   string          `form:"entity_name" binding:"required,min=1,max=200"`   // 实体名称
	TableComment string          `form:"table_comment" binding:"required,min=1,max=200"` // 表描述
	AuthorName   string          `form:"author_name" binding:"required,min=1,max=100"`   // 作者名称
	Remarks      string          `form:"remarks" binding:"max=60"`                       // 备注信息
	GenTpl       string          `form:"gen_tpl" binding:"oneof=crud tree"`              // 生成模板方式: [crud=单表, tree=树表]
	ModuleName   string          `form:"module_name" binding:"required,min=1,max=60"`    // 生成模块名
	FunctionName string          `form:"function_name" binding:"required,min=1,max=60"`  // 生成功能名
	GenType      int             `form:"gen_type" binding:"oneof=0 1"`                   // 生成代码方式: [0=zip压缩包, 1=自定义路径]
	GenPath      string          `form:"gen_path,default=/" binding:"required,max=60"`   // 生成路径
	TreePrimary  string          `form:"tree_primary"`                                   // 树表主键
	TreeParent   string          `form:"tree_parent"`                                    // 树表父键
	TreeName     string          `form:"tree_name"`                                      // 树表名称
	SubTableName string          `form:"sub_table_name"`                                 // 子表名称
	SubTableFk   string          `form:"sub_table_fk"`                                   // 子表外键
	Columns      []ReqEditColumn `form:"columns" binding:"required"`                     // 字段列表
}

// ReqDelTable 删除表结构参数
type ReqDelTable struct {
	Ids []uint64 `json:"ids" binding:"required"` // 主键
}

// ReqPreviewCode 预览代码参数
type ReqPreviewCode struct {
	ID uint64 `form:"id" binding:"required,gt=0"` // 主键
}

// ReqGenCode 生成代码参数
type ReqGenCode struct {
	Tables string `form:"tables" binding:"required"` // 生成的表, 用","分隔
}

// ReqDownloadCode 下载代码参数
type ReqDownloadCode struct {
	Tables string `form:"tables" binding:"required"` // 下载的表, 用","分隔
}

// RespDbTableItem 数据表单条返回信息
type RespDbTableItem struct {
	TableName    string `json:"table_name"`    // 表名称
	TableComment string `json:"table_comment"` // 表描述
	CreateTime   uint64 `json:"create_time"`   // 创建时间
	UpdateTime   uint64 `json:"update_time"`   // 更新时间
}

type RespDbTableList render.RespPageJson

// RespGenTable 生成表返回信息
type RespGenTable struct {
	ID           uint64 `json:"id"`            // 主键
	GenType      int    `json:"gen_type"`      // 生成类型
	TableName    string `json:"table_name"`    // 表名称
	TableComment string `json:"table_comment"` // 表描述
	CreateTime   uint64 `json:"create_time"`   // 创建时间
	UpdateTime   uint64 `json:"update_time"`   // 更新时间
}

// RespGenTableBase 生成表基本返回信息
type RespGenTableBase struct {
	ID           uint64 `json:"id"`            // 主键
	TableName    string `json:"table_name"`    // 表的名称
	TableComment string `json:"table_comment"` // 表的描述
	EntityName   string `json:"entity_name"`   // 实体名称
	AuthorName   string `json:"author_name"`   // 作者名称
	Remarks      string `json:"remarks"`       // 备注信息
	CreateTime   uint64 `json:"create_time"`   // 创建时间
	UpdateTime   uint64 `json:"update_time"`   // 更新时间
}

// RespGenTableGen 生成表生成返回信息
type RespGenTableGen struct {
	GenTpl       string `json:"gen_tpl"`        // 生成模板方式: [crud=单表, tree=树表]
	GenType      int    `json:"gen_type"`       // 生成代码方式: [0=zip压缩包, 1=自定义路径]
	GenPath      string `json:"gen_path"`       // 生成代码路径: [不填默认项目路径]
	ModuleName   string `json:"module_name"`    // 生成模块名
	FunctionName string `json:"function_name"`  // 生成功能名
	TreePrimary  string `json:"tree_primary"`   // 树主键字段
	TreeParent   string `json:"tree_parent"`    // 树父级字段
	TreeName     string `json:"tree_name"`      // 树显示字段
	SubTableName string `json:"sub_table_name"` // 关联表名称
	SubTableFk   string `json:"sub_table_fk"`   // 关联表外键
}

// RespGenColumn 生成列返回信息
type RespGenColumn struct {
	ID            uint64 `json:"id"`             // 字段主键
	ColumnName    string `json:"column_name"`    // 字段名称
	ColumnComment string `json:"column_comment"` // 字段描述
	ColumnLength  int    `json:"column_length"`  // 字段长度
	ColumnType    string `json:"column_type"`    // 字段类型
	JavaType      string `json:"go_type"`        // Go类型
	JavaField     string `json:"go_field"`       // Go字段
	IsRequired    uint8  `json:"is_required"`    // 是否必填
	IsInsert      uint8  `json:"is_insert"`      // 是否为插入字段
	IsEdit        uint8  `json:"is_edit"`        // 是否编辑字段
	IsList        uint8  `json:"is_list"`        // 是否列表字段
	IsQuery       uint8  `json:"is_query"`       // 是否查询字段
	QueryType     string `json:"query_type"`     // 查询方式: [等于、不等于、大于、小于、范围]
	HtmlType      string `json:"html_type"`      // 显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]
	DictType      string `json:"dict_type"`      // 字典类型
	CreateTime    uint64 `json:"create_time"`    // 创建时间
	UpdateTime    uint64 `json:"update_time"`    // 更新时间
}

// RespGenTableDetail 生成表详情返回信息
type RespGenTableDetail struct {
	Base    RespGenTableBase `json:"base"`    // 基本信息
	Gen     RespGenTableGen  `json:"gen"`     // 生成信息
	Columns []RespGenColumn  `json:"columns"` // 字段列表
}
