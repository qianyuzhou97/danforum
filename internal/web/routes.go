package web

import (
	"net/http"
)

func (s *Server) routes() {

	// mid.Logger(sugar)

	// p := PostService{db: db}
	// c := CommunityService{db: db}
	// u := UserService{db: db}

	s.Handle(http.MethodGet, "/v1/posts", s.ListAllPosts, Authenticate())
	s.Handle(http.MethodPost, "/v1/posts", s.CreatePost)
	s.Handle(http.MethodGet, "/v1/posts/{id}", s.GetPostByID)
	s.Handle(http.MethodPut, "/v1/posts/{id}", s.UpdatePostByID)
	s.Handle(http.MethodDelete, "/v1/posts/{id}", s.DeletePostByID)

	s.Handle(http.MethodGet, "/v1/community", s.ListAllCommunity)
	s.Handle(http.MethodPost, "/v1/community", s.CreateCommunity)
	s.Handle(http.MethodGet, "/v1/community/{id}", s.GetCommunityByID)

	s.Handle(http.MethodPut, "/v1/user", s.CreateUser)
	s.Handle(http.MethodGet, "/v1/user/token", s.Authenticate)
}
