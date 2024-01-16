package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/jaronnie/ragingcd/core/pkg/restc"

	"github.com/gin-gonic/gin"
)

const XForwardAuthHeaderKey = "X-Forward-Auth-Header"

type XForwardAuth struct {
	UserId int `json:"userId"`
}

func DevAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get XForwardAuthHeader from local auth_url
		var xForwardAuth XForwardAuth

		var client restc.Interface
		var err error

		client, err = restc.New(restc.WithUrl(viper.GetString("server.dev.auth_url")), restc.WithHeaders(c.Request.Header))

		headers, err := client.Get().SubPath("/api/v1.0/system/user/auth").Do(context.Background()).Header()
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		err = json.Unmarshal([]byte(headers.Get(XForwardAuthHeaderKey)), &xForwardAuth)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(XForwardAuthHeaderKey, xForwardAuth)
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	// 不允许设置其他配置文件
	if _, err := os.Stat("config/config-dev.toml"); err == nil {
		viper.SetConfigFile("config/config-dev.toml")
		_ = viper.ReadInConfig()
		if viper.GetString("server.mode") == "dev" {
			return DevAuth()
		}
	}
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
