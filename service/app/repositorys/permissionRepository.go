package repositorys

import (
	"errors"
	"ginedu2/service/app/models"
	"github.com/gin-gonic/gin"
	"sync"
)

type PermissionRepository struct {
	CustomClaims *models.CustomClaims
}

var (
	permission *PermissionRepository
	once       = &sync.Once{}
)

func newPermissionRepository(customClaims *models.CustomClaims) *PermissionRepository {
	return &PermissionRepository{
		CustomClaims: customClaims,
	}
}

func newDefaultPermissionRepository(c *gin.Context) (*PermissionRepository, error) {
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
	if permission == nil {
		once.Do(func() {
			customClaims, ok := GetCustomClaims(c)
			if ok == false {
				panic("无法获取授权信息")
			}
			permission = newPermissionRepository(customClaims)
		})
	}
	return permission
}

//判断是否某个用户组
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
