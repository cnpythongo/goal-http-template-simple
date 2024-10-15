package {{{ .PackageName }}}

import (
	"errors"
    "goal-app/pkg/log"
    "goal-app/pkg/utils"
    "gorm.io/gorm"
    "time"
)

//{{{ .EntityName }}} {{{ .FunctionName }}}模型
type {{{ .EntityName }}} struct {
    BaseModel
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- if eq .GenTpl "tree"}}}
    Children []*{{{ .EntityName }}} `gorm:"foreignKey:parent_id;references:id" json:"children,omitempty"`
    {{{- end }}}
}

func (m *{{{ .EntityName }}}) TableName() string {
	return "{{{ .Name }}}"
}

func New{{{ .EntityName }}}() *{{{ .EntityName }}} {
	return &{{{ .EntityName }}}{}
}


func New{{{ .EntityName }}}List() []*{{{ .EntityName }}} {
	return make([]*{{{ .EntityName }}}, 0)
}

func (m *{{{ .EntityName }}}) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Unix()
	m.CreateTime = now
	m.UpdateTime = now
	return nil
}

func Create{{{ .EntityName }}}(tx *gorm.DB, obj *{{{ .EntityName }}}) (*{{{ .EntityName }}}, error) {
    err := tx.Create(&obj).Error
    if err != nil {
        log.GetLogger().Errorf("model.{{{ .EntityName }}}.Create{{{ .EntityName }}} Error ==> %v", err)
        return nil, err
    }
    return obj, nil
}

func Update{{{ .EntityName }}}(tx *gorm.DB, obj *{{{ .EntityName }}}) error {
	err := tx.Save(&obj).Error
	if err != nil {
		log.GetLogger().Errorf("model.{{{ .EntityName }}}.Update{{{ .EntityName }}} Error ==> %v", err)
	}
	return err
}

func Delete{{{ .EntityName }}}(tx *gorm.DB, ids []int64) error {
	err := tx.Model(New{{{ .EntityName }}}()).Where("id in ?", ids).UpdateColumns(map[string]interface{}{
        "delete_time": time.Now().Unix(),
    }).Error
    if err != nil {
        log.GetLogger().Errorf("model.{{{ .EntityName }}}.Delete{{{ .EntityName }}} Error ==> %v", err)
    }
    return err
}

func Get{{{ .EntityName }}}Instance(tx *gorm.DB, conditions map[string]interface{}) (*{{{ .EntityName }}}, error) {
	result := New{{{ .EntityName }}}()
	err := tx.Where(conditions).Take(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Infof("model.{{{ .EntityName }}}.Get{{{ .EntityName }}}Instance conditions ==> %v", conditions)
			log.GetLogger().Errorf("model.{{{ .EntityName }}}.Get{{{ .EntityName }}}Instance Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}


func Get{{{ .EntityName }}}List(tx *gorm.DB, page, limit int, query interface{}, args []interface{}) ([]*{{{ .EntityName }}}, int64, error) {
	qs := tx.Model(New{{{ .EntityName }}}()).Where("delete_time == 0")
	if query != nil && args != nil && len(args) > 0 {
		qs = qs.Where(query, args...)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("model.{{{ .EntityName }}}.Get{{{ .EntityName }}}List Count Error ==> %v", err)
		return nil, 0, err
	}
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		qs = qs.Limit(limit).Offset(offset)
	}
	result := New{{{ .EntityName }}}List()
	err = qs.Find(&result).Error
	if err != nil {
		log.GetLogger().Errorf("model.{{{ .EntityName }}}.Get{{{ .EntityName }}}List Query Error ==> %v", err)
		return nil, 0, err
	}
	return result, total, nil
}

func GetAll{{{ .EntityName }}}(tx *gorm.DB) ([]*{{{ .EntityName }}}, error) {
	result := New{{{ .EntityName }}}List()
	err := tx.Where("delete_time == 0").Find(&result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.GetLogger().Errorf("model.{{{ .EntityName }}}.GetAll{{{ .EntityName }}} Error ==> %v", err)
		}
		return nil, err
	}
	return result, nil
}

{{{- if eq .GenTpl "tree" }}}
func Build{{{ .EntityName }}}Tree(rows []*{{{ .EntityName }}}) []*{{{ .EntityName }}} {
	rootNodes := New{{{ .EntityName }}}List()

	tmpMap := make(map[int64]*{{{ .EntityName }}})
	for _, r := range rows {
		r.Children = make([]*{{{ .EntityName }}}, 0)
		tmpMap[r.ID] = r
	}

	for _, r := range rows {
		if r.ParentID == 0 {
			rootNodes = append(rootNodes, r)
		} else {
			parent, ok := tmpMap[r.ParentID]
			if ok && parent != nil {
				parent.Children = append(parent.Children, r)
			}
		}
	}
	return rootNodes
}
{{{- end }}}