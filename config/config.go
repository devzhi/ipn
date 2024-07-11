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
		_, err = toml.Decode(string(tomlData), &conf)
		if err != nil {
			log.Fatalln("配置文件错误：", err)
		}
	})
	return &conf
}
