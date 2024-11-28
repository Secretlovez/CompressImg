# 图片压缩工具

## 简介

这是一个简单的 Go 语言编写的图片压缩工具。它可以扫描指定目录或当前目录中的所有图片文件，进行图片压缩,图片压缩默认名称为xxx_thumb.jpg/xxx_thumb.png/xxxx_thumb。

## 安装

确保你已经安装了 Go 语言环境。如果没有安装，可以从 [Go 官方网站](https://golang.org/) 下载并安装。

## 使用方法

运行程序赋权限，后输入要压缩的图片路径。未输入，则遍历当前目录下所有jpg\png格式的图片进行压缩。
```sh
chmod 777 -R ./compressImg
./compressImg  /path/to/directory
```
or
```sh
chmod 777 -R ./compressImg
./compressImg
```