package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	port int `json:"port"`
	ReadTimeout int `json:"ReadTimeout"`
	WriteTimeout int `json:"WriteTimeout"`
	EtcdAndPoints []string `json:"etcdAndPoints"`
	EtcdDialTimeout int `json:"etcdDialTimeout"`
}

var (
	G_config *Config
)

func InitConfig(fileName string)(err error)  {
	var (
		content []byte
		conf Config
	)
	// 读取配置文件
	if content, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}
	// 反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return err
	}
	// 赋值单例
	G_config = &conf
	return nil
}