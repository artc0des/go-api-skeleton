package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"template.com/api/utils"
)

func Authenticate(context *gin.Context) {
	//JWT validation
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	isValid, userType, userId, jwtErr := utils.ValidateToken(token)

	if jwtErr != nil || !isValid {
		context.AbortWithError(http.StatusUnauthorized, errors.New("unathorized user"))
		return
	}

	context.Set("userType", userType)
	context.Set("userId", userId)
}
