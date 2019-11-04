package restful

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gy-kim/search-service/logging"
)

// Server is the HTTP REST server
type Server struct {
	address string
	server  *http.Server
	cfg     Config

	handlerList     http.Handler
	handlerNotFound http.HandlerFunc
}

// Listen starts HTTP service.
func (s *Server) Listen(stop <-chan struct{}) {
	s.logger().Info("Start server..")
	router := s.route()

	// create http server.
	s.server = &http.Server{
		Handler: router,
		Addr:    s.address,
	}

	go func() {
		defer func() {
			_ = s.server.Close()
		}()

		<-stop
	}()

	s.logger().Info("[SERVER] Server Address: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		s.logger().Info("[SERVER] Stop server: %s", err)
	}
}

// route registes url router
func (s *Server) route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", health).Methods("GET")
	router.NotFoundHandler = s.handlerNotFound

	sub := router.PathPrefix("v1").Subrouter()
	sub.Handle("/", s.handlerList).Methods("GET")

	router.Use(s.middleware)

	return router
}

func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (s *Server) logger() logging.Logger {
	return s.cfg.Logger()
}
