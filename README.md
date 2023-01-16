## 一、介绍
> 在python3中有`http.server`模块可以实现将一个目录里的文件变为文件服务器，使用浏览器下载

> 该项目使用go的gin框架，实现了将一个目录快速变为一个文件服务器，使用浏览器快速进行文件下载,
> 并且性能比较高

## 二、使用
### 1、安装
```bash
go install github.com/Lyzin/ftrans
```
#### 1.1 查看本工具帮助信息

```bash
ftrans -h
```

### 2、使用
```bash
# 进入需要变为文件服务器的目录，执行如下命令
cd your-file-dir

# 启动服务，会默认将当前目录作为文件服务器的目录
ftrans serve
```
### 3、查看效果
浏览器打开[http://your-ip:8000](http://your-ip:8000)，即可看到文件列表
