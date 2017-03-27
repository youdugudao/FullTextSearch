package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errJson struct {
	Success   bool
	ErrorInfo string
	DebugInfo interface{}
}

type okJson struct {
	Success bool
	Data    interface{}
}

func CheckErr(err error, c *gin.Context, debugInfo interface{}) {
	if err != nil {
		ResponseError(err.Error(), c, debugInfo)
	}
}

func ResponseError(errString string, c *gin.Context, debugInfo interface{}) {
	c.JSON(http.StatusOK, errJson{
		Success:   false,
		ErrorInfo: errString,
		DebugInfo: debugInfo,
	})
	c.Abort()
}

func ResponseOk(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, okJson{
		Success: true,
		Data:    data,
	})
	c.Abort()
}
