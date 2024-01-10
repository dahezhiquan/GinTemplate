package dto

// 博客详情页返回包

type BlogDetailResp struct {
	Blog BlogDisplay `json:"blog"`
}

type BlogDisplay struct {
	ID      int64  `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
