package test

import (
	"log"
	"m-sec/pkg/domain"
	"testing"
)

// 测试查看博客内容
func TestGetBlogDetail(t *testing.T) {
	d := domain.NewBlogDomain()
	de, _ := d.BlogDetail(nil, "1234567890")
	log.Println("======测试查看博客内容======")
	log.Println(de)
}
