//网络协议
//  cmd(uint8) | length(uint16) | from([16]byte)    | to([16]byte) | msg(string)
//	cmd最大256  | length最大65536 |
package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"tool"
)

const (
	CMD_ERROR uint8 = iota //错误命令
	SUCCESS                //成功消息
	Login                  //登陆命令
	Text                   //文本消息
	File                   //文件消息
)

const (
	SERVER_ID = "0000000000000000"
)

const (
	MAX_MSG_SIZE int = 1 << 16
)

type Message struct {
	Cmd    uint8    //标识该消息类型
	Length uint16   //消息正文长度
	From   [16]byte //消息发送者
	To     [16]byte //消息接收者
	Msg    string   //消息正文
}

//构造消息体，如果长度超过最大长度则返回错误
func NewMessage(cmd uint8, from [16]byte, to [16]byte, msg string) (Message, error) {
	var message Message
	length := len(msg)

	//如果命令字无效，返回错误
	if invaild(cmd) {
		return message, errors.New("未知命令")
	}

	//长度太长返回错误
	if len(msg) > MAX_MSG_SIZE {
		return message, errors.New("正文长度不能超多" + string(MAX_MSG_SIZE))
	}
	message.Cmd = cmd
	message.Length = uint16(length)
	message.From = from
	message.To = to
	message.Msg = msg
	return message, nil
}

//命令字是否有效,无效返回true
func invaild(cmd uint8) bool {
	switch cmd {
	case SUCCESS, Login, Text:
		return false
	default:
		return true
	}

}

//封装协议
func (message Message) Pack() []byte {
	var body []byte
	buf := bytes.NewBuffer(body)
	buf.WriteByte(message.Cmd) //写入命令字

	buf.Write(tool.Uint16toBytes(message.Length)) //写入长度值

	for i := 0; i < 16; i++ {
		buf.WriteByte(message.From[i])
	} //写入发送人

	for i := 0; i < 16; i++ {
		buf.WriteByte(message.To[i])
	} //写入收件人

	buf.WriteString(message.Msg) //写入消息体

	body = buf.Bytes()
	return body
}

//解析协议
func (message *Message) Unpack(conn net.Conn) {
	cmd := make([]byte, 1)
	length := make([]byte, 2)
	from := make([]byte, 16)
	to := make([]byte, 16)
	conn.Read(cmd)
	conn.Read(length)
	conn.Read(from)
	conn.Read(to)

	msg := make([]byte, binary.LittleEndian.Uint16(length))
	conn.Read(msg)

	message.Cmd = cmd[0]
	message.Length = binary.LittleEndian.Uint16(length)
	message.From = tool.BytesToByte16(from)
	message.To = tool.BytesToByte16(to)
	message.Msg = string(msg)

}
