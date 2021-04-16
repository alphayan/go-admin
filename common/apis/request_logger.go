package apis

import (
	"strings"

	"github.com/alphayan/go-admin-core/sdk"
	"github.com/alphayan/go-admin-core/sdk/pkg"
	"github.com/alphayan/go-admin-core/sdk/pkg/logger"
	"github.com/gin-gonic/gin"
)

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Logger {
	requestId := pkg.GenerateMsgIDFromContext(c)
	log := sdk.Runtime.GetLogger().Fields(map[string]interface{}{
		strings.ToLower(pkg.TrafficKey): requestId,
	})
	return &logger.Logger{Logger: log}
}
