package system

import (
	"ginedu2/service/app/models"
	"ginedu2/service/app/repositorys"
	"ginedu2/service/app/requests"
	"ginedu2/service/global"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
	res            global.Response
	userRepository repositorys.UserRepository
	menuRepository repositorys.SystemMenuRepository
}

//登陆
func (u UserController) Login(c *gin.Context) {
	var (
		LoginForm  requests.Login
		roles      []string
		permission []string
	)
	apiList := make(map[string]string)

	err := c.ShouldBind(&LoginForm)

	if !captcha.VerifyString(LoginForm.CaptchaId, LoginForm.CaptchaValue) {
		u.res.Failed(c, "验证码错误")
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "422",
			"error": global.GetError(err.(validator.ValidationErrors), LoginForm),
		})
		return
	}
	isLogin, user := u.userRepository.Login(LoginForm.PassWord, LoginForm.Name)

	if isLogin {

		for _, role := range user.Roles {
			roles = append(roles, role.Alias)
		}

		_ = u.menuRepository.GetApiListToMapByUser(user, &apiList)
		_ = u.menuRepository.GetPermissionByUser(user, &permission)

		token, _ := models.GenToken(models.JwtUser{}.NewJwtUser(
			user.ID,
			user.Name,
			roles,
			apiList,
			permission,
		), global.Config.MyJwt.Secret)
		u.res.Success(c, "登陆成功", gin.H{
			"token":    token,
			"userInfo": user,
		})
	} else {
		u.res.Failed(c, "用户名或密码错误")
	}

}

//注册用户
func (u UserController) Add(c *gin.Context) {
	var userAdd requests.UserAdd
	err := c.ShouldBind(&userAdd)
	if err != nil {
		u.res.Failed(c, global.GetError(err.(validator.ValidationErrors), userAdd))
		return
	}
	result, model := u.userRepository.Add(userAdd.PassWord, userAdd.Name, userAdd)

	if result.Error == nil {
		u.res.Success(c, "ok", model)
	} else {
		u.res.Failed(c, result.Error.Error())
	}
}

func (u UserController) List(c *gin.Context) {
	var params requests.UserList
	_ = c.ShouldBind(&params)

	u.userRepository.Where = params.Where
	u.res.Success(c, "ok", u.userRepository.List(params.Page, params.PageSize, "created_at"))
}

func (u UserController) Up(c *gin.Context) {

	var data requests.UserUpdate
	err := c.ShouldBind(&data)
	if err != nil {
		u.res.Failed(c, global.GetError(err.(validator.ValidationErrors), data))
		return
	}

	reErr := u.userRepository.Update(data)
	if reErr == nil {
		u.res.Success(c, "ok", "")
	} else {
		u.res.Failed(c, reErr.Error())
	}
}

func (u UserController) Del(c *gin.Context) {
	var delIds global.Del
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminUser{}, delIds.Ids)
	if result.Error == nil {
		u.res.Success(c, "ok", "")
	} else {
		u.res.Failed(c, result.Error.Error())
	}
}
