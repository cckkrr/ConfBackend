## 文件夹说明

该文件夹Services存放的是对于该"软件"的服务，比如读取本地配置文件（app.conf），以及数据库、缓存的初始化等。

## 程序启动配置读取
项目目录下 ./etc/app.conf文件中配置属性，修改时在 services/app_conf.go 中对应添加属性名及类别（属性名首字母大写，否则无法导出）

## Services的使用

直接输入S.S.[你想使用的服务]即可

## 服务的注册

在services/service.go中初始化你的服务。
