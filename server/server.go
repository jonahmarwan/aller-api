package server

import (
	"errors"
	"net/http"

	api "github.com/jonahmarwan/aller-api"
	"github.com/jonahmarwan/aller-api/crypt"
)

type Server struct {
	config api.ConfigStruct
	router Router	
}

func NewServer(cfg *api.ConfigStruct, rtr *Router) (Server, error){
	if cfg.AutoKeys == true {
		pubKey, privKey, err := crypt.GenerateKeys()
		if err != nil {
			return Server{
				config: api.ConfigStruct{CORS: true}, 
				router: *rtr}, errors.New("Couldn't generate API keys.")
		}
		cfg.ApiPrivateKey = privKey
		cfg.ApiPublicKey = pubKey
	} else {
		if cfg.ApiPrivateKey == [32]byte{0} || cfg.ApiPublicKey == [32]byte{0} {
			return Server{
				config: api.ConfigStruct{CORS: true}, 
				router: *rtr}, errors.New("API Keys cannot be None while Autokeys is false.")
			
		}
	}

	var s Server
	s.config = *cfg
	s.router  = *rtr
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, err := s.router.HandleRequest(r.URL.Path)
	if err != nil {
		return
	}
	handler(w, r)
}

	
func (s *Server) Self() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.router.Config.CORS {
			w.Header().Set("Access-Control-Allow-Origin", "https://google.com")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		s.router.ServeHTTP(w, r)
	})
}