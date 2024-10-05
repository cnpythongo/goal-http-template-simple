package {{{ .PackageName }}}


// {{{ .EntityName }}}ListReq {{{ .FunctionName }}}列表请求参数
type {{{ .EntityName }}}ListReq struct {
    render.Pagination
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

type {{{ .EntityName }}}TreeReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}DetailReq {{{ .FunctionName }}}详情请求参数
type {{{ .EntityName }}}DetailReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}CreateReq {{{ .FunctionName }}}创建请求参数
type {{{ .EntityName }}}CreateReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}UpdateReq {{{ .FunctionName }}}更新请求参数
type {{{ .EntityName }}}UpdateReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}DeleteReq {{{ .FunctionName }}}删除请求参数
type {{{ .EntityName }}}DeleteReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}ItemResp {{{ .FunctionName }}}单条详情
type {{{ .EntityName }}}ItemResp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" structs:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// {{{ .EntityName }}}TreeResp {{{ .FunctionName }}}树结构数据
type {{{ .EntityName }}}TreeResp struct {
    {{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" structs:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    Children []*{{{ .EntityName }}}TreeResp `json:"children"`    // 子节点
}
