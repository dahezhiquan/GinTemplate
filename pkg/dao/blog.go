package dao

import (
	"context"
	"m-sec/pkg/connections/database"
	"m-sec/pkg/connections/database/gorms"
	"m-sec/pkg/model"
	"m-sec/pkg/public"
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
	err := conn.Session(ctx).Where("hash = ? and status=1 and deleted = 0", hash).First(&blog).Error
	if err == gorms.ErrRecordNotFound {
		return nil, nil
	}
	return blog, err
}

func (b *BlogDao) FindBlogList(tx database.DbConn, ctx context.Context, page int) (list []*model.Blog, total int64, err error) {
	conn := b.getConn(tx)
	db := conn.Session(ctx).Model(&model.Blog{})
	err = db.Scopes(gorms.Paginate(page, public.VUL_PAGE_SIZE)).Where("status=1 and deleted =0").Order("is_top DESC, id DESC").Find(&list).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
