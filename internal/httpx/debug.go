package httpx

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	}
}

func debugPrintRoute(httpMethod, absolutePath string, nuMiddlewares int, handlers ...HandlerFunc) {
	if gin.IsDebugging() {
		nuHandlers := len(handlers)
		if nuHandlers == 0 {
			return
		}
		handlerName := nameOfFunction(handlers[len(handlers)-1])
		debugPrint("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuMiddlewares+nuHandlers)

	}
}

func combineHandlers(handlers ...HandlerFunc) []HandlerFunc {
	finalSize := len(handlers)
	mergedHandlers := make([]HandlerFunc, finalSize)
	copy(mergedHandlers, handlers)
	return mergedHandlers
}

func calculateAbsolutePath(basePath, relativePath string) string {
	return joinPaths(basePath, relativePath)
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func nameOfFunction(f any) string {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func debugPrint(format string, values ...any) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
	}
}
