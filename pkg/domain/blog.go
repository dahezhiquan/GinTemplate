package domain

import (
	"GinTemplate/common/errs"
	"GinTemplate/pkg/connections/cache"
	"GinTemplate/pkg/connections/cache/redis"
	"GinTemplate/pkg/connections/database"
	"GinTemplate/pkg/dao"
	"GinTemplate/pkg/model"
	"GinTemplate/pkg/repo"
	"context"
	"time"
)

type BlogDomain struct {
	blogRepo repo.BlogRepo
	cache    cache.Cache
}

func NewBlogDomain() *BlogDomain {
	return &BlogDomain{
		blogRepo: dao.NewBlogDao(),
		cache:    redis.Rc,
	}
}

func (d *BlogDomain) BlogDetail(conn database.DbConn, hash string) (*model.Blog, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	blog, err := d.blogRepo.FindBlogByHash(conn, ctx, hash)
	if err != nil {
		return nil, errs.DBError
	}
	if blog == nil {
		return nil, errs.BlogNotExistError
	}
	return blog, nil
}
