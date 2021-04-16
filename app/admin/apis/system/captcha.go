package system

import (
	"go-admin/common/apis"

	"github.com/alphayan/go-admin-core/sdk/pkg/captcha"
	"github.com/gin-gonic/gin"
)

type System struct {
	apis.Api
}

func (e *System) GenerateCaptchaHandler(c *gin.Context) {
	log := e.GetLogger(c)
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		log.Errorf("DriverDigitFunc error, %s", err.Error())
		e.Error(c, 500, err, "验证码获取失败")
		return
	}
	e.Custom(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
