package requests

import (
	"github.com/sonhineboy/gsadminGen/pkg"
)

type GenFields struct {
	TableName string `form:"table_name" binding:"required" required_msg:"当前字段必填"`
}

type GenCode struct {
	Fields            []pkg.Field `json:"fields" binding:"required"`
	Checkbox          []string    `json:"checkbox" binding:"required"`
	ControllerPackage string      `json:"controllerPackage" binding:"required"`
	TableDiyName      string      `json:"tableDiyName" binding:"required"`
}
