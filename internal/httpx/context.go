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
