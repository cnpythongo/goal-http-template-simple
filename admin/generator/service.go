package generator

import (
	"errors"
	"github.com/jinzhu/copier"
	"goal-app/admin/generator/util"
	"goal-app/model"
	"goal-app/pkg/log"
	"goal-app/pkg/render"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type IGeneratorService interface {
	List(req *ReqGenTableList) ([]*RespGenTableItem, int64, int, error)
	GetDbTableList(req *ReqDbTableList) ([]*RespDbTable, int64, int, error)
	Create(req *ReqGenTableCreate) (int, error)
	Update(req *ReqUpdateGenTable) (int, error)
	Delete(req *ReqDelTable) (int, error)
	Preview(req *ReqPreview) ([]*RespPreviewItem, int, error)
	GenCode(tableName string) (int, error)

	GetGenColumnList(tableId int64) ([]*RespGenColumn, int, error)
	UpdateGenColumn(req *ReqUpdateGenColumn) (int, error)
	DeleteGenTableColumns(req *ReqDelGenTableColumn) (int, error)
}

type generatorService struct {
}

func NewGeneratorService() IGeneratorService {
	return &generatorService{}
}

func (s *generatorService) List(req *ReqGenTableList) ([]*RespGenTableItem, int64, int, error) {
	var query []string
	var args []interface{}
	if req.Name != "" {
		query = append(query, "name like concat('%s', ?, '%')")
		args = append(args, req.Name)
	}
	if req.TableComment != "" {
		query = append(query, "table_comment concat('%s', ?, '%')")
		args = append(args, req.TableComment)
	}
	if req.CreateTimeStart > 0 {
		query = append(query, "create_time >= ?")
		args = append(args, req.CreateTimeStart)
	}
	if req.CreateTimeEnd > 0 {
		query = append(query, "create_time <= ?")
		args = append(args, req.CreateTimeEnd)
	}
	rows, total, err := model.GetGenTableList(model.GetDB(), req.Page, req.Limit, query, args)
	if err != nil {
		return nil, total, render.QueryError, err
	}

	result := make([]*RespGenTableItem, 0)
	for _, row := range rows {
		item := new(RespGenTableItem)
		err = copier.Copy(item, row)
		if err != nil {
			log.GetLogger().Error(err)
			return nil, total, render.DBAttributesCopyError, err
		}
		result = append(result, item)
	}
	return result, total, render.OK, nil
}

func (s *generatorService) GetDbTableList(req *ReqDbTableList) ([]*RespDbTable, int64, int, error) {
	result := make([]*RespDbTable, 0)
	query := util.GenUtil.GetDbTables(model.GetDB(), req.TableName, req.TableComment)
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		log.GetLogger().Errorf("GenUtil.GetDbTables Count Error ==> %v", err)
		return result, total, render.QueryError, err
	}
	if total == 0 {
		return result, total, render.OK, nil
	}
	err = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Find(&result).Error
	return result, total, render.OK, nil
}

func (s *generatorService) Create(req *ReqGenTableCreate) (int, error) {
	var dbTables []*RespDbTable
	err := util.GenUtil.GetDbTablesByName(model.GetDB(), strings.Split(req.Tables, ",")).Find(&dbTables).Error
	if err != nil {
		log.GetLogger().Errorf("GenUtil.GetDbTablesByName Error ==> %v", err)
		return render.QueryError, err
	}
	if len(dbTables) == 0 {
		log.GetLogger().Errorf("dbTables is empty")
		return render.QueryError, nil
	}

	tables := make([]*model.GenTable, 0)
	err = copier.Copy(&tables, &dbTables)
	if err != nil {
		log.GetLogger().Errorf("service.ImportTable copier.Copy Error ==> %v", err)
		return render.DBAttributesCopyError, err
	}

	tx := model.GetDB().Begin()
	for _, table := range tables {
		//生成表信息
		genTable := util.GenUtil.CreateGenTable(table)
		txErr := tx.Create(&genTable).Error
		if txErr != nil {
			tx.Rollback()
			log.GetLogger().Errorf("GenUtil.CreateGenTable Error ==> %v", txErr)
			return render.CreateError, txErr
		}

		// 生成列信息
		var columns []*model.GenTableColumn
		txErr = util.GenUtil.GetDbTableColumnsByTableName(model.GetDB(), table.Name).Find(&columns).Error
		if txErr != nil {
			tx.Rollback()
			log.GetLogger().Errorf("GenUtil.GetDbTableColumnsByTableName Error ==> %v", txErr)
			return render.QueryError, txErr
		}
		for j := 0; j < len(columns); j++ {
			column := util.GenUtil.CreateGenColumn(genTable.ID, columns[j])
			txErr = tx.Create(&column).Error
			if txErr != nil {
				tx.Rollback()
				log.GetLogger().Errorf("GenUtil.CreateGenColumn Error ==> %v", txErr)
				return render.QueryError, txErr
			}
		}
	}
	tx.Commit()

	return render.OK, nil
}

// getSubTableInfo 根据主表获取子表主键和列信息
func (s *generatorService) getSubTableInfo(genTable *model.GenTable) (pkCol *model.GenTableColumn, cols []*model.GenTableColumn, e error) {
	if genTable.SubTableName == "" || genTable.SubTableFk == "" {
		return
	}
	table := model.NewGenTable()
	err := model.GetDB().Where("name = ?", genTable.SubTableName).Limit(1).First(&table).Error
	if err != nil {
		return nil, nil, err
	}
	err = util.GenUtil.GetDbTableColumnsByTableName(model.GetDB(), genTable.SubTableName).Find(&cols).Error
	if err != nil {
		return nil, nil, err
	}
	pkCol = util.GenUtil.CreateGenColumn(table.ID, util.GenUtil.GetTablePriCol(cols))
	return
}

// renderCodeByTable 根据主表和模板文件渲染模板代码
func (s *generatorService) renderCodeByTable(genTable *model.GenTable) (map[string]string, error) {
	columns := make([]*model.GenTableColumn, 0)
	err := model.GetDB().Where("gen_table_id = ?", genTable.ID).Order("sort asc").Find(&columns).Error
	if err != nil {
		return nil, err
	}

	//获取子表信息
	pkCol, cols, err := s.getSubTableInfo(genTable)
	if err != nil {
		return nil, err
	}

	//获取模板变量信息
	vars := util.TemplateUtil.PrepareVars(genTable, columns, pkCol, cols)
	//生成模板内容
	result := make(map[string]string)
	for _, tplPath := range util.TemplateUtil.GetTemplatePaths(genTable.GenTpl) {
		result[tplPath], err = util.TemplateUtil.Render(tplPath, vars)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *generatorService) Preview(req *ReqPreview) ([]*RespPreviewItem, int, error) {
	genTable := model.NewGenTable()
	err := model.GetDB().Where("id = ?", req.ID).Limit(1).First(&genTable).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, nil
		}
		log.GetLogger().Errorf("generator.service.PreviewCode Query Error ==> %v", err)
		return nil, render.QueryError, err
	}

	//获取模板内容
	tplCodeMap, err := s.renderCodeByTable(genTable)
	if err != nil {
		log.GetLogger().Errorf("generator.service.PreviewCode renderCodeByTable Error ==> %v", err)
		return nil, render.Error, err
	}

	result := make([]*RespPreviewItem, 0)
	for tplPath, tplCode := range tplCodeMap {
		result = append(result, &RespPreviewItem{
			Name:     strings.ReplaceAll(tplPath, ".tpl", ""),
			Language: util.TemplateUtil.GetTplLang(tplPath),
			Content:  tplCode,
		})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, render.OK, nil
}

func (s *generatorService) GenCode(tableName string) (int, error) {
	genTable, err := model.GetGenTableInstance(model.GetDB(), map[string]interface{}{"name": tableName})
	if err != nil {
		log.GetLogger().Error(err)
		return render.QueryError, err
	}
	//获取模板内容
	tplCodeMap, err := s.renderCodeByTable(genTable)
	if err != nil {
		log.GetLogger().Errorf("generator.service.GenCode renderCodeByTable Error ==> %v", err)
		return render.Error, err
	}
	//获取生成根路径
	basePath := util.TemplateUtil.GetGenPath(genTable)
	// 生成代码文件
	err = util.TemplateUtil.GenCodeFiles(tplCodeMap, genTable.ModuleName, basePath)
	if err != nil {
		log.GetLogger().Errorf("generator.service.GenCode GenCodeFiles Error ==> %v", err)
		return render.Error, err
	}
	return render.OK, nil
}

func (s *generatorService) Update(req *ReqUpdateGenTable) (int, error) {
	table, err := model.GetGenTableInstance(model.GetDB(), map[string]interface{}{"id": req.ID})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return render.DataNotExistError, nil
		}
		log.GetLogger().Errorf("generator.service.Update Query Error ==> %v", err)
		return render.QueryError, err
	}

	err = copier.Copy(&table, &req)
	if err != nil {
		log.GetLogger().Errorf("generator.service.Update copier.Copy Error ==> %v", err)
		return render.DBAttributesCopyError, err
	}
	err = model.UpdateGenTable(model.GetDB(), table)
	if err != nil {
		return render.UpdateError, err
	}
	return render.OK, nil
}

func (s *generatorService) Delete(req *ReqDelTable) (int, error) {
	tx := model.GetDB().Begin()
	err := tx.Where("gen_table_id in ?", req.Ids).Delete(&model.GenTableColumn{}).Error
	if err != nil {
		tx.Rollback()
		return render.DeleteError, err
	}
	err = tx.Where("id in ?", req.Ids).Delete(&model.GenTable{}).Error
	if err != nil {
		tx.Rollback()
		return render.DeleteError, err
	}
	tx.Commit()
	return render.OK, nil
}

func (s *generatorService) GetGenColumnList(tableId int64) ([]*RespGenColumn, int, error) {
	genTable, err := model.GetGenTableInstance(model.GetDB(), map[string]interface{}{"id": tableId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, render.DataNotExistError, nil
		}
		log.GetLogger().Errorf("generator.service.GetGenColumnList Query Error ==> %v", err)
		return nil, render.QueryError, err
	}

	rows, err := model.GetGenTableColumnList(model.GetDB(), genTable.ID)
	if err != nil {
		log.GetLogger().Errorf("generator.service.GetGenColumnList Query Error ==> %v", err)
		return nil, render.QueryError, err
	}
	result := make([]*RespGenColumn, 0)
	err = copier.Copy(&result, &rows)
	if err != nil {
		log.GetLogger().Errorf("service.GetGenColumnList copier.Copy Error ==> %v", err)
		return nil, render.DBAttributesCopyError, err
	}
	return result, render.OK, nil
}

func (s *generatorService) UpdateGenColumn(req *ReqUpdateGenColumn) (int, error) {
	column, err := model.GetGenTableColumnInstance(model.GetDB(), req.ID)
	if err != nil {
		log.GetLogger().Errorf("generator.service.UpdateGenColumn Query Error ==> %v", err)
		return render.DataNotExistError, err
	}

	err = copier.Copy(&column, &req)
	if err != nil {
		log.GetLogger().Errorf("service.UpdateGenColumn copier.Copy Error ==> %v", err)
		return render.DBAttributesCopyError, err
	}
	err = model.UpdateGenTableColumn(model.GetDB(), column)
	if err != nil {
		return render.UpdateError, err
	}
	return render.OK, nil
}

func (s *generatorService) DeleteGenTableColumns(req *ReqDelGenTableColumn) (int, error) {
	tx := model.GetDB().Begin()
	err := tx.Where("id in ?", req.Ids).Delete(&model.GenTableColumn{}).Error
	if err != nil {
		tx.Rollback()
		return render.DeleteError, err
	}
	tx.Commit()
	return render.OK, nil
}
