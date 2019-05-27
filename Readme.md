# bili-archive

## 如何使用

运行可执行文件，打开浏览器`localhost:8080`。  
（请勿关闭 `Console`）

## 如何构建

1. `clone`该项目

```bash
git clone https://github.com/Yesterday17/bili-archive.git
cd ./bili-archive
```

2. 获得 `statik`

```bash
go get github.com/rakyll/statik
```

3. 编译前端

```bash
cd ./bili-archive-frontend
yarn install
yarn build
```

4. 打包前端

```bash
cd bili-archive-frontend
statik -src=./build -f
```

5. 构建

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
