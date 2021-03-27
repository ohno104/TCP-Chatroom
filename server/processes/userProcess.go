package processes

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/server/model"
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

	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()

		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服務器內部錯誤"
		}

	} else {
		loginResMsg.Code = 200
		fmt.Printf("%v 登入成功\n", user.UserName)
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
