import { get, post } from '@/request';
import { PageParams, PageList } from '@/api/types';

export interface {{{.EntityName}}}ListParams extends PageParams {
  {{{- range .Columns }}}
  {{{- if .IsQuery }}}
  {{{ toSnakeCase .ColumnName }}}: {{{- if eq .HtmlType "number" }}}number{{{- else }}}string{{{- end }}};
  {{{- end }}}
  {{{- end }}}
}

export interface {{{.EntityName}}}CreateBody {
{{{- range .Columns }}}
{{{- if .IsInsert }}}
{{{ toSnakeCase .ColumnName }}}: {{{- if eq .HtmlType "number" }}}number{{{- else }}}string{{{- end }}};
{{{- end }}}
{{{- end }}}
}

export interface {{{.EntityName}}}UpdateBody {
  {{{- range .Columns }}}
  {{{- if .IsEdit }}}
  {{{ toSnakeCase .ColumnName }}}: {{{- if eq .HtmlType "number" }}}number{{{- else }}}string{{{- end }}};
  {{{- end }}}
  {{{- end }}}
}

export interface {{{.EntityName}}}DeleteBody {
  ids: Array<number>;
}

export interface {{{.EntityName}}}Item {
  {{{- range .Columns }}}
  {{{- if .IsList }}}
  {{{ toSnakeCase .ColumnName }}}: {{{- if eq .HtmlType "number" }}}number{{{- else }}}string{{{- end }}};
  {{{- end }}}
  {{{- end }}}
  children?: Array<{{{.EntityName}}}Item>;
}

export default {
  // {{{.EntityName}}}列表
  list: (params: {{{.EntityName}}}ListParams) =>
    get<PageList<{{{.EntityName}}}Item>>('{{{.GenPath}}}/list', params),
  // {{{.EntityName}}}详情
  detail: (id: number) => get<{{{.EntityName}}}Item>('{{{.GenPath}}}/detail', { id }),
  // {{{.EntityName}}}新增
  create: (data: {{{.EntityName}}}CreateBody) => post<any>('{{{.GenPath}}}/create', data),
  // {{{.EntityName}}}更新
  update: (data: {{{.EntityName}}}UpdateBody) => post<any>('{{{.GenPath}}}/update', data),
  // {{{.EntityName}}}删除
  delete: (data: {{{.EntityName}}}DeleteBody) => post<any>('{{{.GenPath}}}/delete', data),
  {{{- if eq .GenTpl "tree" -}}}}
  // 获取树结构数据
  getTreeData: () => get<Array<{{{.EntityName}}}Item>>('{{{.GenPath}}}/tree'),
  {{{- end -}}}}
};
