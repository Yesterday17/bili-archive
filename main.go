package main

import "log"

func main() {
	// 加载配置文件
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}
	// 立即保存（初始设置）
	QuickSaveConfig()
	// 服务器
	CreateBiliArchiveServer()
}
