package v1

import (
	"encoding/json"
	"gin-basic-framework/global"
	"gin-basic-framework/model/common/request"
	"gin-basic-framework/service/baoxin"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type RuleApi struct {
}

func (ruleApi *RuleApi) CreateRuleTask(c *gin.Context) {
	ruleConfigsJSON := c.PostForm("ruleConfigs")
	triggerConfigsJSON := c.PostForm("triggerConfigs")

	// 解析 ruleConfigsJSON 和 triggerConfigsJSON 成相应的结构体
	var ruleConfigs []interface{}
	err := json.Unmarshal([]byte(ruleConfigsJSON), &ruleConfigs)
	if err != nil {
		// 处理解析错误
	}

	var triggerConfigs request.TriggerConfigsModel
	err = json.Unmarshal([]byte(triggerConfigsJSON), &triggerConfigs)
	if err != nil {
		// 处理解析错误
	}
	ruleTaskInfo := &request.RuleModel{
		WarnId:         c.PostForm("warnId"),
		RuleConfigs:    ruleConfigs,
		TriggerConfigs: triggerConfigs,
	}

	scheduleService := &baoxin.ScheduleService{}
	// 把规则任务下发给nsq根据triggerConfigs执行
	var taskInfo map[string]interface{}
	decode_err := mapstructure.Decode(ruleTaskInfo, &taskInfo)
	if err != nil {
		global.GLOBAL_LOGGER.Info("Erro decoding rulTaskInfo:", zap.Error(decode_err))
	}
	scheduleService.PublishScheduleTask(taskInfo)

}

func (ruleApi *RuleApi) StopRuleTask(c *gin.Context) {

}
