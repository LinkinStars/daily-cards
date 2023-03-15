package main

import (
	"flag"
	"fmt"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/db"
	"github.com/LinkinStars/dc/internal/base/router"
	"github.com/LinkinStars/go-scaffold/contrib/log/zap"
	"github.com/LinkinStars/go-scaffold/logger"
)

var (
	initConfig bool
	Version    = "v0.0.0"
)

func init() {
	flag.BoolVar(&initConfig, "init", false, "init config file")
	flag.StringVar(&config.DataPath, "c", "", "save data path")
}

func main() {
	showVersion := flag.Bool("v", false, "print version")
	flag.Parse()
	if *showVersion {
		fmt.Println(Version)
		return
	}

	if initConfig {
		err := config.WriteConfigFile()
		if err != nil {
			fmt.Println("init config failed: ", err.Error())
		} else {
			fmt.Printf("init config at [%s] success\n", config.DataPath)
		}
		return
	}
	config.Init()

	newLogger := zap.NewLogger(logger.LevelDebug,
		zap.WithName("daily-cards-log"), zap.WithPath(config.DataPath))
	logger.SetLogger(newLogger)

	db.Init()

	router.Run()
}
