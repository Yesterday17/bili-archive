package bilibili

import (
	"fmt"
	"testing"
)

func Test_GetFavoriteListItems(t *testing.T) {
	items, err := GetFavoriteListItems("4", "1273667", "1", "")
	if err != nil {
		t.Error("[×][收藏内容] " + err.Error())
	}

	if len(items) != 1 {
		t.Error(fmt.Sprintf("[×][收藏内容] 收藏视频数量不符！应为 1，实为%d。", len(items)))
	}

	if items[0].AID != 20728609 {
		t.Error(fmt.Sprintf("[×][收藏内容] 收藏视频AID不符！应为 20728609，实为%d。", items[0].AID))
	}
	t.Log(fmt.Sprintf("[×][收藏内容] %s(av%d)", items[0].Title, items[0].AID))
}
