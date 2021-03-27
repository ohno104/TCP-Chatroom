package processes

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dail err =", err)
		return
	}
	//記得關閉
	defer conn.Close()

	var msg message.Message
	msg.Type = message.RegisterMsgType

	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	//data內容序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	msg.Data = string(data)

	//整個package序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	tf := utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("註冊發送信息錯誤 =", err)
	}

	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) error =", err)
	}

	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)
	if registerResMsg.Code == 200 {
		fmt.Println("註冊成功")

	} else {
		fmt.Println(registerResMsg.Error)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf("userId = %d userPwd=%s \n", userId, userPwd)
	// return nil

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dail err =", err)
		return
	}
	//記得關閉
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginMsgType

	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	//data內容序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	msg.Data = string(data)

	//整個package序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//先發送data長度給服務器用來判斷是否丟包
	//需要將data轉成byte切片才能使用conn.Write()
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//fmt.Printf("客戶端, 發送消息的長度=%d\n", len(data))

	//發送消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write(data) fail =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err =", err)
		return
	}

	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		//fmt.Println("登入成功")

		//紀錄自己的user
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		fmt.Println("目前的線上用戶:")
		for _, v := range loginResMsg.UserIds {
			if v == userId {
				continue
			}
			fmt.Println(v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Printf("\n\n")

		//需要在客戶端啟動一個協程
		//該協程保持和服務器的通訊, 如果服務器有數據推送給客戶端
		//則接收並顯示在客戶端的終端
		go serverProcessMsg(conn)

		//登入後的二級菜單
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMsg.Error)
	}

	return
}
