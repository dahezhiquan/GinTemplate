package domain

import (
	"context"
	"m-sec/common/errs"
	"m-sec/pkg/connections/cache"
	"m-sec/pkg/connections/cache/redis"
	"m-sec/pkg/connections/database"
	"m-sec/pkg/dao"
	"m-sec/pkg/model"
	"m-sec/pkg/repo"
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
