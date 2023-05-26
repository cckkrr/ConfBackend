# 程序介绍

本程序为后端软件

## 项目开发环境
- 系统 GOOS=Darwin
- 架构 GOARCH=amd64
- Go version: 1.19.1

## 任意环境下打包为单个Linux环境下二进制可运行文件命令（交叉编译）
在项目目录下：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go


## 项目结构

程序在server包中启动，每个路由调用view包中视图，视图从control中获取服务。

## 程序启动配置读取
项目目录下 ./etc/app.conf文件中配置属性，修改时在 services/app_conf.go 中对应添加属性名及类别（属性名首字母大写，否则无法导出）


## 开发过程中
- main.go中的main函数新建go协程监听端口，发送小车命令
- http返回见common包response.go，使用通用返回，使用方法见view包代码
