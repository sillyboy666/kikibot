package kikibot

import (
	"fmt"

	"github.com/fatih/color"
	"kikibot/util"
)

func HandleFmMessage(msg FmMessage) {
	switch msg.Type {
	case TypeMember:
		handleMember(msg)
	case TypeMessage:
		handleMessage(msg)
	}
}

// ------------------------------------------------------------------- //

// handleMember 处理类型为 Member 的消息
func handleMember(msg FmMessage) {
	switch msg.Event {
	case EventJoinQueue:
		handleMemberJoinQueue(msg)
	}
}

func handleMessage(msg FmMessage) {
	switch msg.Event {
	case EventNew:
		handleMessageNew(msg)
	}
}

// ------------------------------------------------------------------- //

func handleMemberJoinQueue(msg FmMessage) {
	for _, v := range msg.Queue {
		name := v.Username
		if name == "" {
			// 匿名用户
			util.Print("匿名用户进入直播间", color.FgHiCyan)
		} else {
			util.Print(fmt.Sprintf("用户 @%s 进入直播间", name), color.FgHiCyan)
		}
	}
}

func handleMessageNew(msg FmMessage) {
	util.Print(fmt.Sprintf("@%s: %s", msg.User.Username, msg.Message), color.FgGreen)
}
