package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	c.Set("error", err)
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func Success(c *gin.Context, response interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
		"data":    response,
	})
}
