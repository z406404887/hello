package login

import (
	"io/ioutil"
	"encoding/json"
)

type DbInfo struct {
	User string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Address string `json:"address"`
	DbName string `json:"dbName"`
	Driver string `json:"driver"`
}
	

type Configuration struct {
	Id      uint16 `json:"id"`
	Type    uint16 `json:"type"`
	Addr    string `json:"addr"`
	Db DbInfo `json:"db"`
}

func NewConfiguration(path string) (*Configuration,error)  {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil,err
	}

	cfg := &Configuration{}
	err = json.Unmarshal(data,cfg)

	if err != nil {
		return  nil, err
	}
	return  cfg ,nil
}

