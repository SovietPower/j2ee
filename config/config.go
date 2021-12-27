package config

import (
	"j2ee/model"
	"j2ee/util/logging"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		logging.Info(err)
		panic(err)
	}

	// 连接数据库
	model.InitDatabase(os.Getenv("MYSQL_DSN"))
	// model.InitDatabase("")
}
