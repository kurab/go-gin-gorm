package controller

// more, see https://godoc.org/github.com/gin-gonic/gin
type Context interface {
	Param(key string) string
	BindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	String(code int, format string, values ...interface{})
}
