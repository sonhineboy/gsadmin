package global

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadmin/service/config"
	"github.com/sonhineboy/gsadmin/service/src"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"time"
)

var (
	GAD_R           *gin.Engine
	GAD_APP_PATH    string
	Config          *config.Config
	Db              *gorm.DB
	SuperAdmin      string
	EventDispatcher src.EventDispatcher
	Limiter         *rate.Limiter
)

func GetError(errs validator.ValidationErrors, r interface{}) string {
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_msg"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("msg")
		if errTagText != "" {
			return errTagText
		}
		if errText != "" {
			return errText
		}
		return fieldError.Field() + ":" + fieldError.Tag()
	}
	return ""
}

/**
分页
*/
func Pages(page int, pageSize int, total int, rows interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	data["page"] = page
	data["pageSize"] = pageSize
	data["rows"] = rows
	data["total"] = total
	return data
}

//即将废弃，请勿使用
func IsSuperAdmin(roles []string, role string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}

	}
	return false
}

//验证码
func CaptchaServe(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func GetEventDispatcher(c *gin.Context) *src.EventDispatcher {

	v, ok := c.Get("e")

	if ok == false {
		fmt.Print("无法获取对象")
		return nil
	}

	e, ok := v.(src.EventDispatcher)

	if ok == false {
		fmt.Print("类型不正确")
		return nil
	}

	return &e
}
