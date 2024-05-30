package requests

type Login struct {
	Name     string `form:"username" binding:"required" msg:"用户名不能为空" json:"username"`
	PassWord string `form:"password" binding:"required" msg:"密码不能为空" json:"password"`

	CaptchaId    string `json:"captchaId"`
	CaptchaValue string `json:"captchaValue"`
}
