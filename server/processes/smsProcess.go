package processes

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(msg *message.Message) {

	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		this.sendMsgToEachOnlineUsers(data, up.Conn)
	}
}

func (this *SmsProcess) sendMsgToEachOnlineUsers(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("轉發消息失敗")
	}
}
