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
statik -src=/path/to/your/project/public
```

5. 构建

```bash
go build
```
