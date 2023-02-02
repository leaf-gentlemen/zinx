package utils

import (
	"github.com/leaf-gentlemen/zinx/ziface"

	"github.com/spf13/viper"
)

var conf *Configure

func Interface() *Configure {
	if conf != nil {
		return conf
	}
	panic("configure fail")
}

type Configure struct {
	TCPServer      ziface.IServer // github.com/leaf-gentlemen/zinx 全局的 server 对象
	Port           int            // 网关端口号
	Host           string         // 地址
	Name           string         // 服务名称
	MaxConn        int            // 最大连接数
	MaxPackageSize int            // 最大数据包长度
	Version        string         // github.com/leaf-gentlemen/zinx 版本
	WorkerPoolSize uint32         // 最新任务池数量
	WorkerMsgLen   uint32         // 最大任务池数量
	MsgBuffChanLen uint32         // 待缓存的写管道的长度
	HedaLen        int
}

func InitConf(path string) {
	conf = &Configure{
		Name:           "[github.com/leaf-gentlemen/zinx v0.4]",
		Version:        "v0.4",
		Port:           8081,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
		WorkerPoolSize: 10,
		WorkerMsgLen:   20,
		MsgBuffChanLen: 1024,
		HedaLen:        8,
	}
	reload(conf, path)
}

// reload
//
//	@Description: 加载配置文件
//	@param c
func reload(c *Configure, path string) {
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	c.Host = viper.GetString("Host")
	c.Name = viper.GetString("Name")
	c.Port = viper.GetInt("Port")
	c.MaxConn = viper.GetInt("MaxConn")
	c.MaxPackageSize = viper.GetInt("MaxPackageSize")
	c.Version = viper.GetString("Version")
	conf = c
}
