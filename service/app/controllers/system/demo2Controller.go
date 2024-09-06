package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"strconv"
)

type Demo2Controller struct{}

func (controller *Demo2Controller) Index(ctx *gin.Context) {

	var (
		params global.List
		re = repositorys.NewDemo2Repository()
	)
	_ = ctx.ShouldBindBodyWith(&params, binding.JSON)
	response.Success(ctx, "ok", re.Page(params.Where, params.Page, params.PageSize, "created_at"))
}

func (controller *Demo2Controller) Save(ctx *gin.Context) {
	var (
		data  requests.Demo2Request
		err   error
		model models.Demo2
		re    = repositorys.NewDemo2Repository()
	)
	err = ctx.ShouldBindBodyWith(&data,binding.JSON)
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

func (controller *Demo2Controller) Edit(ctx *gin.Context) {
	var (
		err          error
		id           int
		request      requests.Demo2Request
		re           = repositorys.NewDemo2Repository()
		rowsAffected int64
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	err = ctx.ShouldBindBodyWith(&request, binding.JSON)
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

func (controller *Demo2Controller) Delete(ctx *gin.Context) {

	var (
		ids          requests.DeleteDemo2Request
		err          error
		rowsAffected int64
		re           = repositorys.NewDemo2Repository()
	)

	err = ctx.ShouldBindBodyWith(&ids, binding.JSON)
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

func (controller *Demo2Controller) Get(ctx *gin.Context) {

	var (
		err   error
		id    int
		model models.Demo2
		re    = repositorys.NewDemo2Repository()
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
