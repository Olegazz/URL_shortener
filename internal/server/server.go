package server

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// NewServer создает новый экземпляр Server
func NewServer(addr string) *Server {
	mux := http.NewServeMux()

	//shortenHandler := handlers.NewShortenHandler(storage)
	//mux.HandleFunc("/shorten", shortenHandler.HandleShorten)

	//redirectHandler := handlers.NewRedirectHandler(storage)
	//mux.HandleFunc("/redirect", redirectHandler.HandleRedirect)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &Server{httpServer: httpServer}
}

func (s *Server) Start() error {
	log.Printf("Server started on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Println("Server stopped")
	return s.httpServer.Shutdown(nil)
}
