package model

import (
	"encoding/json"
	"io/ioutil"
)

const (
	fpath = "conf/config.json"
)

//ClientConfiguration ：
type ClientConfiguration struct {
	APIUserName string //ApiUserName : api account
	APIKey      string //ApiKey : api key
}

//Config ：
type Config struct {
	Client ClientConfiguration
}

//LoadPerformUploadAPIUserName ：
func LoadPerformUploadAPIUserName() string {
	DefJSONParse := NewJSONStruct()
	readConfig := Config{}
	DefJSONParse.Load(fpath, &readConfig)
	return readConfig.Client.APIUserName
}

//LoadPerformUploadAPIKey ：
func LoadPerformUploadAPIKey() string {
	DefJSONParse := NewJSONStruct()
	readConfig := Config{}
	DefJSONParse.Load(fpath, &readConfig)
	return readConfig.Client.APIKey
}

//JSONStruct ：
type JSONStruct struct {
}

//Load ：
func (jst *JSONStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

//NewJSONStruct ：
func NewJSONStruct() *JSONStruct {
	return &JSONStruct{}
}
