package service

import (
	"github.com/gin-gonic/gin"
	"m-sec/common"
	"m-sec/common/copier"
	"m-sec/pkg/connections/database"
	"m-sec/pkg/connections/database/gorms"
	"m-sec/pkg/connections/database/transaction"
	"m-sec/pkg/domain"
	"m-sec/pkg/dto"
	"net/http"
)

type HandlerBlog struct {
	blogDomain *domain.BlogDomain
	tx         *transaction.Transaction
	conn       database.DbConn
}

func NewHandlerBlog() *HandlerBlog {
	return &HandlerBlog{
		blogDomain: domain.NewBlogDomain(),
		tx:         transaction.NewTransaction(),
		conn:       gorms.NewConn(),
	}
}

func (b *HandlerBlog) BlogDetail(ctx *gin.Context) {
	var result = &common.Result{}
	var resp = dto.BlogDetailResp{}

	// 获取blog_hash 并进行参数校验
	hash := ctx.Param("blog_hash")

	// 通过用户hash获取作者信息
	blog, err := b.blogDomain.BlogDetail(nil, hash)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	// 模型转换与组装
	var blogDisplay dto.BlogDisplay
	if err := copier.Copy(&blogDisplay, blog); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	resp.Blog = blogDisplay
	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
