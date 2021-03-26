package main

import (
	"TCP-Chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err =", err)
		return
	}
	fmt.Println("讀到的 buf=", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("read pkg header err", err)
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarsha err =", err)
		return
	}

	return
}

func process(conn net.Conn) {
	//記得關閉
	defer conn.Close()

	for {
		fmt.Println("讀取客戶端發送的數據...")
		//將讀取工作封裝
		msg, err := readPkg(conn)
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
