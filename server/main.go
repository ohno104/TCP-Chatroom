package main

import (
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}

	err := processor.processDistributor()
	if err != nil {
		fmt.Println("客戶端與服務器通訊協程錯誤= ", err)
	}
}

func init() {
	//初始化redis的連接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化redis操作物件
	initUserDao()
}

func main() {
	//初始化redis的連接池
	//initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化redis操作物件
	//initUserDao()

	fmt.Println("[新的結構]服務器在8889監聽...")
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
