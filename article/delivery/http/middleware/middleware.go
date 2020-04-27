package middleware

import "github.com/labstack/echo"

// GoMiddleware represent the data-struct for middleware
// GoMiddlewareはミドルウェアのデータ構造を表します
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
// CORSはCORSミドルウェアを処理します
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware initialize the middleware
// InitMiddlewareがミドルウェアを初期化します
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
