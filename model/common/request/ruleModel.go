package request

type UpDownConfigsModel struct {
	TagName   string  `mapstructure:"tagName" json:"tagName"`
	UpThres   float32 `mapstructure:"upThres" json:"upThres"`
	DownThres float32 `mapstructure:"downThres" json:"downThres"`
}

type DualSignalsDeviationModel struct {
	TagName1 string `mapstructure:"tagName1" json:"tagName1"`
	TagName2 string `mapstructure:"tagName2" json:"tagName2"`
	Delta    string `mapstructure:"delta" json:"delta"`
}

type TriggerConfigsModel struct {
	Freqs       int    `mapstructure:"freqs" json:"freqs"`
	FreqsUnit   string `mapstructure:"freqsUnit" json:"freqsUnit"`
	StartHour   int    `mapstructure:"startHour" json:"startHour"`
	StartMinute int    `mapstructure:"startMinute" json:"startMinute"`
	EndHour     int    `mapstructure:"endHour" json:"endHour"`
	EndMinute   int    `mapstructure:"endMinute" json:"endMinute"`
}

type RuleModel struct {
	WarnId         string              `mapstructure:"warnId" json:"warnId"`
	RuleConfigs    []interface{}       `mapstructure:"ruleConfigs" json:"ruleConfigs"`
	TriggerConfigs TriggerConfigsModel `mapstructure:"triggerConfigs" json:"triggerConfigs"`
}
