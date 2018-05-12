package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Value これはviperで取得した設定をsubしたものを格納
var Value *viper.Viper

// Set 設定ファイルを読み込みます
func Set() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s ", err))
	}
	if os.Getenv("ENV") == "development" {
		Value = viper.Sub("development")
	} else {
		Value = viper.Sub("production")
	}
	return nil
}
