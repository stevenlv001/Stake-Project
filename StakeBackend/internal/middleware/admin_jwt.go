package middleware

import (
	"StakeBackend/internal/config"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AdminClaims 管理员JWT Claims
type AdminClaims struct {
	AdminID   string `json:"admin_id"`
	Role      string `json:"role"` // admin, super_admin
	jwt.RegisteredClaims
}

// GenerateAdminToken 生成管理员JWT
func GenerateAdminToken(adminID string, role string) (string, error) {
	claims := AdminClaims{
		AdminID: adminID,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GlobalConfig.JWT.Expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GlobalConfig.JWT.Secret + "_admin"))
}

// AdminJWTMiddleware 管理员JWT鉴权中间件
func AdminJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未授权"})
			c.Abort()
			return
		}

		claims, err := ParseAdminToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token无效"})
			c.Abort()
			return
		}

		// 设置管理员信息到上下文
		c.Set("admin_id", claims.AdminID)
		c.Set("admin_role", claims.Role)
		c.Next()
	}
}

// ParseAdminToken 解析管理员Token
func ParseAdminToken(token string) (*AdminClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法类型，防止算法混淆攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 使用不同的密钥区分管理员和普通用户
		return []byte(config.GlobalConfig.JWT.Secret + "_admin"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaims.Claims.(*AdminClaims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid admin token claims")
}

// RequireRole 角色权限检查中间件
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("admin_role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "无权访问"})
			c.Abort()
			return
		}

		// 超级管理员拥有所有权限
		if roleStr, ok := role.(string); ok {
			if roleStr == "super_admin" || roleStr == requiredRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "权限不足"})
		c.Abort()
	}
}
