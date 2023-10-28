package controllers

import "github.com/lsortudo/blog-golang/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET") // Corrigir problema la no middlewares so ta reconhecendo o auth

}
