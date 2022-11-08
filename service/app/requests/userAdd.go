package requests

type UserAdd struct {
	Name      string `json:"name" binding:"required" msg:"用户名不能为空"`
	PassWord  string `json:"password" binding:"required,min=3,eqfield=CPassWord" min_msg:"长度最小大于3" eqfield_msg:"两次输入密码不一致" msg:"密码不能为空"`
	CPassWord string `json:"password2" binding:"required,min=3" min_msg:"长度最小大于3" msg:"密码不能为空"`
	RealName  string `json:"real_name" binding:"required,min=2" min_msg:"长度最小大于2" msg:"真实姓名不能为空"`
	Avatar    string `json:"avatar" binding:"required,min=3" min_msg:"长度最小大于3" msg:"通向不能为空"`
	Roles     []uint `json:"group"`
}
