package processes

import (
	"TCP-Chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

func outputOnlineUser() {
	fmt.Println("\n當前在線用戶列表:")
	for id, _ := range onlineUsers {
		fmt.Println(id)
	}
}

func updateUserStatus(msg *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[msg.UserId]
	if !ok {
		user = &message.User{
			UserId: msg.UserId,
		}
	}
	user.UserStatus = msg.Status
	onlineUsers[msg.UserId] = user
	outputOnlineUser()
}
