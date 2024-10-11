package {{{ .PackageName }}}


// Req{{{ .EntityName }}}List {{{ .FunctionName }}}列表请求参数
type Req{{{ .EntityName }}}List struct {
    render.Pagination
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Req{{{ .EntityName }}}Tree {{{ .FunctionName }}}树结构请求参数
type Req{{{ .EntityName }}}Tree struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Req{{{ .EntityName }}}Detail {{{ .FunctionName }}}详情请求参数
type Req{{{ .EntityName }}}Detail struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Req{{{ .EntityName }}}Create {{{ .FunctionName }}}创建请求参数
type Req{{{ .EntityName }}}Create struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Req{{{ .EntityName }}}Update {{{ .FunctionName }}}更新请求参数
type Req{{{ .EntityName }}}Update struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" form:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Req{{{ .EntityName }}}Delete {{{ .FunctionName }}}删除请求参数
type Req{{{ .EntityName }}}Delete struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ title (toCamelCase .GoField) }}}s []{{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}s" form:"{{{ toSnakeCase .GoField }}}s"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

// Resp{{{ .EntityName }}}Item {{{ .FunctionName }}}单条详情
type Resp{{{ .EntityName }}}Item struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" structs:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

{{{- if eq .GenTpl "tree" }}}
// Resp{{{ .EntityName }}}Tree {{{ .FunctionName }}}树结构数据
type Resp{{{ .EntityName }}}Tree struct {
    {{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `json:"{{{ toSnakeCase .GoField }}}" structs:"{{{ toSnakeCase .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    ParentName string                `json:"parent_name" structs:"parent_name"` // '父级名称'
    Children []*Resp{{{ .EntityName }}}Tree `json:"children"`    // 子节点
}
{{{- end }}}
