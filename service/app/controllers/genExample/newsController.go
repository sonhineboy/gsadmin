package genExample

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"strconv"
)

type NewsController struct{}

func (controller *NewsController) Index(ctx *gin.Context) {

	var (
		params global.List
		re = repositorys.NewNewsRepository()
	)
	_ = ctx.ShouldBind(&params)
	response.Success(ctx, "ok", re.Page(params.Where, params.Page, params.PageSize, "created_at"))
}

func (controller *NewsController) Save(ctx *gin.Context) {
	var (
		data  requests.NewsRequest
		err   error
		model models.News
		re    = repositorys.NewNewsRepository()
	)
	err = ctx.ShouldBind(&data)
	if err != nil {
		response.Failed(ctx, global.GetError(err, data))
		return
	}

	model, err = re.Insert(data)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}

func (controller *NewsController) Edit(ctx *gin.Context) {
	var (
		err          error
		id           int
		request      requests.NewsRequest
		re           = repositorys.NewNewsRepository()
		rowsAffected int64
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	err = ctx.ShouldBind(&request)
	if err != nil {
		response.Failed(ctx, global.GetError(err, request))
		return
	}
	rowsAffected, err = re.UpdateById(id, request)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
}

func (controller *NewsController) Delete(ctx *gin.Context) {

	var (
		ids          requests.DeleteNewsRequest
		err          error
		rowsAffected int64
		re           = repositorys.NewNewsRepository()
	)

	err = ctx.ShouldBind(&ids)
	if err != nil {
		response.Failed(ctx, global.GetError(err, ids))
		return
	}
	rowsAffected, err = re.DelByIds(ids.Ids)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
	return
}

func (controller *NewsController) Get(ctx *gin.Context) {

	var (
		err   error
		id    int
		model models.News
		re    = repositorys.NewNewsRepository()
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	model, err = re.FindById(id)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}
