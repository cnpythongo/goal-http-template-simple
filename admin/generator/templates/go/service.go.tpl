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
	Create(payload *Req{{{ .EntityName }}}Create) (*Resp{{{ .EntityName }}}Item, int, error)
	Update(payload *Req{{{ .EntityName }}}Update) (*Resp{{{ .EntityName }}}Item, int, error)
	Delete(payload *Req{{{ .EntityName }}}Delete) (int, error)
	GetAll{{{ .EntityName }}}() ([]*model.{{{ .EntityName }}}, int, error)
	{{{- if eq .GenTpl "tree" }}}
	Tree(req *Req{{{ .EntityName }}}Tree) ([]*Resp{{{ .EntityName }}}Tree, int, error)
    Convert{{{ .EntityName }}}TreeToJSON(nodes []*model.{{{ .EntityName }}}, parent *model.{{{ .EntityName }}}) ([]*Resp{{{ .EntityName }}}Tree, int, error)
    {{{- end }}}
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
	limit := req.Limit
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
	res = &Resp{{{ .EntityName }}}Item{}
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
	err = copier.Copy(&res, &obj)
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
func (s *{{{ lowerFirst .EntityName }}}Service) Create(payload *Req{{{ .EntityName }}}Create) (*Resp{{{ .EntityName }}}Item, int, error) {
	obj := model.New{{{ .EntityName }}}()
    err := copier.Copy(&obj, &payload)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    obj, err = model.Create{{{ .EntityName }}}(model.GetDB(), obj)
    if err != nil {
        return nil, render.CreateError, err
    }
    res := &Resp{{{ .EntityName }}}Item{}
    err = copier.Copy(&res, obj)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    return res, render.OK, nil
}

// Update {{{ .FunctionName }}}更新
func (s *{{{ lowerFirst .EntityName }}}Service) Update(payload *Req{{{ .EntityName }}}Update) (*Resp{{{ .EntityName }}}Item, int, error) {
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
    res := &Resp{{{ .EntityName }}}Item{}
    err = copier.Copy(&res, &obj)
    if err != nil {
        log.GetLogger().Error(err)
        return nil, render.DBAttributesCopyError, err
    }
    return res, render.OK, nil
}

// Delete {{{ .FunctionName }}}删除
func (s *{{{ lowerFirst .EntityName }}}Service) Delete(payload *Req{{{ .EntityName }}}Delete) (int, error) {
    // 删除
    err := model.Delete{{{ .EntityName }}}(model.GetDB(), payload.IDs)
    if err != nil {
        return render.DeleteError, err
    }
    return render.OK, nil
}

// GetAll{{{ .EntityName }}} {{{ .FunctionName }}}获取所有有效数据
func (s *{{{ lowerFirst .EntityName }}}Service) GetAll{{{ .EntityName }}}() ([]*model.{{{ .EntityName }}}, int, error) {
	result, err := model.GetAll{{{ .EntityName }}}(model.GetDB())
	if err != nil {
		return nil, render.QueryError, err
	}
	return result, render.OK, err
}


{{{- if eq .GenTpl "tree" }}}
// Tree {{{ .FunctionName }}}树
func (s *{{{ lowerFirst .EntityName }}}Service) Tree(req *Req{{{ .EntityName }}}Tree) ([]*Resp{{{ .EntityName }}}Tree, int, error) {
	rows, code, err := s.GetAll{{{ .EntityName }}}()
    if err != nil {
        return nil, code, err
    }
    roots := model.Build{{{ .EntityName }}}Tree(rows)
    result, code, err := s.Convert{{{ .EntityName }}}TreeToJSON(roots, nil)
	return result, code, err
}


// Convert{{{ .EntityName }}}TreeToJSON 行模型数据转成JSON树结构
func (s *{{{ lowerFirst .EntityName }}}Service) Convert{{{ .EntityName }}}TreeToJSON(nodes []*model.{{{ .EntityName }}}, parent *model.{{{ .EntityName }}}) ([]*Resp{{{ .EntityName }}}Tree, int, error) {
	result := make([]*Resp{{{ .EntityName }}}Tree, 0)
	if len(nodes) == 0 {
		return result, render.OK, nil
	}
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	for _, node := range nodes {
        children, code, err := s.Convert{{{ .EntityName }}}TreeToJSON(node.Children, node)
        if err != nil {
            return nil, code, err
        }
        item := &Resp{{{ .EntityName }}}Tree{
            ParentName: pName,
            Children:   children,
        }
        err = copier.Copy(&item, &node)
        if err != nil {
            log.GetLogger().Error(err)
            return nil, render.DBAttributesCopyError, err
        }
        result = append(result, item)
    }
    return result, render.OK, nil
}
{{{- end }}}
