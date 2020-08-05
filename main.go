package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/sunyd/go-demo1/common"
	"os"
)

func main() {

	//加载配置文件
	InitConfig()

	db := common.InitSqlDb()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	//默认设置端口8080
	//r.Run()
}

func InitConfig() {
	workDir, _ := os.Getwd()
	fmt.Printf("工作目录为：%v \n", workDir)
	//设置读取配置文件名
	viper.SetConfigName("application")
	// 设置配置文件类型
	viper.SetConfigType("yml")
	//设置配置文件路径,搜索第一个路径，可以设置多个路径
	viper.AddConfigPath(workDir + "/config")
	// 搜索路径并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监视配置文件，重新读取配置数据
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
}
