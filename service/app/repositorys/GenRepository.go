package repositorys

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadminGen"
	"github.com/sonhineboy/gsadminGen/pkg"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

//
// GenRepository
//  @Description: 代码生成核心业务罗杰
//
type GenRepository struct {
	BaseRepository
}

func NewGenRepository() *GenRepository {
	return &GenRepository{}
}

//
// GetTables
//  @Description: 获取数据表
//  @receiver     r *GenRepository
//  @return       tableSlice
//  @return       err
//
func (r *GenRepository) GetTables() (tableSlice []map[string]string, err error) {
	tableSlice = make([]map[string]string, 0, 10)
	tables, err := r.getDb().Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	for _, v := range tables {
		tableSlice = append(tableSlice, map[string]string{"label": v, "value": v})
	}
	return
}

//
// TableField
//  @Description: 表字段
//  @receiver     r *GenRepository
//  @param        name string
//  @param        function func(fieldsSlices []map[string]interface{}, columnType gorm.ColumnType, r *GenRepository) []map[string]interface{}
//  @return       []map[string]interface{}
//  @return       error
//
func (r *GenRepository) TableField(name string, function func(fieldsSlices []map[string]interface{}, columnType gorm.ColumnType, r *GenRepository) []map[string]interface{}) ([]map[string]interface{}, error) {

	column, err := r.getDb().Migrator().ColumnTypes(name)

	if err != nil {
		return nil, err
	}

	ignoreField := r.getIgnoreField()

	fieldsSlices := make([]map[string]interface{}, 0, 20)

	for _, v := range column {
		_, ok := ignoreField[v.Name()]
		if ok {
			continue
		}
		fieldsSlices = function(fieldsSlices, v, r)
	}

	return fieldsSlices, nil
}

//
// GetIndexType
//  @Description: 获取索引类型
//  @receiver     r *GenRepository
//  @param        column string
//  @param        Indexes map[string]map[string]interface{}
//  @return       string
//
func (r *GenRepository) GetIndexType(column string, Indexes map[string]map[string]interface{}) string {

	v, ok := Indexes[column]
	if ok {
		if v["Non_unique"] == 0 {
			return "UNIQUE"
		}
		if v["Index_type"] == "FULLTEXT" {
			return "FULLTEXT"
		}
		return "NORMAL"
	}
	return "Null"
}

//
// GetTablesIndexes
//  @Description: 获取索引
//  @receiver     r *GenRepository
//  @param        tables string
//  @return       indexMap
//
func (r *GenRepository) GetTablesIndexes(tables string) (indexMap map[string]map[string]interface{}) {
	var indexes []map[string]interface{}
	r.getDb().Raw(fmt.Sprint("show Index from ", tables)).Scan(&indexes)
	indexMap = make(map[string]map[string]interface{}, 20)
	for _, v := range indexes {
		columnName, _ := v["Column_name"]
		if columnNameString, ok := columnName.(string); ok {
			indexMap[columnNameString] = map[string]interface{}{
				"Non_unique": v["Non_unique"],
				"Index_type": v["Index_type"],
			}
		} else {
			continue
		}
	}
	return
}

//
// getIgnoreField
//  @Description: 获取忽略字段
//  @receiver     r *GenRepository
//  @return       map[string]string
//
func (r *GenRepository) getIgnoreField() map[string]string {
	return map[string]string{
		"id":         "true",
		"created_at": "true",
		"updated_at": "true",
		"deleted_at": "true",
	}
}

//
// GenCode
//  @Description: 根据数据生成业务代码
//  @receiver     r *GenRepository
//  @param        data requests.GenCode
//  @return       error
//
func (r *GenRepository) GenCode(data requests.GenCode) error {

	// 如果有表前缀去掉
	prefixReg := regexp.MustCompile(fmt.Sprint("^", global.Config.Db.TablePrefix))
	data.TableDiyName = prefixReg.ReplaceAllString(data.TableDiyName, "")

	v := pkg.TableModal{
		Name:   data.TableDiyName,
		Fields: data.Fields,
	}

	var err error

	if global.SlicesHasStr(data.Checkbox, "生成Controller") {
		if err = gsadminGen.GenController("./app/controllers/"+data.ControllerPackage+"/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Controller.go", v, data.ControllerPackage); err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成Model") {
		if err = gsadminGen.GenModel("./app/models/"+gsadminGen.UnderToConvertSoreLow(v.Name)+".go", v); err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成Request") {
		if err = gsadminGen.GenRequest("./app/requests/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Request.go", v); err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成Repository") {
		err = gsadminGen.GenRepository("./app/repositorys/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"Repository.go", v)
		if err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成前端模板") {
		err = r.genWebTemp(v)
		if err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成路由") {

		routerWriter := pkg.NewWriterRouter(fmt.Sprint(global.GAD_APP_PATH, "router/systemApi.go"), "//router gen start not delete", data.ControllerPackage)
		err = routerWriter.Write(r.getRouters(v, data))
		if err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}

	}

	if global.SlicesHasStr(data.Checkbox, "生成数据库") {

		dbTableWriter := pkg.NewWriterAutoModel(fmt.Sprint(global.GAD_APP_PATH, "initialize/dbInit.go"), "//slot start not delete")
		err = dbTableWriter.Write([]string{
			fmt.Sprint("\t\t&models.", strings.Title(gsadminGen.UnderToConvertSoreLow(v.Name)), "{},"),
		})

		if err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	if global.SlicesHasStr(data.Checkbox, "生成菜单") {
		err := r.genMenu(v, data)
		if err != nil {
			return fmt.Errorf("genCode err： %v ", err)
		}
	}

	return nil
}

//
// genWebTemp
//  @Description: 生成前台模板
//  @receiver     r *GenRepository
//  @param        v pkg.TableModal
//  @return       err
//
func (r *GenRepository) genWebTemp(v pkg.TableModal) error {
	err := gsadminGen.GenIndex(`../web/scui/src/views/`+gsadminGen.UnderToConvertSoreLow(v.Name)+"/"+"index.vue", v)
	if err != nil {
		return fmt.Errorf("genWebTemp Index Err %v", err)
	}

	err = gsadminGen.GenForm("../web/scui/src/views/"+gsadminGen.UnderToConvertSoreLow(v.Name)+"/"+"form.vue", v)
	if err != nil {
		return fmt.Errorf("genWebTemp Form Err %v", err)
	}
	err = gsadminGen.GenApi("../web/scui/src/api/model/"+gsadminGen.UnderToConvertSoreLow(v.Name)+".js", v) //
	if err != nil {
		return fmt.Errorf("genWebTemp Api Err %v", err)
	}
	return nil
}

//
// getRouters
//  @Description: 获取要生成路由的数据
//  @receiver     r
//  @param        v pkg.TableModal
//  @param        data requests.GenCode
//  @return       []string
//
func (r *GenRepository) getRouters(v pkg.TableModal, data requests.GenCode) []string {

	return []string{
		"",
		fmt.Sprint("\t", "//gen_", gsadminGen.UnderToConvertSoreLow(v.Name)),
		fmt.Sprint("\t", gsadminGen.UnderToConvertSoreLow(v.Name), " :=", " r.Group(\"", gsadminGen.UnderToConvertSoreLow(v.Name), "\")"),
		fmt.Sprint("\t", "{"),
		fmt.Sprint("\t\t", "var ", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", " ", data.ControllerPackage, ".", strings.Title(gsadminGen.UnderToConvertSoreLow(v.Name)), "Controller"),
		fmt.Sprint("\t\t", gsadminGen.UnderToConvertSoreLow(v.Name), ".GET(\"/index\",", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", ".Index)"),
		fmt.Sprint("\t\t", gsadminGen.UnderToConvertSoreLow(v.Name), ".POST(\"/save\",", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", ".Save)"),
		fmt.Sprint("\t\t", gsadminGen.UnderToConvertSoreLow(v.Name), ".POST(\"/delete\",", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", ".Delete)"),
		fmt.Sprint("\t\t", gsadminGen.UnderToConvertSoreLow(v.Name), ".GET(\"/:id\",", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", ".Get)"),
		fmt.Sprint("\t\t", gsadminGen.UnderToConvertSoreLow(v.Name), ".POST(\"/edit/:id\",", gsadminGen.UnderToConvertSoreLow(v.Name), "Controller", ".Edit)"),
		fmt.Sprint("\t", "}"),
		"",
	}
}

//
// getMenus
//  @Description: 获取要生成路由的数据
//  @receiver     r
//  @param        v pkg.TableModal
//  @param        data requests.GenCode
//  @return       []requests.MenuPost
//
func (r *GenRepository) getMenus(v pkg.TableModal, data requests.GenCode) []requests.MenuPost {

	return []requests.MenuPost{
		{
			Component: gsadminGen.UnderToConvertSoreLow(v.Name),
			Name:      gsadminGen.UnderToConvertSoreLow(v.Name),
			ParentId:  0,
			Path:      fmt.Sprint("/", gsadminGen.UnderToConvertSoreLow(v.Name)),
			Meta: map[string]interface{}{
				"icon":  "el-icon-menu",
				"type":  "menu",
				"title": data.MenuName,
			},
			ApiList: []map[string]string{
				{"url": fmt.Sprint("/api/", gsadminGen.UnderToConvertSoreLow(v.Name), "/index"), "code": "get"},
			},
			Sort: 0,
		},
		{
			Name: fmt.Sprint(gsadminGen.UnderToConvertSoreLow(v.Name), ".save"),
			Meta: map[string]interface{}{
				"icon":  "el-icon-menu",
				"type":  "button",
				"title": "新增",
			},
			ApiList: []map[string]string{
				{"url": fmt.Sprint("/api/", gsadminGen.UnderToConvertSoreLow(v.Name), "/save"), "code": "post"},
			},
			Sort: 1,
		},
		{
			Name: fmt.Sprint(gsadminGen.UnderToConvertSoreLow(v.Name), ".del"),
			Meta: map[string]interface{}{
				"icon":  "el-icon-menu",
				"type":  "button",
				"title": "删除",
			},
			ApiList: []map[string]string{
				{"url": fmt.Sprint("/api/", gsadminGen.UnderToConvertSoreLow(v.Name), "/delete"), "code": "post"},
			},
			Sort: 2,
		},
		{
			Name: fmt.Sprint(gsadminGen.UnderToConvertSoreLow(v.Name), ".get"),
			Meta: map[string]interface{}{
				"icon":  "el-icon-menu",
				"type":  "button",
				"title": "查看",
			},
			ApiList: []map[string]string{
				{"url": fmt.Sprint("/api/", gsadminGen.UnderToConvertSoreLow(v.Name), "/:id"), "code": "get"},
			},
			Sort: 3,
		},
		{
			Name: fmt.Sprint(gsadminGen.UnderToConvertSoreLow(v.Name), ".edit"),
			Meta: map[string]interface{}{
				"icon":  "el-icon-menu",
				"type":  "button",
				"title": "编辑",
			},
			ApiList: []map[string]string{
				{"url": fmt.Sprint("/api/", gsadminGen.UnderToConvertSoreLow(v.Name), "/edit/:id"), "code": "post"},
			},
			Sort: 4,
		},
	}

}

//
// genMenu
//  @Description: 生成数据库菜单
//  @receiver     r *GenRepository
//  @param        v pkg.TableModal
//  @param        data requests.GenCode
//  @return       error
//
func (r *GenRepository) genMenu(v pkg.TableModal, data requests.GenCode) error {

	var err error
	menuRe := &SystemMenuRepository{}

	menus := r.getMenus(v, data)

	err = r.getDb().Transaction(func(tx *gorm.DB) error {

		menuRe.SetDb(tx)

		var (
			model  models.AdminMenu
			result *gorm.DB
		)
		for i, menu := range menus {
			//每次需要清空值，否则会有问题
			menuRe.MenuModel = models.AdminMenu{}
			if i == 0 {
				result, model = menuRe.Add(menu)
			} else {
				menu.ParentId = model.ID
				result, _ = menuRe.Add(menu)
			}

			if result.Error != nil {
				return result.Error
			}

		}
		return nil
	})
	return err
}
