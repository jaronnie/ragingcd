package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(g *gin.Context, data interface{}) {
	g.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "message": "success"})
}

func Fail(g *gin.Context, err error, code int) {
	rCode := http.StatusInternalServerError
	if http.StatusText(code) != "" {
		rCode = code
	}
	if err != nil {
		g.JSON(http.StatusOK, gin.H{"code": rCode, "data": nil, "message": err.Error()})
	} else {
		g.JSON(http.StatusOK, gin.H{"code": rCode, "data": nil, "message": http.StatusText(rCode)})
	}
}
