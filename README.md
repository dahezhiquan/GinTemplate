# GinTemplate❤️
Go-Gin后端模板项目，帮助你快速构建一个Gin后端服务

# 配置文件说明⚙

### 服务信息

在这个部分，列出该项目的服务信息配置和参数说明。

- **Name:** 服务名称
  - 介绍：配置服务的名称
  - 示例：`demo-hahaha`

- **Addr:** 服务地址
  - 介绍：配置服务的地址和端口
  - 示例：`127.0.0.1:8989`

### MySQL信息

在这个部分，列出项目中 MySQL 的配置和参数说明。

- **Username:** 用户名
  - 介绍：MySQL 数据库用户名
  - 示例：`root`

- **Password:** 密码
  - 介绍：MySQL 数据库密码
  - 示例：`password`

- **Host:** 主机地址
  - 介绍：MySQL 主机地址
  - 示例：`192.168.0.1`

- **Port:** 端口
  - 介绍：MySQL 端口
  - 示例：`3306`

- **DB:** 数据库名称
  - 介绍：MySQL 数据库名称
  - 示例：`demo`

- **MaxOpenConns:** 最大打开连接数
  - 介绍：MySQL 最大打开连接数
  - 示例：`100`

- **MaxIdleConns:** 最大空闲连接数
  - 介绍：MySQL 最大空闲连接数
  - 示例：`20`

### Redis信息

在这个部分，列出项目中 Redis 的配置和参数说明。

- **Host:** 主机地址
  - 介绍：Redis 主机地址
  - 示例：`127.0.0.1`

- **Port:** 端口
  - 介绍：Redis 端口
  - 示例：`6379`

- **Password:** 密码
  - 介绍：Redis 密码
  - 示例：`""`（空）

- **DB:** 数据库索引
  - 介绍：Redis 数据库索引
  - 示例：`0`

### JWT凭据

在这个部分，列出项目中 JWT 的配置和参数说明。

- **AccessExp:** 访问令牌过期时间（秒）
  - 介绍：访问令牌的过期时间，以秒为单位
  - 示例：`70000`

- **RefreshExp:** 刷新令牌过期时间（秒）
  - 介绍：刷新令牌的过期时间，以秒为单位
  - 示例：`140000`

- **AccessSecret:** 访问令牌密钥
  - 介绍：访问令牌的密钥
  - 示例：`xxx`

- **RefreshSecret:** 刷新令牌密钥
  - 介绍：刷新令牌的密钥
  - 示例：`xxx`

### Zap日志配置

在这个部分，列出项目中 Zap 日志的配置和参数说明。

- **DebugFileName:** 调试日志文件保存位置
  - 介绍：配置调试日志文件的保存位置
  - 示例：`""`（空字符串表示未设置）

- **InfoFileName:** 信息日志文件保存位置
  - 介绍：配置信息日志文件的保存位置
  - 示例：`""`（空字符串表示未设置）

- **WarnFileName:** 警告日志文件保存位置
  - 介绍：配置警告日志文件的保存位置
  - 示例：`""`（空字符串表示未设置）

- **MaxSize:** 最大日志文件大小（MB)
  - 介绍：配置日志文件的最大大小，以 MB 为单位
  - 示例：`500`

- **MaxAge:** 日志文件保存期限（天)
  - 介绍：配置日志文件的最大保存期限，以天为单位
  - 示例：`28`

- **MaxBackups:** 最大日志文件备份个数
  - 介绍：配置日志文件的最大备份数量
  - 示例：`3`


# 启动方法🚀

配置文件一切准备就绪后，在MySQL中导入项目下的初始SQL文件，之后运行main.go即可


运行之后，你会收到如下的提示：

![image](https://github.com/dahezhiquan/GinTemplate/assets/76278560/15e0883c-716b-4c32-bb1a-6adc2735dae6)

访问接口地址：http://127.0.0.1:8989/blog/detail/1234567890

如果一切正常，你会得到服务返回的json数据：

![image](https://github.com/dahezhiquan/GinTemplate/assets/76278560/69ab834b-b7fa-43a7-a89f-3fbd884e0a73)

同时在运行端会输出对应的日志信息：

![image](https://github.com/dahezhiquan/GinTemplate/assets/76278560/ba0e84d9-53b0-40ac-b229-db067cf289bf)









