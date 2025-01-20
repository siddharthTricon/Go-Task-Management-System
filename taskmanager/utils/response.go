package utils

import "github.com/gin-gonic/gin"

// func RespondJSON(c *gin.Context, status int, payload interface{}){
// 	c.JSON(status, payload)
// }

func RespondJSON(c *gin.Context, statusCode int, message string, data interface{}) {
    c.JSON(statusCode, gin.H{
        "message": message,
        "data":    data,
    })
}

