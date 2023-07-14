package baoxin

import (
	"encoding/json"
	"fmt"
	"gin-basic-framework/global"
	"gin-basic-framework/utils"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"time"
)

type ScheduleService struct {
}

func (ss *ScheduleService) PublishScheduleTask(taskInfo map[string]interface{}) {
	message := []byte("Schedule rule task")
	message = append(message, utils.UtilsGroup.EncodeJSON(taskInfo)...)

	err := global.GLOBAL_NSQ_PRODUCER.Publish("schedule_task", message)

	if err != nil {
		global.GLOBAL_LOGGER.Fatal("Error publishing schedule_task", zap.Error(err))
	}

	global.GLOBAL_LOGGER.Info("schedule_task has been published successfully")
}

type ScheduleHandler struct {
}

func (sh *ScheduleHandler) HandleSchedule(message *nsq.Message) {
	scheduleTaskInfo := make(map[string]interface{})
	err := json.Unmarshal(message.Body, &scheduleTaskInfo)
	if err != nil {
		fmt.Println("Error decoding schedule message:", err)
		return
	}
	// 根据triggerConfigs的内容获取开始时间和结束时间
	startTime := time.Now()
	endTime := time.Now()

	frequency := 1 * time.Second
	sleepDuration := startTime.Sub(time.Now())
	if sleepDuration > 0 {
		global.GLOBAL_LOGGER.Info("Waiting for start time ")
		time.Sleep(sleepDuration)
	}
	for {
		if time.Now().After(endTime) {
			global.GLOBAL_LOGGER.Info("")
			break
		}

		time.Sleep(frequency)
	}
}
