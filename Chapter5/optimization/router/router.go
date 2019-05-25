package router

import (
	"log"
	"net/http"
)

type Routers struct {
	tree map[string]*node
}

type node struct {
	childNode []*node
	Method    string
	Path      string
	Priority  int
	Handle    Handle
}

func (n *node) addNode(path string, handler Handle) {}
func (n *node) getNode(path string) (Handle, Params) {
	var a Handle
	return a, Params{}
}

func New() *Routers {
	return &Routers{
		tree: make(map[string]*node),
	}
}

type Handle func(http.ResponseWriter, *http.Request, Params)

type HandleFunc func(handler Handle)

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (p Params) Get(key string) string {
	if len(p) == 0 {
		return ""
	}
	for _, param := range p {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

func (r *Routers) GET(path string, h Handle) {
	r.Handle(http.MethodGet, path, h)
}

func (r *Routers) POST(path string, h Handle) {
	r.Handle(http.MethodPost, path, h)

}

func (r *Routers) PATCH(path string, h Handle) {
	r.Handle(http.MethodPatch, path, h)

}

func (r *Routers) PUT(path string, h Handle) {
	r.Handle(http.MethodPut, path, h)

}

func (r *Routers) DELETE(path string, h Handle) {
	r.Handle(http.MethodDelete, path, h)

}

func (r *Routers) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 调用相应的 handle 触发动作
	node := r.tree[request.Method]
	h, params := node.getNode(request.URL.Path)
	h(writer, request, params)
}

func (r *Routers) Handle(method string, path string, h Handle) {
	// 添加路由
	root := r.tree[method]
	root.addNode(path, h)
	log.Println(r.tree)
}
