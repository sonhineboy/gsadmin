package requests

import "github.com/sonhineboy/gsadmin/service/global"

type UserList struct {
	global.List
}

type UserUpdate struct {
	Id        uint   `json:"id"`
	Name      string `json:"name" binding:"required" msg:"用户名不能为空"`
	PassWord  string `json:"password"`
	CPassWord string `json:"password2"`
	RealName  string `json:"real_name" binding:"required,min=2" min_msg:"长度最小大于2" msg:"真实姓名不能为空"`
	Avatar    string `json:"avatar" binding:"required,min=3" min_msg:"长度最小大于3" msg:"通向不能为空"`
	Roles     []uint `json:"group"`
}
