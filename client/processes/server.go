package processes

import (
	"TCP-Chatroom/utils"
	"fmt"
	"net"
	"os"
)

//顯示登入成功介面
func ShowMenu() {
	fmt.Println("----------歡迎登入----------")
	fmt.Printf("\t\t 1 顯示在線用戶列表\n")
	fmt.Printf("\t\t 2 發送消息\n")
	fmt.Printf("\t\t 3 信息列表\n")
	fmt.Printf("\t\t 4 退出系統\n")
	fmt.Printf("請選擇操作(1 - 4): ")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("顯示在線用戶列表")
	case 2:
		fmt.Println("發送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系統")
		os.Exit(0)
	default:
		fmt.Println("輸入錯誤, 請重新輸入")
	}
}

//保持和服務器通訊
func serverProcessMsg(Conn net.Conn) {
	//持續讀取服務器的發送的消息
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err =", err)
			return
		}
		fmt.Printf("你收到了一則訊息: %v", msg)
	}
}
