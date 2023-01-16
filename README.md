## 一、介绍
> 该项目主要是使用gin框架，实现了将一个目录快速变为一个文件服务器，使用浏览器快速进行文件下载

> 在python3中有`http.server`模块可以实现将一个目录里的文件变为文件服务器，使用浏览器下载

## 二、使用
1、安装
```bash
go install github.com/Lyzin/ftrans
```
2、使用
> 进入需要变为文件服务器的目录，执行如下命令
```bash
ftrans serve
```
3、浏览器打开http://your-ip:8000，即可看到文件列表
