package middleware

import (
	"fmt"
	"gin-basic-framework/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法和密钥
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			key := global.GLOBAL_CONFIG.JWT.SigningKey
			return key, nil
		})

		if err != nil {
			// Token 验证失败，返回错误响应
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token 验证成功，可以从 claims 中获取用户信息
			userID := claims["userID"].(string)
			// 将 userID 存储到上下文中，方便后续处理
			c.Set("userID", userID)
			c.Next()
		} else {
			// Token 验证失败，返回错误响应
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
	}
}
