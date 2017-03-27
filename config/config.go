package config

import (
	"io/ioutil"
	"encoding/json"
)

var FilterWords map[string]int8
var MongoDB struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func init() {
	if len(FilterWords) == 0 {
		FilterWordsStr, err := ioutil.ReadFile("./config/filterWords.json")
		if err != nil {
			panic("读取filterWords失败，原因：" + err.Error())
		}
		FilterWords = make(map[string]int8)
		for _, v := range string(FilterWordsStr) {
			FilterWords[string(v)] = 1
		}
	}
	if MongoDB.Port == "" {
		MongoDBStr, err := ioutil.ReadFile("./config/mongodb.json")
		if err != nil {
			panic("读取mongodb配置文件失败，原因：" + err.Error())
		}
		if err := json.Unmarshal([]byte(MongoDBStr), &MongoDB); err != nil {
			panic("解析mongodb配置文件失败，原因：" + err.Error())
		}
	}
}
