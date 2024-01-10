package blog

import (
	"GinTemplate/common"
	"GinTemplate/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
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
	g.GET("/blog/detail/:blog_hash", h.BlogDetail)
}
