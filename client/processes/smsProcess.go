package processes

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMsg(content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType

	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = curUser.UserId
	smsMsg.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMsg json.Marshal fail =", err.Error())
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg json.Marshal fail =", err.Error())
	}

	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg err =", err.Error())
	}
	return
}
