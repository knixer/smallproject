package api

func (s *Server) Routes() {
	s.Router.GET("/car/:brand", s.handleCarGet())
}
