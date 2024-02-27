package initialize

import (
	"github.com/spf13/viper"
	"mxshop/user/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//设置环境变量，通过true和false控制用哪个config文件
}

func InitConfig() {
	//从配置文件中读取响应的配置
	debug := GetEnvInfo("MX_shop")
	var ConfigName string
	if debug {
		ConfigName = "user/config-debug.yaml"
	} else {
		ConfigName = "user/config-pro.yaml"
	}
	v := viper.New()
	v.SetConfigFile(ConfigName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
