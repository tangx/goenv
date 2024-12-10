package envutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// refFilename 根据 ref 信息返回对应的配置文件名称
func refFilename(ref string) string {
	// feat/xxxx
	parts := strings.Split(ref, "/")
	feat := parts[len(parts)-1]               // xxxx
	return fmt.Sprintf("config.%s.yml", feat) // config.xxxx.yml
}

// refConfig 根据 gitlab ci 环境变量创建与分支对应的配置文件
func refConfig() string {
	// gitlab ci
	ref := os.Getenv("CI_COMMIT_REF_NAME")

	if len(ref) != 0 {
		return refFilename(ref)
	}

	return `local.yml`
}

// Configure dump default config to default.yml
// and load config from default.yml, config.yml, refConfig, env
// prefix is AppName
// config is a pointer to a struct
func Configure(prefix string, config interface{}) error {

	// call SetDefaults, 设置默认值。 提供变量值的模版
	if err := CallSetDefaults(config); err != nil {
		return err
	}

	// write config
	data, err := Marshal(config, prefix)
	if err != nil {
		return err
	}
	_ = os.MkdirAll("./config", 0755)
	_ = os.WriteFile("./config/default.yml", data, 0644)

	// load config from files
	for _, _conf := range []string{"default.yml", "config.yml", refConfig()} {
		_conf := filepath.Join("./config/", _conf)
		log.Printf("load config from %s", _conf)

		err = UnmarshalFile(config, prefix, _conf)
		if err != nil {
			log.Println(err)
		}
	}

	// load config from env
	err = UnmarshalEnv(config, prefix)
	if err != nil {
		log.Print(err)
	}

	// Call Init, 初始化对象。 比如使用连接地址创建数据库连接。
	if err := CallInitialize(config); err != nil {
		return err
	}

	return nil
}
