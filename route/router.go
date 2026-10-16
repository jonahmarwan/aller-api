package route

import (
	api "github.com/jonahmarwan/aller-api"
	"errors"
	"net/http"

	"github.com/armon/go-radix"
)

type Router struct {
	tree   *radix.Tree
	Config *api.ConfigStruct
}

func NewRouter(config *api.ConfigStruct) (Router, error) {
	var router Router
	router.tree = radix.New()
	router.Config = config
	return router, nil
}

func (router *Router) AddRoute(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	router.tree.Insert(path, handlerFunc)
}

// HandleRequest in the future take Request type as input
func (router *Router) HandleRequest(path string) (func(w http.ResponseWriter, r *http.Request), error) {
	m, _ := router.tree.Get(path)
	if m == nil {
		err := errors.New("404 NOT FOUND")
		return nil, err
	}
	fn := m.(func(w http.ResponseWriter, r *http.Request))
	return fn, nil
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, err := router.HandleRequest(r.URL.Path)
	if err != nil {
		return
	}
	handler(w, r)
}

func (router *Router) Self() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if router.Config.CORS {
			w.Header().Set("Access-Control-Allow-Origin", "https://google.com")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		router.ServeHTTP(w, r)
	})
}

func (router *Router) EnableCORS() {
	router.Config.CORS = true
}

func (router *Router) DisableCORS() {
	router.Config.CORS = false
}
