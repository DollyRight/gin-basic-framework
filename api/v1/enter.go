package v1

type ApiGroup struct {
	UserApi
	RuleApi
}

var ApiGroupApp = new(ApiGroup)
