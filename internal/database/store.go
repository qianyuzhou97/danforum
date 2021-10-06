package database

import (
	"context"
)

type Store interface {
	ListAllCommunity(ctx context.Context) ([]Community, error)
	GetCommunityByID(ctx context.Context, communityID int64) (*Community, error)
	CreateCommunity(ctx context.Context, nc NewCommunity) error

	ListAllPosts(ctx context.Context) ([]Post, error)
	GetPostByID(ctx context.Context, postID int64) (*Post, error)
	CreatePost(ctx context.Context, np NewPost) error
	UpdatePostByID(ctx context.Context, up UpdatePost) error
	DeletePostByID(ctx context.Context, postID string) error

	CreateUser(ctx context.Context, n NewUser) error
	Authenticate(ctx context.Context, username, password string) error
}
