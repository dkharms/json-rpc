package api

type Path string

type Route struct {
	path    Path
	handler handler
}
