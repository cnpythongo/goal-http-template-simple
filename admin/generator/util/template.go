package util

import (
	"archive/zip"
	"bytes"
	"fmt"
	"goal-app/model"
	"goal-app/pkg/config"
	"goal-app/pkg/utils"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
)

// genUtil 模板工具
type templateUtil struct {
	basePath string
	tpl      *template.Template
}

// TplVars 模板变量
type TplVars struct {
	GenTpl          string
	Name            string
	AuthorName      string
	PackageName     string
	EntityName      string
	EntitySnakeName string
	ModuleName      string
	FunctionName    string
	JavaCamelField  string
	DateFields      []string
	PrimaryKey      string
	PrimaryField    string
	AllFields       []string
	SubPriCol       *model.GenTableColumn
	SubPriField     string
	SubTableFields  []string
	ListFields      []string
	DetailFields    []string
	DictFields      []string
	IsSearch        bool
	ModelOprMap     map[string]string
	Table           model.GenTable
	Columns         []*model.GenTableColumn
	SubColumns      []*model.GenTableColumn
	GenPath         string
	//ModelTypeMap    map[string]string
}

var TemplateUtil = &templateUtil{
	basePath: "admin/generator/templates",
	tpl: template.New("").Delims("{{{", "}}}").Funcs(
		template.FuncMap{
			"sub":                     sub,
			"slice":                   slice,
			"title":                   utils.UpperFirst,
			"toSnakeCase":             utils.ToSnakeCase,
			"toCamelCase":             utils.ToCamelCase,
			"toCamelCaseWithoutFirst": utils.ToCamelCaseWithoutFirst,
			"contains":                utils.ToolsUtil.Contains,
			"lowerFirst":              utils.LowerFirst,
		}),
}

// sub 模板-减函数
func sub(a, b int) int {
	return a - b
}

// slice 模板-创建切片
func slice(items ...interface{}) []interface{} {
	return items
}

// zFile 待加入zip的文件
type zFile struct {
	Name string
	Body string
}

// PrepareVars 获取模板变量信息
func (t *templateUtil) PrepareVars(table *model.GenTable, columns []*model.GenTableColumn,
	oriSubPriCol *model.GenTableColumn, oriSubCols []*model.GenTableColumn) TplVars {
	subPriField := "id"
	isSearch := false
	primaryKey := "id"
	primaryField := "id"
	functionName := "【请填写功能名称】"
	var allFields []string
	var subTableFields []string
	var listFields []string
	var detailFields []string
	var dictFields []string
	var subColumns []*model.GenTableColumn
	var oriSubColNames []string
	for _, column := range oriSubCols {
		oriSubColNames = append(oriSubColNames, column.ColumnName)
	}
	if oriSubPriCol != nil && oriSubPriCol.ID > 0 {
		subPriField = oriSubPriCol.ColumnName
		subColumns = append(subColumns, oriSubPriCol)
	}
	for _, column := range columns {
		allFields = append(allFields, column.ColumnName)
		if utils.ToolsUtil.Contains(oriSubColNames, column.ColumnName) {
			subTableFields = append(subTableFields, column.ColumnName)
			subColumns = append(subColumns, column)
		}
		if column.IsList == 1 {
			listFields = append(listFields, column.ColumnName)
		}
		if column.IsEdit == 1 {
			detailFields = append(detailFields, column.ColumnName)
		}
		if column.IsQuery == 1 {
			isSearch = true
		}
		if column.IsPk == 1 {
			primaryKey = column.GoField
			primaryField = column.ColumnName
		}
		if column.DictType != "" && !utils.ToolsUtil.Contains(dictFields, column.DictType) {
			dictFields = append(dictFields, column.DictType)
		}
	}
	//QueryType转换查询比较运算符
	modelOprMap := map[string]string{
		"=":    "==",
		"LIKE": "like",
	}
	if table.FunctionName != "" {
		functionName = table.FunctionName
	}
	return TplVars{
		GenTpl:          table.GenTpl,
		Name:            table.Name,
		AuthorName:      table.AuthorName,
		PackageName:     table.ModuleName,
		EntityName:      table.EntityName,
		EntitySnakeName: utils.ToSnakeCase(table.EntityName),
		ModuleName:      table.ModuleName,
		FunctionName:    functionName,
		DateFields:      SqlConstants.ColumnTimeName,
		PrimaryKey:      primaryKey,
		PrimaryField:    primaryField,
		AllFields:       allFields,
		SubPriCol:       oriSubPriCol,
		SubPriField:     subPriField,
		SubTableFields:  subTableFields,
		ListFields:      listFields,
		DetailFields:    detailFields,
		DictFields:      dictFields,
		IsSearch:        isSearch,
		ModelOprMap:     modelOprMap,
		Columns:         columns,
		SubColumns:      subColumns,
		GenPath:         table.GenPath,
	}
}

// GetTemplatePaths 获取模板路径
func (t *templateUtil) GetTemplatePaths(genTpl string) []string {
	tplPaths := []string{
		"go/route.go.tpl",
		"go/model.go.tpl",
		"go/schema.go.tpl",
		"go/handler.go.tpl",
		"go/service.go.tpl",
		"react/api.ts.tpl",
		"react/index.tsx.tpl",
		//"react/edit.tsx.tpl",
	}
	if genTpl == GenConstants.TplTree {
		tplPaths = append(tplPaths, "react/index-tree.tsx.tpl")
	}
	return tplPaths
}

// Render 渲染模板
func (t *templateUtil) Render(tplPath string, tplVars TplVars) (res string, e error) {
	tpl, err := t.tpl.ParseFiles(path.Join(config.GetConfig().App.RootPath, t.basePath, tplPath))
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = tpl.ExecuteTemplate(buf, path.Base(tplPath), tplVars)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetGenPath 获取生成路径
func (t *templateUtil) GetGenPath(table *model.GenTable) string {
	if table.GenPath == "/" {
		return path.Join(config.GetConfig().App.RootPath, GenConfig.GenRootPath)
	}
	return table.GenPath
}

// GetFilePaths 获取生成文件相对路径
func (t *templateUtil) GetFilePaths(tplCodeMap map[string]string, moduleName string) map[string]string {
	//模板文件对应的输出文件
	fmtMap := map[string]string{
		"go/route.go.tpl":     "go/%s/route.go",
		"go/model.go.tpl":     "go/%s/model.go",
		"go/schema.go.tpl":    "go/%s/schema.go",
		"go/handler.go.tpl":   "go/%s/handler.go",
		"go/service.go.tpl":   "go/%s/service.go",
		"react/api.ts.tpl":    "react/%s/api.ts",
		"react/index.tsx.tpl": "react/%s/index.tsx",
		//"react/edit.tsx.tpl":       "react/%s/edit.tsx",
		"react/index-tree.tsx.tpl": "react/%s/index-tree.tsx",
	}
	filePath := make(map[string]string)
	for tplPath, tplCode := range tplCodeMap {
		file := fmt.Sprintf(fmtMap[tplPath], moduleName)
		filePath[file] = tplCode
	}
	return filePath
}

// GenCodeFiles 生成代码文件
func (t *templateUtil) GenCodeFiles(tplCodeMap map[string]string, moduleName string, basePath string) error {
	filePaths := t.GetFilePaths(tplCodeMap, moduleName)
	for file, tplCode := range filePaths {
		filePath := path.Join(basePath, file)
		dir := path.Dir(filePath)
		if !utils.ToolsUtil.IsFileExist(dir) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
		err := os.WriteFile(filePath, []byte(tplCode), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, file zFile) error {
	header := &zip.FileHeader{
		Name:   file.Name,
		Method: zip.Deflate,
	}
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.WriteString(writer, file.Body)
	if err != nil {
		return err
	}
	return nil
}

// GenZip 生成代码压缩包
func (t *templateUtil) GenZip(zipWriter *zip.Writer, tplCodeMap map[string]string, moduleName string) error {
	filePaths := t.GetFilePaths(tplCodeMap, moduleName)
	files := make([]zFile, 0)
	for file, tplCode := range filePaths {
		files = append(files, zFile{
			Name: file,
			Body: tplCode,
		})
	}
	for _, file := range files {
		err := addFileToZip(zipWriter, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *templateUtil) GetTplLang(tpl string) string {
	if strings.Contains(tpl, ".go.tpl") {
		return "go"
	} else if strings.Contains(tpl, ".ts.tpl") {
		return "ts"
	} else {
		return "tsx"
	}
}
