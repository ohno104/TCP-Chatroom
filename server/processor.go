package main

import (
	"TCP-Chatroom/common/message"
	"TCP-Chatroom/server/processes"
	"TCP-Chatroom/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverPreocessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		//登入邏輯處理
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerPreocessLogin(msg)

	case message.RegisterMsgType:
		//註冊處理
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerPreocessRegister(msg)

	default:
		fmt.Println("消息類型錯誤")
	}
	return
}

func (this *Processor) processDistributor() (err error) {
	for {
		fmt.Println("讀取客戶端發送的數據...")
		//將讀取工作封裝

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出, 服務器端也退出...")
				return err
			} else {
				fmt.Println("readPkg err =", err)
				return err
			}
		}
		fmt.Println("mgs", msg)
		err = this.serverPreocessMsg(&msg)
		if err != nil {
			return err
		}
	}
}
