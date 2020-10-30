package RGin

type Router struct {
	Path    string
	tree    []*Router
	Method  []string
	Execute []Handle
}

type Handle func(ctx RContext)

func AddRoute(pattern, method string, h Handle) {

}

func (r Router) AddGroup(pattern string) *Router {

}

func (r Router) Use(h Handle) {

}
