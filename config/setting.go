package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Conf struct {
	Server ServerConfig `mapstructure:"server"`
	System SystemConfig `mapstructure:"system"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Log    LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Port int64  `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type SystemConfig struct {
	PwdKey PwdKeyConfig `mapstructure:"pwdKey"`
	Qn     QnConfig     `mapstructure:"qn"`
}

type PwdKeyConfig struct {
	PublicKey  string `mapstructure:"publicKey"`
	PrivateKey string `mapstructure:"privateKey"`
}

type QnConfig struct {
	Key       string `mapstructure:"key"`
	Bucket    string `mapstructure:"bucket"`
	AccessKey string `mapstructure:"access-key"`
	SecretKey string `mapstructure:"secret-key"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int64  `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	LogMode      bool   `mapstructure:"log_mode"`
	SqlLog       bool   `json:"sqlLog"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

var GlobalConf = new(Conf)

func InitSetting() {
	viper.SetConfigType("yml")
	viper.SetConfigFile("./conf/conf.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("读取配置文件失败，", err)
	}

	// 映射到结构体
	err = viper.Unmarshal(GlobalConf)
	if err != nil {
		log.Fatalln("参数映射到结构体失败，", err)
	}

	// 添加监听
	viper.WatchConfig()
	// 添加回调逻辑
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 映射到结构体
		err = viper.Unmarshal(GlobalConf)
		if err != nil {
			log.Println("参数映射到结构体失败，", err)
		}
		log.Printf("修改后的配置文件为：%+v", GlobalConf)
	})

}
