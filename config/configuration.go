package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	defaultConfigFile = "config.json" //默认的配置文件名称
)

//Configuration 配置信息
type Configuration struct {
	MySQL struct {
		Host     string `json:"host"`     //mysql地址
		Port     string `json:"port"`     //端口
		User     string `json:"user"`     //登录用户名
		Pass     string `json:"pass"`     //登录用户密码
		Protocol string `json:"protocol"` //"tcp"
		Schema   string `json:"schema"`   //默认的schema
		Charset  string `json:"charset"`  //"utf8"
	} `json:"mysql"`
	Schema []string `json:"schema"` //需要分析的schema
}

//ConnString 连接字符串
func (c *Configuration) ConnString() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
		c.MySQL.User, c.MySQL.Pass, c.MySQL.Host, c.MySQL.Port, c.MySQL.Schema, c.MySQL.Charset)
}

//LoadConfig 加载配置文件
func LoadConfig(file string) (Configuration, error) {
	if len(file) == 0 {
		file = defaultConfigFile
	}

	var cfg Configuration

	result, err := ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("failed to read the configuration file with %v", file)
		return cfg, err
	}

	err = json.Unmarshal(result, &cfg)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal the configuration file with %v", file)
	}
	return cfg, err
}
