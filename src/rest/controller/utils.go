package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MakeErrorJSON(httpCode int, errorCode int, msg any) (int, any) {
	return httpCode, &gin.H{
		"error": errorCode,
		"msg":   fmt.Sprint(msg),
	}
}

func MakeSuccessJSON(data any, msg string) (int, any) {
	return 200, &gin.H{
		"error": 0,
		"msg":   msg,
		"data":  data,
	}
}
