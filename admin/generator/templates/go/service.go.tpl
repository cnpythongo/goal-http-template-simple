package {{{ .PackageName }}}

import (
    "github.com/jinzhu/copier"
    "goal-app/model"
    "goal-app/pkg/log"
    "goal-app/pkg/render"
)

type I{{{ .EntityName }}}Service interface {
	List(req *{{{ .EntityName }}}ListReq) (res []*{{{ .EntityName }}}ItemResp, total int64, code int, err error)
	Detail(req *{{{ .EntityName }}}DetailReq) (res *{{{ .EntityName }}}ItemResp, code int, err error)
	Create(payload *{{{ .EntityName }}}CreateReq) (res *{{{ .EntityName }}}ItemResp, code int, err error)
	Update(payload *{{{ .EntityName }}}UpdateReq) (res *{{{ .EntityName }}}ItemResp, code int, err error)
	Delete(payload *{{{ .EntityName }}}DeleteReq) (code int, e error)
	Tree(req *{{{ .EntityName }}}TreeReq) (res *{{{ .EntityName }}}TreeResp, code int, err error)
	GetAll{{{ title (toCamelCase .EntityName) }}}() ([]*model.{{{ title (toCamelCase .EntityName) }}}, error)
    Build{{{ title (toCamelCase .EntityName) }}}Tree(rows []*{{{ toCamelCaseWithoutFirst .EntityName }}}.{{{ .EntityName }}}) *{{{ toCamelCaseWithoutFirst .EntityName }}}.{{{ .EntityName }}}
    Convert{{{ title (toCamelCase .EntityName) }}}TreeToJSON(root *model.{{{ title (toCamelCase .EntityName) }}}, parent *model.{{{ title (toCamelCase .EntityName) }}}) *{{{ .EntityName }}}TreeResp
}

// {{{ .EntityName }}}Service {{{ .FunctionName }}}服务实现类
type {{{ .EntityName }}}Service struct {}

// New{{{ .EntityName }}}Service 初始化
func New{{{ .EntityName }}}Service() I{{{ .EntityName }}}Service {
	return &{{{ .EntityName }}}Service{}
}


// List {{{ .FunctionName }}}列表
func (svc *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) List(req {{{ .EntityName }}}ListReq) (resp Resp{{{ .EntityName }}}List, total int64, code int, err error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	model := model.GetDB().Model(&{{{ toCamelCaseWithoutFirst .EntityName }}}.{{{ .EntityName }}}{})
	{{{- range .Columns }}}
	{{{- if .IsQuery }}}
	{{{- $queryOpr := index $.ModelOprMap .QueryType }}}
	{{{- if and (eq .GoType "string") (eq $queryOpr "like") }}}
	if payload.{{{ title (toCamelCase .ColumnName) }}} != "" {
        model = model.Where("{{{ .ColumnName }}} like ?", "%"+payload.{{{ title (toCamelCase .ColumnName) }}}+"%")
    }
    {{{- else }}}
    if payload.{{{ title (toCamelCase .ColumnName) }}} {{{ if eq .GoType "string" }}}!= ""{{{ else }}}>=0{{{ end }}} {
        model = model.Where("{{{ .ColumnName }}} = ?", payload.{{{ title (toCamelCase .ColumnName) }}})
    }
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
	{{{- if contains .AllFields "is_delete" }}}
	model = model.Where("is_delete = ?", 0)
	{{{- end }}}
	// 总数
	var total int64
	err := model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	// 数据
	var objs []*{{{ toCamelCaseWithoutFirst .EntityName }}}.{{{ .EntityName }}}
	err = model.Limit(limit).Offset(offset).Order("id desc").Find(&objs).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	result := make([]*resp.{{{ .EntityName }}}Resp, 0)
	response.Copy(&result, objs)
	return resp.{{{ .EntityName }}}ListResp{
	    PageResp: response.PageResp{
            PageNo:   page.PageNo,
            PageSize: page.PageSize,
            Count:    count,
        },
	    Lists:    result,
	}, nil
}

// Detail {{{ .FunctionName }}}详情
func (svc *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) Detail(req {{{ .EntityName }}}DetailReq) (res *{{{ .EntityName }}}ItemResp, code int, err error) {
	// var obj *{{{ toCamelCaseWithoutFirst .EntityName }}}.{{{ .EntityName }}}
	obj, err := model.Get{{{ title (toCamelCase .EntityName) }}}Instance(
	    model.GetDB(),
	    map[string]interface{}{
	        {{{ $.PrimaryKey }}}: req.ID,
	        {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
	    }
    )
	if err != nil {
	    if errors.Is(err, gorm.ErrRecordNotFound) {
	        return nil, render.DataNotExistError, err
	    }
		return nil, render.QueryErr, err
	}
	response.Copy(&res, obj)
	{{{- range .Columns }}}
    {{{- if and .IsEdit (contains (slice "image" "avatar" "logo" "img") .GoField) }}}
    res.Avatar = util.UrlUtil.ToAbsoluteUrl(res.Avatar)
    {{{- end }}}
    {{{- end }}}
	return
}

// Create {{{ .FunctionName }}}创建
func (svc *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) Create(payload {{{ .EntityName }}}CreateReq) (res *{{{ .EntityName }}}ItemResp, code int, err error) {
	obj := model.New{{{ .EntityName }}}()
	copier.Copy(&obj, payload)
	err := model.Create{{{ title (toCamelCase .EntityName) }}}(&obj).Error
	if err != nil {
		return nil, render.CreateErr, err
	}
	return obj, render.OK, nil
}

// Update {{{ .FunctionName }}}更新
func (svc *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) Update(payload {{{ .EntityName }}}UpdateReq) (res *{{{ .EntityName }}}ItemResp, code int, err error) {
	obj, err := model.Get{{{ title (toCamelCase .EntityName) }}}Instance(
        model.GetDB(),
        map[string]interface{}{
            {{{ $.PrimaryKey }}}: req.ID,
            {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
        }
    )
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, render.DataNotExistError, err
        }
        return nil, render.QueryErr, err
    }
	// 更新
	err := model.Update{{{ title (toCamelCase .EntityName) }}}(model.GetDB(), id, *payload)
	if err != nil {
		return nil, render.UpdateErr, err
	}
	return obj, render.OK, nil
}

// Del {{{ .FunctionName }}}删除
func (svc *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) Delete(payload {{{ .EntityName }}}DeleteReq) (code int, e error) {
	obj, err := model.Get{{{ title (toCamelCase .EntityName) }}}Instance(
        model.GetDB(),
        map[string]interface{}{
            {{{ $.PrimaryKey }}}: req.ID,
            {{{ if contains .AllFields "delete_time" }}}"delete_time": 0,{{{ end }}}
        }
    )
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, render.DataNotExistError, err
        }
        return render.QueryErr, err
    }
    // 删除
    obj.DeleteTime = time.Now().Unix()
    err = model.GetDB().Save(&obj).Error
    if err != nil {
        return render.DeleteError, err
    }
	return render.OK, nil
}

// Tree {{{ .FunctionName }}}树
func (s *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) Tree(req {{{ .EntityName }}}TreeReq) (res *{{{ .EntityName }}}TreeResp, code int, err error) {
	rows, _, err := s.GetAllInstances()
    if err != nil {
        return nil, render.DBError, err
    }
    root := s.Build{{{ title (toCamelCase .EntityName) }}}Tree(rows)
    result := s.Convert{{{ title (toCamelCase .EntityName) }}}TreeToJSON(root, nil)
	return result, render.OK, nil
}

// GetAll{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}获取所有有效数据
func (s *{{{ toCamelCaseWithoutFirst .EntityName }}}Service) GetAll{{{ title (toCamelCase .EntityName) }}}() ([]*model.{{{ title (toCamelCase .EntityName) }}}, error) {
	result, err := model.GetAll{{{ title (toCamelCase .EntityName) }}}(model.GetDB())
	if err != nil {
		return nil, err
	}
	return result, err
}

// Convert{{{ title (toCamelCase .EntityName) }}}TreeToJSON 行模型数据转成JSON树结构
func (s *service) Convert{{{ title (toCamelCase .EntityName) }}}TreeToJSON(root *model.{{{ title (toCamelCase .EntityName) }}}, parent *model.{{{ title (toCamelCase .EntityName) }}}) *{{{ .EntityName }}}TreeResp {
	pName := ""
	if parent != nil {
		pName = parent.Name
	}
	result := &RespSystemOrgTree{
		ID:         root.ID,
		ParentID:   root.ParentID,
		ParentName: pName,
		Children:   make([]*{{{ .EntityName }}}TreeResp, len(root.Children)),
	}

	for i, child := range root.Children {
		result.Children[i] = s.Convert{{{ title (toCamelCase .EntityName) }}}TreeToJSON(child, root)
	}
	return result
}
