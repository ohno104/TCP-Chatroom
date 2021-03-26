package utils

import (
	"TCP-Chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func ReadPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err =", err)
		return
	}
	fmt.Println("讀到的 buf=", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("read pkg header err", err)
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarsha err =", err)
		return
	}

	return
}

func WritePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return

}
