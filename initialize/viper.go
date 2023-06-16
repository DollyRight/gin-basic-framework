package initialize

import (
	"fmt"
	"gin-basic-framework/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	vp := viper.New()
	//config文件的路径、名称、类型
	vp.SetConfigFile("config.yaml")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal eroor config file: %s \n", err))
	}
	vp.WatchConfig()

	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = vp.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = vp.Unmarshal(&global.GLOBAL_CONFIG); err != nil {
		fmt.Println(err)
	}

	return vp

}
