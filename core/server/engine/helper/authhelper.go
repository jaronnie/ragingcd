package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
)

type AuthHelper struct {
	Context *gin.Context
}

func (ah *AuthHelper) UserID() int {
	return ah.Context.MustGet(middlewares.XForwardAuthHeaderKey).(middlewares.XForwardAuth).UserId
}
