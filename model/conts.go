package model

import (
	"encoding/json"
	"io/ioutil"
)

const (
	fpath = "conf/config.json"
)

type ClientConfiguration struct {
	ApiUserName string //ApiUserName : api account
	ApiKey      string //ApiKey : api key
}

type Config struct {
	Client ClientConfiguration
}

func LoadPerformUploadApiUserName() string {
	DefJsonParse := NewJsonStruct()
	readConfig := Config{}
	DefJsonParse.Load(fpath, &readConfig)
	return readConfig.Client.ApiUserName
}

func LoadPerformUploadApiKey() string {
	DefJsonParse := NewJsonStruct()
	readConfig := Config{}
	DefJsonParse.Load(fpath, &readConfig)
	return readConfig.Client.ApiKey
}

type JsonStruct struct {
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
