package kikibot

import (
	"fmt"

	"kikibot/util"
)

func HandleFmMessage(msg FmMessage) {
	switch msg.Type {
	case TypeMember:
		handleMember(msg)
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

// ------------------------------------------------------------------- //

func handleMemberJoinQueue(msg FmMessage) {
	for _, v := range msg.Queue {
		name := v.Username
		if name == "" {
			// 匿名用户
			util.Info("匿名用户进入直播间")
		} else {
			util.Info(fmt.Sprintf("用户 @%s 进入直播间", name))
		}
	}
}
