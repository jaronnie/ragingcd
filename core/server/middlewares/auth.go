package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const XForwardAuthHeaderKey = "X-Forward-Auth-Header"

type XForwardAuth struct {
	UserId int `json:"userId"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get(XForwardAuthHeaderKey)
		if auth == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  fmt.Sprintf("%s 为空", XForwardAuthHeaderKey),
			})
			c.Abort()
			return
		}
		var xForwardAuth XForwardAuth
		err := json.Unmarshal([]byte(auth), &xForwardAuth)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  fmt.Sprintf("%s 格式不正确", XForwardAuthHeaderKey),
			})
			c.Abort()
			return
		}
		c.Set(XForwardAuthHeaderKey, xForwardAuth)
		c.Next()
	}
}
