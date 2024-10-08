package model

import (
	"errors"
	"goal-app/pkg/log"
	"gorm.io/gorm"
	"strings"
	"time"
)

// GenTable 代码生成业务实体
type GenTable struct {
	BaseModel
	Name         string `json:"name" gorm:"column:name;not null;default:'';comment:表名称"`
	TableComment string `json:"table_comment" gorm:"column:table_comment;not null;default:'';comment:表描述"`
	SubTableName string `json:"sub_table_name" gorm:"column:sub_table_name;not null;default:'';comment:关联表名称"`
	SubTableFk   string `json:"sub_table_fk" gorm:"column:sub_table_fk;not null;default:'';comment:关联表外键"`
	AuthorName   string `json:"author_name" gorm:"column:author_name;not null;default:'';comment:作者的名称"`
	EntityName   string `json:"entity_name" gorm:"column:entity_name;not null;default:'';comment:实体的名称"`
	ModuleName   string `json:"module_name" gorm:"column:module_name;not null;default:'';comment:生成模块名"`
	FunctionName string `json:"function_name" gorm:"column:function_name;not null;default:'';comment:生成功能名"`
	TreePrimary  string `json:"tree_primary" gorm:"column:tree_primary;not null;default:'';comment:树主键字段"`
	TreeParent   string `json:"tree_parent" gorm:"column:tree_parent;not null;default:'';comment:树父级字段"`
	TreeName     string `json:"tree_name" gorm:"column:tree_name;not null;default:'';comment:树显示字段"`
	GenTpl       string `json:"gen_tpl" gorm:"column:gen_tpl;not null;default:'crud';comment:生成模板方式: [crud=单表, tree=树表]"`
	GenType      int    `json:"gen_type" gorm:"column:gen_type;not null;default:0;comment:生成代码方式: [0=zip压缩包, 1=自定义路径]"`
	GenPath      string `json:"gen_path" gorm:"column:gen_path;not null;default:'/';comment:生成代码路径: [不填默认项目路径]"`
	Remarks      string `json:"remarks" gorm:"column:remarks;not null;default:'';comment:备注信息"`
}

func (g *GenTable) TableName() string {
	return "gen_tables"
}

func NewGenTable() *GenTable {
	return &GenTable{}
}

func NewGenTableList() []*GenTable {
	return make([]*GenTable, 0)
}

// GenTableColumn 代码生成表列实体
type GenTableColumn struct {
	BaseModel
	GenTableID    int64  `json:"gen_table_id" gorm:"column:gen_table_id;not null;default:0;comment:表外键"`
	ColumnName    string `json:"column_name" gorm:"column:column_name;not null;default:'';comment:列名称"`
	ColumnComment string `json:"column_comment" gorm:"column:column_comment;not null;default:'';comment:列描述"`
	ColumnLength  int64  `json:"column_length" gorm:"column:column_length;not null;default:0;comment:列长度"`
	ColumnType    string `json:"column_type" gorm:"column:column_type;not null;default:'';comment:列类型"`
	GoType        string `json:"go_type" gorm:"column:go_type;not null;default:'';comment:类型"`
	GoField       string `json:"go_field" gorm:"column:go_field;not null;default:'';comment:字段名"`
	IsPk          uint8  `json:"is_pk" gorm:"column:is_pk;not null;default:0;comment:是否主键: [1=是, 0=否]"`
	IsIncrement   uint8  `json:"is_increment" gorm:"column:is_increment;not null;default:0;comment:是否自增: [1=是, 0=否]"`
	IsRequired    uint8  `json:"is_required" gorm:"column:is_required;not null;default:0;comment:是否必填: [1=是, 0=否]"`
	IsInsert      uint8  `json:"is_insert" gorm:"column:is_insert;not null;default:0;comment:是否为插入字段: [1=是, 0=否]"`
	IsEdit        uint8  `json:"is_edit" gorm:"column:is_edit;not null;default:0;comment:是否编辑字段: [1=是, 0=否]"`
	IsList        uint8  `json:"is_list" gorm:"column:is_list;not null;default:0;comment:是否列表字段: [1=是, 0=否]"`
	IsQuery       uint8  `json:"is_query" gorm:"column:is_query;not null;default:0;comment:是否查询字段: [1=是, 0=否]"`
	QueryType     string `json:"query_type" gorm:"column:query_type;not null;default:'=';comment:查询方式: [等于、不等于、大于、小于、范围]"`
	HtmlType      string `json:"html_type" gorm:"column:html_type;not null;default:'';comment:显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]"`
	DictType      string `json:"dict_type" gorm:"column:dict_type;not null;default:'';comment:字典类型"`
	Sort          int    `json:"sort" gorm:"column:sort;not null;default:0;comment:排序编号"`
}

func (g *GenTableColumn) TableName() string {
	return "gen_table_columns"
}

func NewGenTableColumn() *GenTableColumn {
	return &GenTableColumn{}
}

func NewGenTableColumnList() []*GenTableColumn {
	return make([]*GenTableColumn, 0)
}

func GetGenTableList(db *gorm.DB, page, size int, query []string, args []interface{}) ([]*GenTable, int64, error) {
	qs := db.Model(NewGenTable())
	if query != nil && args != nil && len(args) > 0 {
		queryStr := strings.Join(query, " AND ")
		qs = qs.Where(queryStr, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.generator.GetGenTableList Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := NewGenTableList()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.generator.GetGenTableList Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetGenTableInstance(tx *gorm.DB, conditions map[string]interface{}) (*GenTable, error) {
	result := NewGenTable()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.generate.GetGenTableInstance conditions Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

func UpdateGenTable(db *gorm.DB, table *GenTable) error {
	table.UpdateTime = int64(time.Now().Unix())
	return db.Save(&table).Error
}

func GetGenTableColumnList(db *gorm.DB, tableId int64) ([]*GenTableColumn, error) {
	result := NewGenTableColumnList()
	err := db.Where("gen_table_id = ?", tableId).Order("sort").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetGenTableColumnInstance(db *gorm.DB, id int64) (*GenTableColumn, error) {
	result := NewGenTableColumn()
	err := db.Where("id = ?", id).Take(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateGenTableColumn(db *gorm.DB, column *GenTableColumn) error {
	column.UpdateTime = int64(time.Now().Unix())
	return db.Save(&column).Error
}
