package main

import (
	"TCP-Chatroom/client/processes"
	"fmt"
)

var userId int
var userPwd string

func main() {
	var key int
	var loop = true
	for {
		fmt.Println("----------歡迎登入多人聊天室----------")
		fmt.Printf("\t\t 1 登入聊天室\n")
		fmt.Printf("\t\t 2 註冊用戶\n")
		fmt.Printf("\t\t 3 退出系統\n")
		fmt.Printf("請選擇操作(1 - 3): ")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("登入聊天室")
			fmt.Printf("請輸入用戶id: ")
			fmt.Scanf("%d\n", &userId)
			fmt.Printf("請輸入用戶密碼: ")
			fmt.Scanf("%s\n", &userPwd)

			up := &processes.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("註冊用戶")
			loop = false
		case 3:
			fmt.Println("退出系統")
			loop = false
		default:
			fmt.Println("輸入錯誤, 請重新輸入")
			fmt.Println("")
		}

		if loop == false {
			break
		}

	}
}
