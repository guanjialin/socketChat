package main

//import (
//	"db"
//	"fmt"
//	"log"
//	"net"
//	"protocol"
//	"time"
//	"tool"
//)

//func main() {
//	user := LoginUi()

//	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3000")
//	if err != nil {
//		log.Fatal("Resolve\n")
//	}

//	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
//	if err != nil {
//		log.Fatal("Dial\n", err)
//	}
//	defer tcpConn.Close()

//	if !Login(user, tcpConn) {
//		fmt.Println("账号或密码错误，5秒后自动退出")
//		time.Sleep(5 * time.Second)
//	}

//	go func() {
//		for {
//			var message protocol.Message
//			message.Unpack(tcpConn)
//			fmt.Printf("对方%s", message.From, ":", message.Msg)
//		}
//	}()

//	for {
//		var cmd string
//		var to string
//		fmt.Print("请选择发送类型消息类型",
//			protocol.Text, "为文本消息",
//			protocol.File, "为文件消息:")
//		fmt.Scanln(&cmd)
//		fmt.Print("请输入接收方:")
//		fmt.Scanln(&to)

//		var message protocol.Message
//		switch cmd[0] {
//		case protocol.Text:
//			var msg string
//			fmt.Print("请输入要发送的文本消息：")
//			fmt.Scanln(&msg)
//			message, _ = protocol.NewMessage(cmd[0],
//				tool.StringToByte16(user.Uid),
//				tool.StringToByte16(to),
//				msg)
//		case protocol.File:
//		default:
//			fmt.Println("未知类型")
//		}

//		_, err := tcpConn.Write(message.Pack())
//		if err != nil {
//			log.Print("发送失败", err)
//		}
//	}
//}

////登陆
//func Login(user db.User, tcpConn net.Conn) bool {
//	message, _ := protocol.NewMessage(protocol.Login,
//		tool.StringToByte16(user.Uid),
//		tool.StringToByte16(protocol.SERVER_ID),
//		user.Password)

//	_, err := tcpConn.Write(message.Pack())
//	if err != nil {
//		log.Print("Write", err)
//	}

//	var loginMsg protocol.Message
//	loginMsg.Unpack(tcpConn)
//	log.Print(loginMsg)

//	if !(loginMsg.Cmd == protocol.SUCCESS) {
//		return false
//	}
//	return true
//}
