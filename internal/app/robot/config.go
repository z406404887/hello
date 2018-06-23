package robot

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Configuration struct {
	SrvAddr    string       `json:"srvAddr"`
	Num int `json:"num"`
	SleepInterval time.Duration `json:"interval"`
	AccFormat string `json:"accFormat"`
	Password string `json:"password"`
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
