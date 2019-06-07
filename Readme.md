# bili-archive

## 简介

从`2019年5月27日`起，`bili-archive-frontend`不再作为`bili-archive`的一部分存在，而是`bili-archive`后端对应的一个前端实现，你可以参考本项目以及[后端接口文档][restful-api-document]实现你自己的版本。

不习惯命令行方式的用户可以移步 [bili-archive-frontend][bili-archive-frontend-repo]，该版本的体验与之前的 `Release` 一致。

[restful-api-document]: https://biliarchive.docs.apiary.io/
[bili-archive-frontend-repo]: https://github.com/Yesterday17/bili-archive-frontend

## 如何使用

通过`--help`，可以获得如下提示信息（省略了文件路径，重新排列了各项顺序）：

```
Usage of bili-archive:
  -path string
        下载视频的根目录。 (default "./Videos/")
  -cookies string
        用户的 cookies，会更新配置文件内的值。
  -uid string
        下载收藏用户的 UID，不指定则为 cookies 对应用户。
  -mode string
        下载的模式，n为通常，b为黑名单，w为白名单，配合wh和bl使用。 (default "n")
  -wl string
        下载收藏列表的白名单，用英文分号分隔，每一项为收藏夹的FID。
  -bl string
        下载收藏列表的黑名单，用英文分号分隔，每一项为收藏夹的FID。
  -s    启动后端模式。
```

## 如何构建

1. `clone`该项目

```bash
git clone https://github.com/Yesterday17/bili-archive.git
cd ./bili-archive
```

2. 构建

```bash
set CGO_ENABLED=0

# 选择架构(amd64, 386, arm)
set GOARCH=amd64

# 选择平台(windows, linux, darwin, freebsd)
set GOOS=windows
go build -o ./build/bili_archive_windows_x64.exe

set GOOS=darwin
go build -o ./build/bili_archive_darwin_amd64

set GOOS=linux
go build -o ./build/bili_archive_linux_amd64
```
