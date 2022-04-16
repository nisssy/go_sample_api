package controllers

type Context interface {
	// Request bind method
	Bind(obj interface{}) error
	// Return type
	Param(key string) string
	JSON(code int, obj interface{})
}
