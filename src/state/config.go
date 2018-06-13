package state

import (
	"io/ioutil"
	"encoding/json"
)


type Configuration struct {
	Id      uint16 `json:"id"`
	Type    uint16 `json:"type"`
	Addr    string `json:"addr"`
	Token   string `json:"token"`
	MgrAddr string `json:"MgrAddr"`
	DbAddr string `json:"dbAddr"`
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

