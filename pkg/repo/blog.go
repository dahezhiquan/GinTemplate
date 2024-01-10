package repo

import (
	"GinTemplate/pkg/connections/database"
	"GinTemplate/pkg/model"
	"context"
)

type BlogRepo interface {
	FindBlogByHash(tx database.DbConn, ctx context.Context, hash string) (*model.Blog, error)
}
