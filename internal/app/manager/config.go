package manager

import (
	"encoding/json"
	"io/ioutil"
)

type ServerItem struct {
	Id   uint16
	Type uint16
	Addr string
}

type Configuration struct {
	Id      uint16       `json:"id"`
	Type    uint16       `json:"type"`
	Addr    string       `json:"addr"`
	Token   string       `json:"token"`
	Servers []ServerItem `json:"servers"`
}

func NewConfiguration(path string) (*Configuration, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	cfg := &Configuration{}
	err = json.Unmarshal(data, cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
