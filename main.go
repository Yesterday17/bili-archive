package main

import (
	"fmt"
	"github.com/Yesterday17/bili-archive/bilibili"
)

func main() {
	fmt.Println(bilibili.GetUserFavoriteListReport("123817257", false))
}
