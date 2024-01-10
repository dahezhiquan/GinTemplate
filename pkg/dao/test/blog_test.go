package test

import (
	"GinTemplate/pkg/dao"
	"context"
	"log"
	"testing"
	"time"
)

// 测试通过hash获取博客详情
func TestGetBlogContentByHash(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d := dao.NewBlogDao()
	blog, _ := d.FindBlogByHash(nil, ctx, "1234567890")
	log.Println("======测试通过hash获取博客详情======")
	log.Println(blog.Content)
}
