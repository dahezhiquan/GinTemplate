package repo

import (
	"context"
	"m-sec/pkg/connections/database"
	"m-sec/pkg/model"
)

type BlogRepo interface {
	FindBlogByHash(tx database.DbConn, ctx context.Context, hash string) (*model.Blog, error)
}
