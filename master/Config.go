package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	port int `json:"port"`
	ReadTimeout int `json:"ReadTimeout"`
	WriteTimeout int `json:"WriteTimeout"`
}

func InitConfig(fileName string)(err error)  {
	var (
		content []byte
		conf Config
	)
	
	if content, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return err
	}

	return nil
}