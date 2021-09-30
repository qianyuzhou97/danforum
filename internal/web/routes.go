package web

import (
	"net/http"
)

func (s *Server) SetRoutes(test bool) *Server {

	// mid.Logger(sugar)

	// p := PostService{db: db}
	// c := CommunityService{db: db}
	// u := UserService{db: db}
	if !test {
		s.mw = []Middleware{Logger(s.sugar), Errors(s.sugar), Metrics()}
	}

	s.Handle(http.MethodGet, "/posts", s.ListAllPosts, Authenticate())
	s.Handle(http.MethodPost, "/posts", s.CreatePost)
	s.Handle(http.MethodGet, "/posts/{id}", s.GetPostByID)
	s.Handle(http.MethodPut, "/posts/{id}", s.UpdatePostByID)
	s.Handle(http.MethodDelete, "/posts/{id}", s.DeletePostByID)

	s.Handle(http.MethodGet, "/community", s.ListAllCommunity)
	s.Handle(http.MethodPost, "/community", s.CreateCommunity)
	s.Handle(http.MethodGet, "/community/{id}", s.GetCommunityByID)

	s.Handle(http.MethodPut, "/user", s.CreateUser)
	s.Handle(http.MethodGet, "/user/token", s.Authenticate)
	return s
}
