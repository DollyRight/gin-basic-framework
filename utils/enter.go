package utils

import "encoding/json"

type Utils struct {
	HttpUtil HttpUtil
}

func (u *Utils) EncodeJSON(params map[string]interface{}) []byte {
	// 在这里使用您喜欢的JSON编码库进行编码
	// 这里使用encoding/json标准库作为示例
	encoded, _ := json.Marshal(params)
	return encoded
}

var UtilsGroup = new(Utils)
