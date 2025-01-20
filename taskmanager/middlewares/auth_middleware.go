package middlewares

import(
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/siddharthTricon/go-task-management-sysytem/utils"
)

func AuthMiddleware() gin.HandlerFunc{
    return func(c *gin.Context){
        token := c.GetHeader("Authorization")
        if token == "" || !utils.ValidateToken(token){
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}