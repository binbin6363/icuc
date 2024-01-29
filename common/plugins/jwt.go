package plugins

import (
	"context"
	"net/http"
	"strings"

	"github.com/binbin6363/icuc/common/api"
	"github.com/binbin6363/icuc/common/err"
	"github.com/binbin6363/icuc/common/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var defaultOpt *Options

func init() {
	defaultOpt = &Options{}
}

type Options struct {
	SkipPaths []string // 跳过token校验的方法
	Secret    string   // jwt的secret
}

type Option func(*Options)

// InitOption 初始化传入Option
func InitOption(opts ...Option) {
	for _, o := range opts {
		o(defaultOpt)
	}
}

// WithSecret jwt的secret
func WithSecret(secret string) Option {
	return func(o *Options) {
		o.Secret = secret
	}
}

// WithSkipPaths 跳过验证的请求path
func WithSkipPaths(path string) Option {
	return func(o *Options) {
		o.SkipPaths = append(o.SkipPaths, path)
	}
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过不校验token的请求
		for _, path := range defaultOpt.SkipPaths {
			if strings.Contains(c.Request.URL.Path, path) {
				return
			}
		}

		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get(api.AuthField)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": err.CodeNoAuth,
				"msg":  "没有认证信息",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Fields(authHeader)
		if !(len(parts) == 2 && parts[0] == api.AuthBearerField) {
			c.JSON(http.StatusOK, gin.H{
				"code": 4001,
				"msg":  "认证信息鉴权失败",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := parseToken(c, parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 4002,
				"msg":  "无效的认证信息",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(api.HeadUid, mc.Id)
		c.Set(api.HeadUserName, mc.Audience)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func parseToken(ctx context.Context, token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(defaultOpt.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}

	log.InfoContextf(ctx, "token invalid, token:%s, jwtToken:%+v", token, jwtToken)
	return nil, err
}
