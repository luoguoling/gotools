
package main

import (
	"checkMds/checklibs"
	"checkMds/config"
	"flag"
	"fmt"
	"os"
	"time"
)


var configPath string

func main(){
	flag.StringVar(&configPath, "config", "config/config.yaml", "assign your config file: -config=your_config_file_path.")
	flag.Parse()
	if err := config.InitConfig(configPath); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	for {
		checklibs.CheckAll()
		time.Sleep(time.Duration(config.GetConfig().Interval) * time.Second)

	}
}