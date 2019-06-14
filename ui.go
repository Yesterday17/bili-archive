package main

import (
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
	"github.com/marcusolsson/tui-go"
	"log"
	"strconv"
)

type DownloadStatus struct {
	ui     *tui.UI
	pg     *tui.Progress
	st     *tui.StatusBar
	active bool

	now    int64
	total  int64
	page   int
	status string
}

type UI struct {
	// 所有收藏夹的数组
	favList *[]bilibili.FavoriteListItemDetail
	// FID -> 收藏信息的 Map
	favMap *map[string]*bilibili.FavoriteListItemDetail
	// FID -> 收藏夹中视频数组的 Map
	favVideoList *map[string]*[]bilibili.FavoriteListItemVideo
	// AID -> 视频信息的 Map
	videoMap *map[string]*bilibili.FavoriteListItemVideo
	// AID -> 视频下载状态的 Map
	videoStatusMap *map[string]*DownloadStatus

	// 当前选中的收藏夹的信息
	activeFav *bilibili.FavoriteListItemDetail
	// 当前选中收藏夹中视频的数组
	activeFavVideoList *[]bilibili.FavoriteListItemVideo
	// 当前选中的视频
	activeVideo *bilibili.FavoriteListItemVideo

	level int
	ui    *tui.UI
	pg    *tui.Progress
	st    *tui.StatusBar
}

func (u *UI) New() {
	// 初始化变量
	u.level = 0

	favList := tui.NewList()
	favList.SetFocused(true)
	for _, list := range *u.favList {
		favList.AddItems(list.Name)
	}
	favListBox := tui.NewVBox(favList)

	favVideoList := tui.NewList()
	favVideoListBox := tui.NewVBox(favVideoList)

	selectBox := tui.NewHBox(favListBox, favVideoListBox)

	progress := tui.NewProgress(100)
	status := tui.NewStatusBar("")
	status.SetPermanentText("nep?!")

	root := tui.NewVBox(
		selectBox,
		tui.NewSpacer(),
		progress,
		status,
	)

	localUI, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}
	u.ui = &localUI
	u.pg = progress
	u.st = status

	// 在收藏列表中上下移动
	favList.OnSelectionChanged(func(list *tui.List) {
		// 边界情况
		if list.Selected() == -1 {
			return
		}

		// 更新当前选中信息
		fa := *u.favList
		u.activeFav = (*u.favMap)[strconv.Itoa(fa[list.Selected()].FID)]
		u.activeFavVideoList = (*u.favVideoList)[strconv.Itoa(fa[list.Selected()].FID)]
	})

	// 在收藏列表中回车
	favList.OnItemActivated(func(list *tui.List) {
		go (*u.ui).Update(func() {
			for _, item := range *u.activeFavVideoList {
				favVideoList.AddItems(item.Title)
			}

			// 聚焦到视频详细内容的 List
			favList.SetFocused(false)
			// 原 List 取消聚焦
			favVideoList.SetFocused(true)
		})
		u.level++
	})

	// 在视频列表中上下移动
	favVideoList.OnSelectionChanged(func(list *tui.List) {
		// 当 ESC 时会调用 此时选中值为 -1
		if list.Selected() == -1 {
			sm := *u.videoStatusMap
			nowStatus := sm[strconv.Itoa(u.activeVideo.AID)]

			nowStatus.active = false
			u.activeVideo = nil

			// 删除进度
			go (*u.ui).Update(func() {
				u.pg.SetCurrent(0)
				u.pg.SetMax(1)
				u.st.SetText("")
			})
			return
		}

		var lastActiveVideo *bilibili.FavoriteListItemVideo
		// 更新信息
		li := *u.activeFavVideoList
		nowActiveVideo := &li[list.Selected()]
		if u.activeVideo != nil {
			// 记录上一个选中的视频
			lastActiveVideo = u.activeVideo
		}

		// 当且仅当选中视频变化时更新（到列表顶部/底部）
		if nowActiveVideo != lastActiveVideo {
			// 更新当前激活的视频
			u.activeVideo = nowActiveVideo

			// 更新状态栏
			var lastStatus, nowStatus *DownloadStatus
			sm := *u.videoStatusMap
			nowStatus = sm[strconv.Itoa(u.activeVideo.AID)]
			if lastActiveVideo != nil {
				lastStatus = sm[strconv.Itoa(lastActiveVideo.AID)]
			}

			// 确保现态存在 ui 控件
			nowStatus.ui = u.ui
			nowStatus.pg = u.pg
			nowStatus.st = u.st

			if lastStatus != nil {
				lastStatus.active = false
			}
			nowStatus.active = true

			// 给当前状态一个初始值
			go (*u.ui).Update(func() {
				nowStatus.pg.SetCurrent(int(nowStatus.now))
				nowStatus.pg.SetMax(int(nowStatus.total))
			})

			switch nowStatus.status {
			case "正在下载":
				nowStatus.st.SetText(fmt.Sprintf(
					"正在下载：[av%d][P%d]%s",
					nowActiveVideo.AID,
					nowStatus.page,
					nowActiveVideo.Title,
				))
			default:
				nowStatus.st.SetText(fmt.Sprintf(
					"%s：[av%d]%s",
					nowStatus.status,
					nowActiveVideo.AID,
					nowActiveVideo.Title,
				))
			}
		}
	})

	// 在视频列表中回车
	favVideoList.OnItemActivated(func(list *tui.List) {
		//
	})

	(*u.ui).SetKeybinding("Esc", func() {
		switch u.level {
		// 收藏选择层
		case 0:
			go (*u.ui).Update(func() {
				status.SetPermanentText("退出请按Ctrl+C！")
			})

		case 1:
			// 视频层
			go (*u.ui).Update(func() {
				// 此时会调用选中值为 -1 的 OnSelectedChange 事件
				favVideoList.RemoveItems()
			})
			favVideoList.SetFocused(false)
			favList.SetFocused(true)
			u.level--
		}
	})

	(*u.ui).SetKeybinding("Ctrl+C", func() {
		(*u.ui).Quit()
	})

}

func (u *UI) Run() {
	if err := (*u.ui).Run(); err != nil {
		log.Fatal(err)
	}
}
