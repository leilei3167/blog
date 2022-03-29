练手项目
https://www.topgoer.cn/docs/jianyugo/jianyugo-1cl3rba8837b5
语言规范:
https://github.com/unknwon/go-code-convention/blob/main/zh-CN/README.md

本项目用到的包:
ini:配置文件的读取
go get github.com/go-ini/ini

gorm:数据库操作
go get -u gorm.io/gorm
go get -u github.com/go-sql-driver/mysql

gin:web框架
go get -u github.com/gin-gonic/gin

学习项目开发规范,思维
1.项目目录规范

templates (views)          # 模板文件
public (static)            # 静态文件 
  css
  fonts
  img
  js
routes                     # 路由逻辑处理
models					 # 数据逻辑层
pkg                        # 子模块 
  setting                # 应用配置存取
cmd                        # 命令行程序命令
conf                       # 默认配置 
  locale                 # i18n 本地化文件
custom                     # 自定义配置
data                       # 应用生成数据文件
log                        # 应用生成日志文件
