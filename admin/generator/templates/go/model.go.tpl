package {{{ .PackageName }}}

import (
	"errors"
    "goal-app/pkg/log"
    "goal-app/pkg/utils"
    "gorm.io/gorm"
    "time"
)

//{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}模型
type {{{ title (toCamelCase .EntityName) }}} struct {
    BaseModel
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

func (m *{{{ title (toCamelCase .EntityName) }}}) TableName() string {
	return "{{{ toSnakeCase .EntityName }}}"
}

func New{{{ title (toCamelCase .EntityName) }}}() *{{{ title (toCamelCase .EntityName) }}} {
	return &{{{ title (toCamelCase .EntityName) }}}{}
}


func New{{{ title (toCamelCase .EntityName) }}}List() []*{{{ title (toCamelCase .EntityName) }}} {
	return make([]*{{{ title (toCamelCase .EntityName) }}}, 0)
}

func (m *{{{ title (toCamelCase .EntityName) }}}) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func Create{{{ title (toCamelCase .EntityName) }}}(tx *gorm.DB, obj *{{{ title (toCamelCase .EntityName) }}}) (obj *{{{ title (toCamelCase .EntityName) }}}, err error) {
    err := tx.Create(&obj).Error
    if err != nil {
        log.GetLogger().Errorf("model.{{{ .FunctionName }}}.Create{{{ title (toCamelCase .EntityName) }}} Error ==> %v", err)
        return nil, err
    }
    return obj, nil
}

func Update{{{ title (toCamelCase .EntityName) }}}(tx *gorm.DB, id int64, data map[string]interface{}) error {
    err := tx.Model(New{{{ title (toCamelCase .EntityName) }}}()).Where("id = ?", id).UpdateColumns(data).Error
    if err != nil {
        log.GetLogger().Errorf("model.{{{ .FunctionName }}}.Update{{{ title (toCamelCase .EntityName) }}} Error ==> %v", err)
    }
    return err
}

func Delete{{{ title (toCamelCase .EntityName) }}}(tx *gorm.DB, id int64) error {
	return Update{{{ title (toCamelCase .EntityName) }}}(tx, id, map[string]interface{}{
		"delete_time": time.Now().Unix(),
	})
}

func Get{{{ title (toCamelCase .EntityName) }}}Instance(tx *gorm.DB, conditions map[string]interface{}) (*{{{ title (toCamelCase .EntityName) }}}, error) {
	result := New{{{ title (toCamelCase .EntityName) }}}()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.{{{ .FunctionName }}}.Get{{{ title (toCamelCase .EntityName) }}}Instance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.{{{ .FunctionName }}}.Get{{{ title (toCamelCase .EntityName) }}}Instance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}


func Get{{{ title (toCamelCase .EntityName) }}}List(tx *gorm.DB, page, size int, query interface{}, args []interface{}) ([]*{{{ title (toCamelCase .EntityName) }}}, int64, error) {
	qs := tx.Model(New{{{ title (toCamelCase .EntityName) }}}()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.{{{ .FunctionName }}}.Get{{{ title (toCamelCase .EntityName) }}}List Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := New{{{ title (toCamelCase .EntityName) }}}List()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.{{{ .FunctionName }}}.Get{{{ title (toCamelCase .EntityName) }}}List Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAll{{{ title (toCamelCase .EntityName) }}}(tx *gorm.DB) ([]*{{{ title (toCamelCase .EntityName) }}}, error) {
	result := New{{{ title (toCamelCase .EntityName) }}}List()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.{{{ .FunctionName }}}.GetAll{{{ title (toCamelCase .EntityName) }}} Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}


func Build{{{ title (toCamelCase .EntityName) }}}Tree(rows []*{{{ title (toCamelCase .EntityName) }}}) *{{{ title (toCamelCase .EntityName) }}} {
	rootNode := New{{{ title (toCamelCase .EntityName) }}}()

	tmpMap := make(map[uint64]*{{{ title (toCamelCase .EntityName) }}})
	for _, r := range rows {
		r.Children = make([]*{{{ title (toCamelCase .EntityName) }}}, 0)
		tmpMap[r.ID] = r
	}

	for _, r := range rows {
		if r.ParentID == 0 {
			rootNode = r
		} else {
			parent, ok := tmpMap[r.ParentID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, r)
			}
		}
	}

	return rootNode
}
