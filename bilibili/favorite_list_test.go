package bilibili

import "testing"

func Test_GetFavoriteList(t *testing.T) {
	// 使用似乎确定死亡的账号
	// UID = 4
	// 收藏夹基本没有变动
	list, err := GetFavoriteList("4", "")
	if err != nil {
		t.Error("[×][收藏列表] " + err.Error())
	}

	if len(list) != 3 {
		t.Error("[×][收藏列表] 收藏夹数量不符！")
	}

	if list[0].MediaID != 62664104 || list[0].FID != 626641 || list[0].Ctime != 1438922096 {
		t.Error("[×][收藏列表] 默认收藏夹基本信息不符！")
	}

	if list[1].Name != "视频" {
		t.Error("[×][收藏列表] 收藏夹1名称不符！应为'视频'，实为'" + list[1].Name + "'。")
	}

	if list[2].Name != "MAD" {
		t.Error("[×][收藏列表] 收藏夹1名称不符！应为'MAD'，实为'" + list[1].Name + "'。")
	}
	t.Log("[√][收藏列表] 收藏夹信息检验成功！")
}
