package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func FindUserIDFromContext(c *gin.Context) (uint,error) {
	claims, exists := c.Get("userID")
	if !exists {
		return 0,errors.New("claims not found")
	}
	userID, ok := claims.(uint)
	if !ok {
		return  0,errors.New("invalid user id type")
	}
	return userID,nil
}
