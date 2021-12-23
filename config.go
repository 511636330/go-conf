package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type StrMap map[string]interface{}

func init() {
	Viper = viper.New()
	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig()

	if err != nil {
		glog.Errorf("Init conifg error: %v", err)

	}

	Viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()

	viper.OnConfigChange(func(e fsnotify.Event) {
		glog.Info(Viper.AllSettings())
	})
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}

func GetDuration(path string, defaultValue ...interface{}) time.Duration {
	return cast.ToDuration(Get(path, defaultValue...))
}

func GetStringMap(path string, defaultValue ...interface{}) map[string]interface{} {
	return cast.ToStringMap(Get(path, defaultValue...))
}
