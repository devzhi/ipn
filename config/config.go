package config

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Device     string
	SmtpHost   string
	SmtpPort   int
	From       string
	Password   string
	To         []string
	AlwaysSend bool
	EnableCron bool
	CronExp    string
}

var conf Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		// 读取TOML
		tomlData, err := ioutil.ReadFile("config.toml")
		if err != nil {
			log.Fatalln(err)
		}
		tomlStr := string(tomlData)
		// 检测编码是否为UTF-8 with BOM
		if tomlStr[0] == 0xEF && tomlStr[1] == 0xBB && tomlStr[2] == 0xBF {
			tomlStr = tomlStr[3:]
		}
		_, err = toml.Decode(tomlStr, &conf)
		if err != nil {
			log.Fatalln("配置文件错误：", err)
		}
	})
	return &conf
}
