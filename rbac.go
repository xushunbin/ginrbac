package ginrbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type RBAC struct {
	Ctx   *gin.Context
	DB    *gorm.DB
	Redis *redis.Client
}

// New 获取RBAC实例
func New(ctx *gin.Context, db *gorm.DB, rds *redis.Client) *RBAC {
	return &RBAC{
		Ctx: ctx,
		DB: db,
		Redis: rds,
	}
}

// Middleware gin中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("RBAC Middleware 123", c.FullPath())
		c.Next()
	}
}