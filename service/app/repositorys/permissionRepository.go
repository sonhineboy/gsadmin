package repositorys

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/app/models"
)

type PermissionRepository struct {
	CustomClaims *models.CustomClaims
}

func NewPermissionRepository(customClaims *models.CustomClaims) *PermissionRepository {
	return &PermissionRepository{
		CustomClaims: customClaims,
	}
}

func NewDefaultPermissionRepository(c *gin.Context) (*PermissionRepository, error) {
	custom, ok := GetCustomClaims(c)
	if ok {
		return &PermissionRepository{
			CustomClaims: custom,
		}, nil
	} else {
		return nil, errors.New("初始化失败")
	}
}

func GetCustomClaims(c *gin.Context) (*models.CustomClaims, bool) {
	v, ok := c.Get("claims")
	claims, err := v.(*models.CustomClaims)
	if ok && err {
		return claims, true
	} else {
		return &models.CustomClaims{}, false
	}
}

func GetPermission(c *gin.Context) *PermissionRepository {
	v, ok := c.Get("permission")
	permission, err := v.(*PermissionRepository)
	if ok && err {
		return permission
	} else {
		return nil
	}
}

// IsRole 判断是否某个用户组
func (p *PermissionRepository) IsRole(role string) bool {
	for _, v := range p.CustomClaims.Roles {
		if v == role {
			return true
		}
	}
	return false
}

//判断是否纯在否个权限
func (p *PermissionRepository) isHasPermission(permission string) bool {
	for _, v := range p.CustomClaims.Permission {
		if v == permission {
			return true
		}
	}
	return false
}

//判断是有api权限
func (p *PermissionRepository) isApi(url string, code string) bool {

	v, ok := p.CustomClaims.ApiList[url]

	if ok == false {
		return false
	}

	if v != code {
		return false
	}
	return true
}
