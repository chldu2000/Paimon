package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaimonAlert(c *gin.Context) {
	c.String(http.StatusOK, "前面的区域，以后再来探索吧！")
}
