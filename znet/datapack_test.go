package znet

import (
	"fmt"
	"io"
	"log"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listen, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Printf("accept err:%s \n", err)
				continue
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					buf := make([]byte, dp.GetHeadLen())
					cnt, err := io.ReadFull(conn, buf)
					if err != nil {
						log.Printf("read full err:%s \n", err)
						break
					}

					if len(buf) != dp.GetHeadLen() {
						log.Printf("head len fail , cnt:%d \n", cnt)
						break
					}

					msg, err := dp.UnPack(buf)
					if err != nil {
						log.Printf("unpack fail err:%s \n", err)
						break
					}

					data := make([]byte, msg.GetMsgLen())
					cnt, err = io.ReadFull(conn, data)
					if cnt != int(msg.GetMsgLen()) {
						log.Printf("data len fail , cnt:%d \n", cnt)
						break
					}

					msg.SetData(data)

					fmt.Println("==> Recv Msg: ID=", msg.GetMsgID(), ", len=", msg.GetMsgLen(), ", data=", string(msg.GetData()))
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	//创建一个封包对象 dp
	dp := NewDataPack()

	//封装一个msg1包
	msg1 := &Message{
		ID:   0,
		Len:  5,
		Data: []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
		return
	}

	msg2 := &Message{
		ID:   1,
		Len:  7,
		Data: []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client temp msg2 err:", err)
		return
	}

	//将sendData1，和 sendData2 拼接一起，组成粘包
	sendData1 = append(sendData1, sendData2...)

	//向服务器端写数据
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}
