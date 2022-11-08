package global

import (
	"bytes"
	"ginedu2/service/config"
	"ginedu2/service/src"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	EventDispatcher *src.EventDispatcher
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
