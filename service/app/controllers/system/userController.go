package system

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"net/http"
)

type UserController struct{}

//登陆
func (u *UserController) Login(c *gin.Context) {
	var (
		LoginForm      requests.Login
		roles          []string
		permission     []string
		userRepository repositorys.UserRepository
		menuRepository repositorys.SystemMenuRepository
	)
	apiList := make(map[string]string)

	err := c.ShouldBind(&LoginForm)

	if !captcha.VerifyString(LoginForm.CaptchaId, LoginForm.CaptchaValue) {
		response.Failed(c, "验证码错误")
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "422",
			"error": global.GetError(err.(validator.ValidationErrors), LoginForm),
		})
		return
	}
	isLogin, user := userRepository.Login(LoginForm.PassWord, LoginForm.Name, c)

	if isLogin {

		for _, role := range user.Roles {
			roles = append(roles, role.Alias)
		}

		_ = menuRepository.GetApiListToMapByUser(user, &apiList)
		_ = menuRepository.GetPermissionByUser(user, &permission)

		token, _ := models.GenToken(models.JwtUser{}.NewJwtUser(
			user.ID,
			user.Name,
			roles,
			apiList,
			permission,
		), global.Config.MyJwt.Secret)
		response.Success(c, "登陆成功", gin.H{
			"token":    token,
			"userInfo": user,
		})
	} else {
		response.Failed(c, "用户名或密码错误")
	}

}

//注册用户
func (u *UserController) Add(c *gin.Context) {
	var (
		userAdd        requests.UserAdd
		userRepository repositorys.UserRepository
	)
	err := c.ShouldBind(&userAdd)
	if err != nil {
		response.Failed(c, global.GetError(err.(validator.ValidationErrors), userAdd))
		return
	}
	result, model := userRepository.Add(userAdd.PassWord, userAdd.Name, userAdd)

	if result.Error == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (u *UserController) List(c *gin.Context) {
	var (
		params         requests.UserList
		userRepository repositorys.UserRepository
	)
	_ = c.ShouldBind(&params)

	userRepository.Where = params.Where
	response.Success(c, "ok", userRepository.List(params.Page, params.PageSize, "created_at"))
}

func (u *UserController) Up(c *gin.Context) {

	var (
		data           requests.UserUpdate
		userRepository repositorys.UserRepository
	)
	err := c.ShouldBind(&data)
	if err != nil {
		response.Failed(c, global.GetError(err.(validator.ValidationErrors), data))
		return
	}

	reErr := userRepository.Update(data)
	if reErr == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, reErr.Error())
	}
}

func (u *UserController) Del(c *gin.Context) {
	var delIds global.Del
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminUser{}, delIds.Ids)
	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}
