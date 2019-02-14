package main

import (
	"./bilibili"
	"fmt"
)

func main() {
	fmt.Println(bilibili.GetUserFavoriteListReport("123817257", false))
}
