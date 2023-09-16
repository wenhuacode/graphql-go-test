package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go-graphql-test/global"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}

func InitConfig() {
	//从配置文件中读取出对应的配置
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-debug.yaml", configFilePrefix)
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息: %v", global.Config)
}
