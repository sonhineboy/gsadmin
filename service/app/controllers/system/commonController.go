package system

import (
	"ginedu2/service/global"
	"github.com/dchest/captcha"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
)

type CommonController struct {
	res            global.Response
	allowType      map[string]string
	uploadPath     string
	uploadPathBase string
}

func NewCommonController() *CommonController {
	return &CommonController{
		allowType: map[string]string{
			"image/jpeg": "jpg",
			"image/png":  "png",
		},
		uploadPathBase: "." + global.Config.App.UploadFile,
		uploadPath:     "." + global.Config.App.UploadFile + "/" + time.Now().Format("20060102"),
	}
}
func (p *CommonController) GetFileBasePath() string {
	return p.uploadPathBase
}

func (p *CommonController) UpLoad(c *gin.Context) {

	file, _ := c.FormFile("file")
	fileType, ok := p.allowType[file.Header.Get("Content-Type")]
	if !ok {
		p.res.Failed(c, "当前类型不允许上传！")
		return
	}

	uuid := uuid2.NewV4()

	dirErr := os.MkdirAll(p.uploadPath, os.ModePerm)

	if dirErr != nil {
		p.res.Failed(c, "文件目录创建错误:"+dirErr.Error())
		return
	}

	fileName := uuid.String() + "." + fileType

	allDir := p.uploadPath + "/" + fileName
	err := c.SaveUploadedFile(file, allDir)

	if err != nil {
		p.res.Failed(c, err.Error())
		return
	}
	p.res.Success(c, "ok", gin.H{
		"id":       uuid,
		"fileName": fileName,
		"src":      global.Config.App.Host + "/api/system/common/file" + allDir[8:],
	})

}

func (p CommonController) CaptchaInfo(c *gin.Context) {
	id := captcha.NewLen(5)
	p.res.Success(c, "ok", gin.H{
		"id":  id,
		"url": global.Config.App.Host + "/api/common/captcha/img/" + string(id),
	})
}

func (p CommonController) CaptchaImage(c *gin.Context) {
	w, _ := strconv.Atoi(c.Param("w"))
	h, _ := strconv.Atoi(c.Param("h"))
	_ = global.CaptchaServe(c.Writer, c.Request, c.Param("id"), ".png", "zh", false, w, h)
}
