package bilibili

import (
	"fmt"
	"testing"
)

func Test_GetVideoPages(t *testing.T) {
	pageNames := []string{
		"#1 Java 8 的安装与卸载",
		"#2 HMCL 启动器入门简介",
		"#3 模组相关资源的获取",
		"#4 模组的安装与汉化资源包的添加",
		"#5 火狐浏览器的安装与插件的使用",
		"#6 JEI 的安装与使用",
		"#7 信息高亮显示模组安装与介绍",
		"#8 Optifine的使用与光影的安装",
		"#9 自动整理插件和鼠标手势的使用",
		"#10 旅行地图的使用",
		"#11 FTB实用工具介绍",
		"#12 平滑字体模组",
		"#13 整合包是什么",
		"#14 整合包下载的一种方法",
		"#15 服务器的选择",
		"#16 服务端的搭建与使用",
	}

	cids := []int{
		32494776,
		32494777,
		32494778,
		32494779,
		32571129,
		32583803,
		32627386,
		32653618,
		32701687,
		32704892,
		32707753,
		32711247,
		32717985,
		32718220,
		33548175,
		33550141,
	}

	pages, err := GetVideoPages("19922956")
	if err != nil {
		t.Error(err)
	}

	for _, page := range pages {
		if page.CID != cids[page.Page-1] {
			t.Error(fmt.Sprintf("[×][分P] CID 不匹配! 应为 %d, 实为 %d。", cids[page.Page-1], page.CID))
		}

		if page.PageName != pageNames[page.Page-1] {
			t.Error(fmt.Sprintf("[×][分P] 分P名称不匹配! 应为 %s, 实为 %s。", pageNames[page.Page-1], page.PageName))
		}
		t.Log(fmt.Sprintf("[√][分P] P%d: %s(%d)", page.Page, page.PageName, page.CID))
	}
}
