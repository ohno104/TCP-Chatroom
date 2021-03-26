package login

import "fmt"

func Login(userId int, userPwd string) (err error) {
	fmt.Printf("userId = %d userPwd=%s \n", userId, userPwd)
	return nil
}
