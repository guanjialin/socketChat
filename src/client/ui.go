package main

import (
	"db"
	"fmt"
)

//用户界面
func Ui() {

}

func LoginUi() db.User {
	var user db.User

	fmt.Print("账号：")
	fmt.Scanln(&user.Uid)
	fmt.Print("密码：")
	fmt.Scanln(&user.Password)
	return user
}
