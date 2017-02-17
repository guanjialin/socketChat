package protocol

import (
	"log"
	"testing"
	"tool"
)

func TestPack(t *testing.T) {
	message, _ := NewMessage(Text,
		tool.StringToByte16("1111111111111111"),
		tool.StringToByte16("2222222222222222"), "你好啊")
	log.Print(message.Pack())
}

//[2 9 0 49 49 49 49 49 49 49 49 49 49 49 49 49 49 49 49 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 228 189 160 229 165 189 229 149 138]
