package main

import (
	"fmt"
	"github.com/leaf-gentlemen/zinx/utils"
	"github.com/leaf-gentlemen/zinx/znet"
	"io"
	"net"
	"time"

	"go.uber.org/zap"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:8082")
	if err != nil {
		fmt.Printf("dial connect error: %s \n", err)
		return
	}
	logger := utils.Logger
	dp := znet.NewDataPack()
	for {
		msg := znet.NewMessage(0, []byte("hello server"))
		buf, err := dp.Pack(msg)
		if err != nil {
			logger.Error("msg pack fail", zap.Error(err))
			return
		}

		if _, err = conn.Write(buf); err != nil {
			logger.Error("msg write fail", zap.Error(err))
			return
		}

		buf = make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			fmt.Printf("read head buf error :%s", err)
			return
		}

		head, err := dp.UnPack(buf)
		if err != nil {
			logger.Error("msg unpack fail", zap.Error(err))
			return
		}

		msgData := make([]byte, head.GetMsgLen())
		cnt, err := io.ReadFull(conn, msgData)
		if err != nil {
			logger.Error("msg data reader fail", zap.Error(err))
			return
		}

		head.SetData(msgData)
		fmt.Printf("read msgID  :%d, cnt: %d  data:%s \n", head.GetMsgID(), cnt, head.GetData())
		time.Sleep(time.Second)
	}
}
