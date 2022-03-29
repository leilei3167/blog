//用于读取配置文件
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

//所需结果公开

var (
	//操作句柄
	Cfg       *ini.File
	HTTP_PORT int

	READ_TIMEOUT  time.Duration
	WRITE_TIMEOUT time.Duration
)

func init() {
	var err error
	//Cfg为赋值
	Cfg, err = ini.Load("conf/my.ini")
	if err != nil {
		log.Fatal("打开配置文件出错", err)
	}

	//读取配置

	LoadBase()
	LoadServer()

}

func LoadBase() {
	Cfg.Section("").Key("RUN_MODE").MustString("debug")

}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("server读取失败:", err)
	}
	HTTP_PORT = sec.Key("HTTP_PORT").MustInt(8080)

	READ_TIMEOUT = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt())
	WRITE_TIMEOUT = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt())

}
