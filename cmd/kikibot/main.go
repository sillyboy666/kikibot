package main

import (
	"fmt"

	"kikibot"
	"kikibot/util"
)

func main() {
	fmt.Print("[KiKiBot] 请输入直播间 ID: ")
	var rid int64
	_, err := fmt.Scanln(&rid)
	if err != nil {
		util.Error("直播间 ID 不正确", nil)
		return
	}

	ch := make(chan kikibot.FmMessage)

	go kikibot.Connect(ch, rid)

	for msg := range ch {
		kikibot.HandleFmMessage(msg)
	}
}
