package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/product/pkg/shared/util"
	"github.com/product/internal/config/db"
	"github.com/product/internal/config/logging"
	"github.com/product/internal/config/server"

	"github.com/spf13/viper"
)

type config struct {
	Server   server.ServerList
	Database db.DatabaseList
	Logger   logging.LoggerConfig
}

var cfg config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(basepath + "/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("grpc.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server config: %v", err))
	}

	viper.AddConfigPath(basepath + "/logging")
	viper.SetConfigType("yaml")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server config: %v", err))
	}

	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("mysql.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot read database config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	viper.Unmarshal(&cfg)

	fmt.Println("=============================")
	fmt.Println(util.Stringify(viper.AllKeys()))
	fmt.Println(util.Stringify(cfg))
	fmt.Println("=============================")
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func GetConfig() *config {
	return &cfg
}
