package routers

type RouterGroup struct {
	Auth AuthRouter
	User UserRouter
}

var RouterGroupApp = new(RouterGroup)
