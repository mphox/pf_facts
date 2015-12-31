package model

import (
	"encoding/json"
	"io/ioutil"
)

type Rule struct {
	Id   int64  `json:"_id"`
	Text string `json:"text"`
}

func LoadRule(ruleId string) (rule *Rule, err error) {
	rule = &Rule{}
	var ruleBytes []byte
	filename := "./data/" + ruleId + ".json"
	ruleBytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(ruleBytes, rule)
	if err != nil {
		return
	}
	return
}
