package processes

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerPreocessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	var loginResMsg message.LoginRes

	//func test
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		loginResMsg.Code = message.SUCCESS
	} else {
		loginResMsg.Code = message.UNREGISTERED
		loginResMsg.Error = "用戶不存在"
	}

	//對回傳資料序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal fail =", err)
	}
	resMsg.Data = string(data)

	//對回傳物件resMsg序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail =", err)
	}

	//...發送data (封裝起來)
	//err = utils.WritePkg(conn, data)
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
