// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package httpx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Context 自定义Context,扩展gin.Context
type Context struct {
	*gin.Context // 继承gin.Context
}

var _ context.Context = &Context{}

// NewContext 创建自定义Context
func NewContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}

// // Success 成功响应
// func (c *Context) Success(data interface{}) {
// 	c.ginCtx.JSON(http.StatusOK, gin.H{
// 		"code": 0,
// 		"msg":  "success",
// 		"data": data,
// 	})
// }

// // Error 错误响应
// func (c *Context) Error(code int, msg string) {
// 	c.ginCtx.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"msg":  msg,
// 	})
// }

// // BadRequest 400错误响应
// func (c *Context) BadRequest(msg string) {
// 	c.ginCtx.JSON(http.StatusBadRequest, gin.H{
// 		"code": http.StatusBadRequest,
// 		"msg":  msg,
// 	})
// }

// // Unauthorized 401错误响应
// func (c *Context) Unauthorized(msg string) {
// 	c.ginCtx.JSON(http.StatusUnauthorized, gin.H{
// 		"code": http.StatusUnauthorized,
// 		"msg":  msg,
// 	})
// }

// // Forbidden 403错误响应
// func (c *Context) Forbidden(msg string) {
// 	c.ginCtx.JSON(http.StatusForbidden, gin.H{
// 		"code": http.StatusForbidden,
// 		"msg":  msg,
// 	})
// }

// // NotFound 404错误响应
// func (c *Context) NotFound(msg string) {
// 	c.ginCtx.JSON(http.StatusNotFound, gin.H{
// 		"code": http.StatusNotFound,
// 		"msg":  msg,
// 	})
// }

// // ServerError 500错误响应
// func (c *Context) ServerError(msg string) {
// 	c.ginCtx.JSON(http.StatusInternalServerError, gin.H{
// 		"code": http.StatusInternalServerError,
// 		"msg":  msg,
// 	})
// }
