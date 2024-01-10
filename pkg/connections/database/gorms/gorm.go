package gorms

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"m-sec/config"
)

var _db *gorm.DB

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

func init() {
	//配置MySQL连接参数
	username := config.Conf.MysqlConfig.Username         //账号
	password := config.Conf.MysqlConfig.Password         //密码
	host := config.Conf.MysqlConfig.Host                 //数据库地址，可以是Ip或者域名
	port := config.Conf.MysqlConfig.Port                 //数据库端口
	Dbname := config.Conf.MysqlConfig.Db                 //数据库名
	maxOpenConns := config.Conf.MysqlConfig.MaxOpenConns //连接池最大连接数
	maxIdleConns := config.Conf.MysqlConfig.MaxIdleConns // 连接池最大空闲连接数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
}

func GetDB() *gorm.DB {
	return _db
}

type GormConn struct {
	tx *gorm.DB
}

func NewConn() *GormConn {
	return &GormConn{tx: GetDB()}
}

func NewTransaction() *GormConn {
	return &GormConn{tx: GetDB()}
}

func (g *GormConn) Session(ctx context.Context) *gorm.DB {
	return g.tx.Session(&gorm.Session{Context: ctx})
}

func (g *GormConn) Begin() {
	g.tx = GetDB().Begin()
}

func (g *GormConn) Rollback() {
	g.tx.Rollback()
}

func (g *GormConn) Commit() {
	g.tx.Commit()
}

// 分页函数的支持

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
