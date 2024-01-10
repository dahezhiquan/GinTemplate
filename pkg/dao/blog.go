package dao

import (
	"GinTemplate/pkg/connections/database"
	"GinTemplate/pkg/connections/database/gorms"
	"GinTemplate/pkg/model"
	"context"
)

type BlogDao struct {
	baseDao
}

func NewBlogDao() *BlogDao {
	return &BlogDao{baseDao{conn: gorms.NewConn()}}
}

func (b *BlogDao) FindBlogByHash(tx database.DbConn, ctx context.Context, hash string) (*model.Blog, error) {
	var blog *model.Blog
	conn := b.getConn(tx)
	err := conn.Session(ctx).Where("hash = ? and deleted = 0", hash).First(&blog).Error
	if err == gorms.ErrRecordNotFound {
		return nil, nil
	}
	return blog, err
}
