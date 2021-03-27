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
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		up.notifyToOtherUsers(userId)
	}
}

func (this *UserProcess) notifyToOtherUsers(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType

	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId = userId
	notifyUserStatusMsg.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyToOtherUsers err =", err)
		return
	}
}

func (this *UserProcess) ServerPreocessRegister(msg *message.Message) (err error) {
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err =", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	var registerResMsg message.RegisterResMsg
	err = model.MyUserDao.Register(&registerMsg.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = err.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "註冊發生未知錯誤"
		}
	} else {
		registerResMsg.Code = 200
	}

	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("json.Marshal fail =", err)
	}

	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail =", err)
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
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
	var loginResMsg message.LoginResMsg

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

		//加入線上用戶列表
		this.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(this)

		//通知其他用戶我上線了
		this.NotifyOthersOnlineUser(this.UserId)

		for _, v := range userMgr.onlineUsers {
			loginResMsg.UserIds = append(loginResMsg.UserIds, v.UserId)
		}

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
