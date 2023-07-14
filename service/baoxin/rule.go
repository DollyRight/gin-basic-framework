package baoxin

import (
	"encoding/json"
	"fmt"
	"gin-basic-framework/model/common/request"
	"gin-basic-framework/utils"
	"github.com/nsqio/go-nsq"
)

type RuleService struct {
}

type RuleHandler struct {
}

func (rh *RuleHandler) HandleRule(message *nsq.Message) {
	ruleTaskInfo := make(map[string]interface{})
	err := json.Unmarshal(message.Body, &ruleTaskInfo)
	if err != nil {
		fmt.Println("Error decoding schedule message:", err)
		return
	}
	//warnId := ruleTaskInfo["warnId"]
	ruleConfigs := ruleTaskInfo["ruleConfigs"]
	rulesNum := len(ruleConfigs)
	anomalyRulesNum := 0
	for _, ruleConfig := range ruleConfigs {
		anomaly := false
		switch ruleConfig["type"] {
		case "upDown":
			anomaly = ruleFunc.UpDownFunc(ruleConfig)
		case "dualSignalDeviation":
			anomaly = ruleFunc.DualSignalDeviationFunc(ruleConfig)
		}
		if anomaly {
			anomalyRulesNum += 1
		}
	}
	if anomalyRulesNum == rulesNum {
		//存储异常记录

	}
}

type RuleFunc struct {
}

func (rf *RuleFunc) UpDownFunc(ruleConfig *request.UpDownConfigsModel) bool {
	tagName := ruleConfig.TagName
	upThres := ruleConfig.UpThres
	downThres := ruleConfig.DownThres
	response := utils.UtilsGroup.HttpUtil.GetRequest("http://localhost:3123/api/fetch-data")
	return true
}

func (rf *RuleFunc) DualSignalDeviationFunc(ruleConfig *request.DualSignalsDeviationModel) bool {
	tagName1 := ruleConfig.TagName1
	tagName2 := ruleConfig.TagName2
	downThres := ruleConfig.Delta
	realtimeData := tagName1
	return true
}

var ruleFunc = new(RuleFunc)
