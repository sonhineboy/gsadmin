package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"gorm.io/gorm"
)

type GenController struct{}

func (gen *GenController) GetTables(ctx *gin.Context) {
	var (
		re = repositorys.NewGenRepository()
	)
	tables, err := re.GetTables()
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", tables)

}

func (gen *GenController) TableFields(ctx *gin.Context) {

	var (
		data requests.GenFields
		re   = repositorys.NewGenRepository()
	)

	err := ctx.ShouldBind(&data)

	if err != nil {
		response.Failed(ctx, global.GetError(err, data))
		return
	}

	v, err := re.TableField(data.TableName, func(fieldsSlices []map[string]interface{}, columnType gorm.ColumnType, r *repositorys.GenRepository) []map[string]interface{} {

		defaultValue, _ := columnType.DefaultValue()
		commentValue, _ := columnType.Comment()
		primary, _ := columnType.PrimaryKey()
		isNull, _ := columnType.Nullable()

		indexes := r.GetTablesIndexes(data.TableName)

		fieldsSlices = append(fieldsSlices, map[string]interface{}{
			"name":      columnType.Name(),
			"type":      columnType.DatabaseTypeName(),
			"isNull":    isNull,
			"default":   defaultValue,
			"describe":  commentValue,
			"json":      columnType.Name(),
			"primary":   primary,
			"transform": commentValue,
			"index":     r.GetIndexType(columnType.Name(), indexes),
		})

		return fieldsSlices
	})
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", v)
}
func (gen *GenController) GenCode(ctx *gin.Context) {
	var (
		data requests.GenCode
		re   = repositorys.NewGenRepository()
	)

	err := ctx.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		global.Logger.Errorf("参数绑定错误：%v", err)
		response.Failed(ctx, global.GetError(err, data)+"xxxx")
		return
	}

	err = re.GenCode(data)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	response.Success(ctx, "ok", nil)
}
