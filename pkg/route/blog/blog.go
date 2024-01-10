package blog

import (
	"github.com/gin-gonic/gin"
	"log"
	"m-sec/common"
	"m-sec/pkg/service"
)

func init() {
	log.Println("init blog route")
	common.Register(&RouterBlog{})
}

type RouterBlog struct {
}

func (*RouterBlog) Route(r *gin.Engine) {
	h := service.NewHandlerBlog()
	g := r.Group("")
	g.GET("/blog/index", h.BlogList)
	g.GET("/blog/detail/:blog_hash", h.BlogDetail)
}
