package db

import "net"

var UserConn = make(map[string]net.Conn)    //用于保存用户的TCPConn
var UserUDPConn = make(map[string]net.Conn) //用于保存用户的UDPConn

type User struct {
	Uid      string
	Password string
}

//User构造函数
func NewUser(uid string, psw string) User {
	return User{uid, psw}
}

//登陆验证
func (u User) Authenticate() bool {
	return true
}

//判断用户是否已经登录
func (u User) IsLogin() bool {
	if UserConn[u.Uid] != nil {
		return true
	} else {
		return false
	}
}

//以下为数据操作
