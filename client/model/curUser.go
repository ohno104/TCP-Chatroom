package model

import (
	"TCP-Chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
