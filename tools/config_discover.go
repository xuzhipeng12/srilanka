package tools

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Cfg *ConfigMap

type ConfigMap struct {
	ListenPort string `yaml:"listen_port"`
	LogPath    string `yaml:"log_path"`

	DBType       string `yaml:"db_type"`
	DBUser       string `yaml:"db_user"`
	DBPwd        string `yaml:"db_pwd"`
	DBHost       string `yaml:"db_host"`
	DBPort       int    `yaml:"db_port"`
	DBName       string `yaml:"db_name"`
	SQLiteDBName string `yaml:"sqlite_db_name"`
	TokenSecret  string `yaml:"token_secret"`
}

func ConfigRead(configFile string) {
	config := new(ConfigMap)
	config.readConfigFromYaml(configFile)
	fmt.Println("Final config: ", config)

	if config.ListenPort == "" {
		panic("init config fail")
	}
	Cfg = config
}

func (this *ConfigMap) readConfigFromYaml(configFile string) {
	if err := configor.Load(this, configFile); err != nil {
		fmt.Println("read config yaml err: " + err.Error())
	}
	fmt.Println("Config from yaml file: ", this)
}
