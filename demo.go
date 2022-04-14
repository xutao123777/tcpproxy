package main

import (
	"log"

	"github.com/xutao123777/tcppack"
)

/*
	用于解决tcp粘包的库
*/

func main() {
	// 打包一个消息(封包)
	buf := tcppack.Pack([]byte("this is message"))

	// 解包一个消息(解包)
	buf = tcppack.Unpack(buf)

	log.Println(string(buf)) // 输出结果 this is message
}
