package {{{ .PackageName }}}

import (
    "errors"
    "github.com/jinzhu/copier"
    "goal-app/model"
    "goal-app/pkg/log"
    "goal-app/pkg/render"
)

type I{{{ .EntityName }}}Service interface {
	List(req *Req{{{ .EntityName }}}List) (res []*Resp{{{ .EntityName }}}Item, total int64, code int, err error)
	Detail(req *Req{{{ .EntityName }}}Detail) (res *Resp{{{ .EntityName }}}Item, code int, err error)
	Create(payload *Req{{{ .EntityName }}}Create) (res *Resp{{{ .EntityName }}}Item, code int, err error)
	Update(payload *Req{{{ .EntityName }}}Update) (res *Resp{{{ .EntityName }}}Item, code int, err error)
	Delete(payload *Req{{{ .EntityName }}}Delete) (code int, e error)
	Tree(req *Req{{{ .EntityName }}}Tree) (res *Resp{{{ .EntityName }}}Tree, code int, err error)
	GetAll{{{ .EntityName }}}() ([]*model.{{{ .EntityName }}}, error)
    Convert{{{ .EntityName }}}TreeToJSON(root *model.{{{ .EntityName }}}, parent *model.{{{ .EntityName }}}) *Resp{{{ .EntityName }}}Tree
}

// {{{ lowerFirst .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ lowerFirst .EntityName }}}Service struct {}

// New{{{ .EntityName }}}Service 初始化
func New{{{ .EntityName }}}Service() I{{{ .EntityName }}}Service {
	return &{{{ lowerFirst .EntityName }}}Service{}
}


// List {{{ .FunctionName }}}列表
func (s *{{{ lowerFirst .EntityName }}}Service) List(req *Req{{{ .EntityName }}}List) (res []*Resp{{{ .EntityName }}}Item, total int64, code int, err error) {
	// 分页信息
	limit := req.Page
	offset := req.Limit * (req.Page - 1)
	// 查询
	query := model.GetDB().Model(&model.{{{ .EntityName }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
	{{{- if and (eq .GoType "string") (eq $queryOpr "like") }}}
	if req.{{{ title (toCamelCase .ColumnName) }}} != "" {
        query = query.Where("{{{ .ColumnName }}} like ?", "%" + req.{{{ title (toCamelCase .ColumnName) }}} + "%")
    }
    {{{- else }}}
    if req.{{{ title (toCamelCase .ColumnName) }}} {{{ if eq .GoType "string" }}}!= ""{{{ else }}}>=0{{{ end }}} {
        query = query.Where("{{{ .ColumnName }}} = ?", req.{{{ title (toCamelCase .ColumnName) }}})
    }
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
	{{{- if contains .AllFields "is_delete" }}}
	query = query.Where("is_delete = ?", 0)
	{{{- end }}}
	// 总数
	err = query.Count(&total).Error
    if err != nil {
        log.GetLogger().Error(err)
        return nil, total, render.QueryError, err
    }
	// 数据
	var objs []*model.{{{ .EntityName }}}
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if err != nil {
	    log.GetLogger().Error(err)
		return nil, total, render.QueryError, err
	}
	err = copier.Copy(&res, objs)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, total, render.DBAttributesCopyError, err
    }
    return res, total, render.OK, nil
}

// Detail {{{ .FunctionName }}}详情
func (s *{{{ lowerFirst .EntityName }}}Service) Detail(req *Req{{{ .EntityName }}}Detail) (res *Resp{{{ .EntityName }}}Item, code int, err error) {
	// var obj *model.{{{ .EntityName }}}
	obj, err := model.Get{{{ .EntityName }}}Instance(
	    model.GetDB(),
	    map[string]interface{}{
	        "{{{ $.PrimaryKey }}}": req.ID,
	        {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
	    },
    )
	if err != nil {
	    if errors.Is(err, gorm.ErrRecordNotFound) {
	        return nil, render.DataNotExistError, err
	    }
		return nil, render.QueryError, err
	}
	err = copier.Copy(&res, obj)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    {{{- range .Columns }}}
    {{{- if and .IsEdit (contains (slice "image" "avatar" "logo" "img") .GoField) }}}
    res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
    {{{- end }}}
    {{{- end }}}
	return
}

// Create {{{ .FunctionName }}}创建
func (s *{{{ lowerFirst .EntityName }}}Service) Create(payload *Req{{{ .EntityName }}}Create) (res *Resp{{{ .EntityName }}}Item, code int, err error) {
	obj := model.New{{{ .EntityName }}}()
    err = copier.Copy(&obj, &payload)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    obj, err = model.Create{{{ .EntityName }}}(model.GetDB(), obj)
    if err != nil {
        return nil, render.CreateError, err
    }
    err = copier.Copy(&res, obj)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    return res, render.OK, nil
}

// Update {{{ .FunctionName }}}更新
func (s *{{{ lowerFirst .EntityName }}}Service) Update(payload *Req{{{ .EntityName }}}Update) (res *Resp{{{ .EntityName }}}Item, code int, err error) {
	obj, err := model.Get{{{ .EntityName }}}Instance(
        model.GetDB(),
        map[string]interface{}{
            "{{{ $.PrimaryKey }}}": payload.ID,
            {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
        },
    )
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, render.DataNotExistError, err
        }
        return nil, render.QueryError, err
    }
    // 更新
    err = copier.Copy(&obj, &payload)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    obj.UpdateTime = time.Now().Unix()
    err = model.Update{{{ .EntityName }}}(model.GetDB(), obj)
    if err != nil {
        return nil, render.UpdateError, err
    }
    err = copier.Copy(&res, &obj)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    return res, render.OK, nil
}

// Delete {{{ .FunctionName }}}删除
func (s *{{{ lowerFirst .EntityName }}}Service) Delete(payload *Req{{{ .EntityName }}}Delete) (code int, e error) {
	_, err := model.Get{{{ .EntityName }}}Instance(
        model.GetDB(),
        map[string]interface{}{
            "{{{ $.PrimaryKey }}}": payload.ID,
            {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
        },
    )
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return render.DataNotExistError, err
        }
        return render.QueryError, err
    }
    // 删除
    err = model.Delete{{{ .EntityName }}}(model.GetDB(), payload.ID)
    if err != nil {
        return render.DeleteError, err
    }
    return render.OK, nil
}

// Tree {{{ .FunctionName }}}树
func (s *{{{ lowerFirst .EntityName }}}Service) Tree(req *Req{{{ .EntityName }}}Tree) (res *Resp{{{ .EntityName }}}Tree, code int, err error) {
	rows, err := s.GetAll{{{ .EntityName }}}()
    if err != nil {
        return nil, render.QueryError, err
    }
    root := model.Build{{{ .EntityName }}}Tree(rows)
    result := s.Convert{{{ .EntityName }}}TreeToJSON(root, nil)
	return result, render.OK, nil
}

// GetAll{{{ .EntityName }}} {{{ .FunctionName }}}获取所有有效数据
func (s *{{{ lowerFirst .EntityName }}}Service) GetAll{{{ .EntityName }}}() ([]*model.{{{ .EntityName }}}, error) {
	result, err := model.GetAll{{{ .EntityName }}}(model.GetDB())
	if err != nil {
		return nil, err
	}
	return result, err
}

// Convert{{{ .EntityName }}}TreeToJSON 行模型数据转成JSON树结构
func (s *{{{ lowerFirst .EntityName }}}Service) Convert{{{ .EntityName }}}TreeToJSON(root *model.{{{ .EntityName }}}, parent *model.{{{ .EntityName }}}) *Resp{{{ .EntityName }}}Tree {
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	result := &Resp{{{ .EntityName }}}Tree{
		ID:         root.ID,
		ParentID:   root.ParentID,
		ParentName: pName,
		Children:   make([]*Resp{{{ .EntityName }}}Tree, len(root.Children)),
	}

	for i, child := range root.Children {
		result.Children[i] = s.Convert{{{ .EntityName }}}TreeToJSON(child, root)
	}
	return result
}
