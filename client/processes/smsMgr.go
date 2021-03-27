package processes

import (
	"TCP-Chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMsg(msg *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err.Error())
		return
	}

	info := fmt.Sprintf("=> %d: %s\n", smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
}
