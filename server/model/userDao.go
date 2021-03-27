package model

import (
	"encoding/json"
	"fmt"

	"TCP-Chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工廠模式, 創建一個UserDao實例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		//表示已經被註冊 所以才查的到
		err = ERROR_USER_EXISTS
		return
	}

	//表示未被註冊
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存註冊用戶錯誤 err =", err)
	}

	return
}
