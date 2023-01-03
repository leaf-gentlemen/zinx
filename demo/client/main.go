package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8081")
	if err != nil {
		fmt.Printf("dial connect error: %s \n", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Printf("write buf error:%s", err)
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read buf error :%s", err)
			return
		}

		fmt.Printf("read buf content :%s, cnt: %d \n", buf, cnt)

		time.Sleep(time.Second)
	}
}
