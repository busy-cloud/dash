package main

import (
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("dash")

	//顺序启动
	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
	}

	//安全退出
	defer boot.Shutdown()

	//启动Web服务
	err = web.Serve()
	if err != nil {
		log.Error(err)
	}

}
