package util

import (
	"goal-app/model"
	"goal-app/pkg/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// genUtil 代码生成工具
type genUtil struct{}

var GenUtil = &genUtil{}

// GetDbTables 获取数据中的数据表信息，排除 gen_ 开头的表， 排除已经导入到gen_tables里的表名
func (g *genUtil) GetDbTables(db *gorm.DB, tableName string, tableComment string) *gorm.DB {
	var where = ""
	if tableName != "" {
		where += `and lower(table_name) like lower("%` + tableName + `%")`
	}
	if tableComment != "" {
		where += `and lower(table_comment) like lower("%` + tableComment + `%")`
	}
	query := db.Table("information_schema.tables").Where(
		`table_schema = (SELECT database()) 
			AND table_name NOT LIKE "gen_%" 
			AND table_name NOT IN (select name from gen_tables) ` + where,
	).Select(
		"table_name as name, table_comment, unix_timestamp(create_time) as create_time, unix_timestamp(update_time)as update_time",
	)
	return query
}

// GetDbTablesByName 根据表名集查询表
func (g *genUtil) GetDbTablesByName(db *gorm.DB, tableNames []string) *gorm.DB {
	query := db.Table("information_schema.tables").Where(
		`table_schema = (SELECT database()) 
			AND table_name NOT LIKE "qrtz_%" 
			AND table_name NOT LIKE "gen_%" 
			AND table_name IN ?`, tableNames,
	).Select("table_name as name, table_comment, unix_timestamp(create_time) as create_time, unix_timestamp(update_time)as update_time")
	return query
}

// GetDbTableColumnsByTableName 根据表名查询列信息
func (g *genUtil) GetDbTableColumnsByTableName(db *gorm.DB, tableName string) *gorm.DB {
	query := db.Table("information_schema.columns").Where(
		`table_schema = (SELECT database()) 
			AND table_name = ?`, tableName).Order("ordinal_position").Select(
		`column_name, 
			(CASE WHEN (is_nullable = "no" && column_key != "PRI") THEN "1" ELSE NULL END) AS is_required,
			(CASE WHEN column_key = "PRI" THEN "1" ELSE "0" END) AS is_pk,
			ordinal_position AS sort, column_comment,
			(CASE WHEN extra = "auto_increment" THEN "1" ELSE "0" END) AS is_increment, column_type`)
	return query
}

// CreateGenTable 初始化表
func (g *genUtil) CreateGenTable(table *model.GenTable) model.GenTable {
	return model.GenTable{
		Name:         table.Name,
		TableComment: table.TableComment,
		AuthorName:   "",
		EntityName:   g.ToClassName(table.Name),
		ModuleName:   g.ToModuleName(table.Name),
		FunctionName: strings.Replace(table.TableComment, "表", "", -1),
	}
}

// CreateGenColumn 初始化字段列
func (g *genUtil) CreateGenColumn(tableId int64, column *model.GenTableColumn) *model.GenTableColumn {
	columnType := g.GetDbType(column.ColumnType)
	columnLen := g.GetColumnLength(column.ColumnType)
	col := &model.GenTableColumn{
		GenTableID:    tableId,
		ColumnName:    column.ColumnName,
		ColumnComment: column.ColumnComment,
		ColumnType:    columnType,
		ColumnLength:  columnLen,
		GoField:       column.ColumnName,
		GoType:        GoConstants.TypeString,
		QueryType:     GenConstants.QueryEq,
		Sort:          column.Sort,
		IsPk:          column.IsPk,
		IsIncrement:   column.IsIncrement,
		IsRequired:    column.IsRequired,
	}
	if utils.ToolsUtil.Contains(append(SqlConstants.ColumnTypeStr, SqlConstants.ColumnTypeText...), columnType) {
		//文本域组
		if columnLen >= 500 || utils.ToolsUtil.Contains(SqlConstants.ColumnTypeText, columnType) {
			col.HtmlType = HtmlConstants.HtmlTextarea
		} else {
			col.HtmlType = HtmlConstants.HtmlInput
		}
	} else if utils.ToolsUtil.Contains(SqlConstants.ColumnTypeTime, columnType) {
		//日期字段
		col.GoType = GoConstants.TypeDateTime
		col.HtmlType = HtmlConstants.HtmlDatetime
	} else if utils.ToolsUtil.Contains(SqlConstants.ColumnTimeName, col.ColumnName) {
		//时间字段
		col.GoType = GoConstants.TypeDateTime
		col.HtmlType = HtmlConstants.HtmlDatetime
	} else if utils.ToolsUtil.Contains(SqlConstants.ColumnTypeNumber, columnType) {
		//数字字段
		col.HtmlType = HtmlConstants.HtmlInput
		if strings.Contains(columnType, ",") {
			col.GoType = GoConstants.TypeFloat
		} else {
			col.GoType = GoConstants.TypeInt
		}
	}
	//非必填字段
	if utils.ToolsUtil.Contains(SqlConstants.ColumnNameNotEdit, col.ColumnName) {
		col.IsRequired = 0
	}
	//需插入字段
	if !utils.ToolsUtil.Contains(SqlConstants.ColumnNameNotAdd, col.ColumnName) {
		col.IsInsert = GenConstants.Require
	}
	//需编辑字段
	if !utils.ToolsUtil.Contains(SqlConstants.ColumnNameNotEdit, col.ColumnName) {
		col.IsEdit = GenConstants.Require
		col.IsRequired = GenConstants.Require
	}
	//需列表字段
	if !utils.ToolsUtil.Contains(SqlConstants.ColumnNameNotList, col.ColumnName) && col.IsPk == 0 {
		col.IsList = GenConstants.Require
	}
	//需查询字段
	if !utils.ToolsUtil.Contains(SqlConstants.ColumnNameNotQuery, col.ColumnName) && col.IsPk == 0 {
		col.IsQuery = GenConstants.Require
	}
	lowerColName := strings.ToLower(col.ColumnName)
	//模糊查字段
	if strings.HasSuffix(lowerColName, "name") || utils.ToolsUtil.Contains([]string{"title", "mobile"}, lowerColName) {
		col.QueryType = GenConstants.QueryLike
	}
	//根据字段设置
	if strings.HasSuffix(lowerColName, "status") || utils.ToolsUtil.Contains([]string{"is_show", "is_disable"}, lowerColName) {
		//状态字段设置单选框
		col.HtmlType = HtmlConstants.HtmlRadio
	} else if strings.HasSuffix(lowerColName, "type") || strings.HasSuffix(lowerColName, "sex") {
		//类型&性别字段设置下拉框
		col.HtmlType = HtmlConstants.HtmlSelect
	} else if strings.HasSuffix(lowerColName, "image") {
		//图片字段设置图片上传
		col.HtmlType = HtmlConstants.HtmlImageUpload
	} else if strings.HasSuffix(lowerColName, "file") {
		//文件字段设置文件上传
		col.HtmlType = HtmlConstants.HtmlFileUpload
	} else if strings.HasSuffix(lowerColName, "content") {
		//富文本字段设置富文本编辑器
		col.HtmlType = HtmlConstants.HtmlEditor
	}
	return col
}

// ToModuleName 表名转业务名
func (g *genUtil) ToModuleName(name string) string {
	return strings.ReplaceAll(name, "_", "")
}

// ToClassName 表名转类名
func (g *genUtil) ToClassName(name string) string {
	return utils.ToCamelCase(name)
}

// GetDbType 获取数据库类型字段
func (g *genUtil) GetDbType(columnType string) string {
	index := strings.IndexRune(columnType, '(')
	if index < 0 {
		return columnType
	}
	return columnType[:index]
}

// GetColumnLength 获取字段长度
func (g *genUtil) GetColumnLength(columnType string) int64 {
	index := strings.IndexRune(columnType, '(')
	if index < 0 {
		return 0
	}
	length, err := strconv.Atoi(columnType[index+1 : strings.IndexRune(columnType, ')')])
	if err != nil {
		return 0
	}
	return int64(length)
}

// GetTablePriCol 获取主键列名称
func (g *genUtil) GetTablePriCol(columns []*model.GenTableColumn) (res *model.GenTableColumn) {
	for _, col := range columns {
		if col.IsPk == 1 {
			res = col
			return
		}
	}
	return
}
