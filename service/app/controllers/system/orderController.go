package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"strconv"
)

type OrderController struct{}

func (orderController *OrderController) Index(ctx *gin.Context) {

	var (
		params global.List
		re     repositorys.OrderRepository
	)
	_ = ctx.ShouldBind(&params)

	response.Success(ctx, "ok", re.Page(params.Where, params.Page, params.PageSize, "created_at"))
}

func (orderController *OrderController) Save(ctx *gin.Context) {
	var (
		data  requests.OrderRequest
		err   error
		model models.Order
		re    = repositorys.NewOrderRepository()
	)
	err = ctx.ShouldBind(&data)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, data))
		} else {
			response.Failed(ctx, err.Error())
		}
		return
	}

	model, err = re.Insert(data)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}

func (orderController *OrderController) Edit(ctx *gin.Context) {
	var (
		err          error
		id           int
		request      requests.OrderRequest
		re           = repositorys.NewOrderRepository()
		rowsAffected int64
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	err = ctx.ShouldBind(&request)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, request))
		} else {
			response.Failed(ctx, err.Error())
		}
		return
	}
	rowsAffected, err = re.UpdateById(id, request)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
}

func (orderController *OrderController) Delete(ctx *gin.Context) {

	var (
		ids          requests.DeleteOrderRequest
		err          error
		rowsAffected int64
		re           = repositorys.NewOrderRepository()
	)

	err = ctx.ShouldBind(&ids)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, ids))
		} else {
			response.Failed(ctx, err.Error())
		}
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

func (orderController *OrderController) Get(ctx *gin.Context) {

	var (
		err   error
		id    int
		model models.Order
		re    = repositorys.NewOrderRepository()
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
