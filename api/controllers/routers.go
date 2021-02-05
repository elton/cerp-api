package controllers

import "github.com/elton/cerp-api/api/middlewares"

func (s *Server) initializeRouters() {
	v1 := s.Router.Group("api/v1")
	{
		v1.GET("/status", middlewares.SetMiddlewareJSON(), HealthCheck)
	}
}
