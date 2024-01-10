// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBlog = "blog"

// Blog mapped from table <blog>
type Blog struct {
	ID      int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"` // 主键id
	Title   string `gorm:"column:title;type:varchar(180);not null" json:"title"`              // 文章标题
	Author  string `gorm:"column:author;type:varchar(30);not null" json:"author"`             // 作者
	Content string `gorm:"column:content;type:mediumtext;not null" json:"content"`            // 内容
	Deleted int64  `gorm:"column:deleted;type:tinyint;not null" json:"deleted"`               // 标记删除
	Hash    string `gorm:"column:hash;type:varchar(32)" json:"hash"`                          // 哈希标识
}

// TableName Blog's table name
func (*Blog) TableName() string {
	return TableNameBlog
}