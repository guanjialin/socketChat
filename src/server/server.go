package main

//import (
//	"db"
//	"log"
//	"net"
//	"protocol"
//	"tool"
//)

//func tcpConn() {

//}

//func main() {
//	tcpAdd, err := net.ResolveTCPAddr("tcp", ":3000")
//	if err != nil {
//		log.Fatal("Resolve\n")
//	}

//	tcpListener, err := net.ListenTCP("tcp", tcpAdd)
//	if err != nil {
//		log.Fatal("Listen\n")
//	}

//	defer tcpListener.Close()

//	for {
//		tcpConn, err := tcpListener.AcceptTCP()
//		if err != nil {
//			log.Fatal("Accept\n")
//		}
//		go Mux(tcpConn)
//	}

//}

////消息选择器
//func Mux(conn net.Conn) {
//	for {
//		var message protocol.Message
//		message.Unpack(conn)
//		if message.Cmd != protocol.CMD_ERROR {
//			log.Print(message)

//			switch message.Cmd {
//			case protocol.Login:
//				LoginMessage(message, conn)
//			case protocol.Text:
//				TextMessage(message)
//			}
//		}
//	}
//}

////登录处理
//func LoginMessage(message protocol.Message, conn net.Conn) {
//	user := db.NewUser(tool.Byte16ToString(message.From), message.Msg)
//	if user.Authenticate() {
//		//用户验证通过
//		success, err := protocol.NewMessage(protocol.SUCCESS,
//			tool.StringToByte16(protocol.SERVER_ID), message.From, "")
//		if err != nil {
//			log.Println("newmessage\n")
//			return
//		}
//		log.Print(success)
//		conn.Write(success.Pack())
//		db.UserConn[user.Uid] = conn
//	} else {
//		//验证失败
//	}
//}

////文本信息处理
//func TextMessage(message protocol.Message) {
//	//	log.Print(message, "\n")
//	toUser := db.NewUser(tool.Byte16ToString(message.To), "")
//	if toUser.IsLogin() { //接收方已经登录，转发消息到接收方
//		toConn := db.UserConn[toUser.Uid]
//		toConn.Write(message.Pack())
//	}
//}

//////错误信息处理
////func ErrorMessage() {
////	log.Print("error msg")
////}
