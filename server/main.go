package main

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// func readPkg(conn net.Conn) (msg message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		//fmt.Println("conn.Read err =", err)
// 		return
// 	}
// 	fmt.Println("讀到的 buf=", buf[:4])

// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])

// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		//fmt.Println("read pkg header err", err)
// 		return
// 	}

// 	err = json.Unmarshal(buf[:pkgLen], &msg)
// 	if err != nil {
// 		fmt.Println("json.Unmarsha err =", err)
// 		return
// 	}

// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}

// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}
// 	return

// }

func serverPreocessLogin(conn net.Conn, msg *message.Message) (err error) {
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
	err = utils.WritePkg(conn, data)
	return
}

//根據客戶端發送消息的種類決定調用哪個函數來處理
func serverPreocessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		//登入邏輯處理
		err = serverPreocessLogin(conn, msg)
	//case message.RegisterMsgType:
	//註冊處理
	default:
		fmt.Println("消息類型錯誤")
	}
	return
}

func process(conn net.Conn) {
	//記得關閉
	defer conn.Close()

	for {
		fmt.Println("讀取客戶端發送的數據...")
		//將讀取工作封裝
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出, 服務器端也退出...")
				return
			} else {
				fmt.Println("readPkg err =", err)
				return
			}
		}
		fmt.Println("mgs", msg)
		err = serverPreocessMsg(conn, &msg)
		if err != nil {
			return
		}
	}
}

func main() {
	fmt.Println("服務器在8889監聽...")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	for {
		fmt.Println("等待客戶端連接服務器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
		}

		go process(conn)
	}

}
