package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//記得關閉
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		fmt.Println("讀取客戶端發送的數據...")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err =", err)
			return
		}
		fmt.Println("讀到的 buf=", buf[:4])
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
