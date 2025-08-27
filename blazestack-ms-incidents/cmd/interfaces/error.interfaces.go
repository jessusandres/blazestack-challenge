package interfaces

import "github.com/gin-gonic/gin"

type IAborter interface {
	Abort()
	Error(err error) *gin.Error
}
