package app

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Config appConfig
)

type appConfig struct {
	Http struct {
		Domain string `mapstructure:"domain"`
		Port   string `mapstructure:"port"`
	} `mapstructure:"http"`
	Rpc struct {
		Domain string `mapstructure:"domain"`
		Port   string `mapstructure:"port"`
	} `mapstructure:"rpc"`
	DB struct {
		Host               string `mapstructure:"host"`
		Port               string `mapstructure:"port"`
		User               string `mapstructure:"user"`
		Password           string `mapstructure:"password"`
		Name               string `mapstructure:"name"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections"`
		MaxOpenConnections int    `mapstructure:"max_idle_connections"`
	} `mapstructure:"db"`
	Secret struct {
		JwtKey       string `mapstructure:"auth_key"`
		PassHashSalt string `mapstructure:"pass_hash_salt"`
	} `mapstructure:"secret"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		PoolSize int    `mapstructure:"pool_size"`
	} `mapstructure:"redis"`
	Kuaidi100 struct {
		CustomerID string `mapstructure:"customer_id"`
		Key        string `mapstructure:"key"`
		Url        string `mapstructrue:"url"`
	} `mapstructure:"kuaidi100"`
	UserRpc struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"user_rpc"`
	Pingpp struct {
		AppID      string `mapstructure:"app_id"`
		Key        string `mapstructure:"key"`
		PrivateKey string `mapstructure:"private_key"`
	} `mapstructure:"pingpp"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	fmt.Println("config", Config)

	//remoteConfig := config.Get("order-server")
	//fmt.Println("get remote config", remoteConfig)
	//
	//viper.SetConfigType("yml")
	//viper.ReadConfig(bytes.NewBufferString(remoteConfig))
	//if err := viper.Unmarshal(&Config); err != nil {
	//	panic(err)
	//}
}
