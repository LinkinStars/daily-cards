package db

import (
	"path/filepath"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/model"
	"github.com/LinkinStars/go-scaffold/logger"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	Engine *xorm.Engine
)

func Init() {
	conn := filepath.Join(config.DataPath, "daily-cards.db")
	logger.Debugf("try to connect to database %s", conn)
	engine, err := xorm.NewEngine("sqlite", conn)
	if err != nil {
		panic(err)
	}
	err = engine.Ping()
	if err != nil {
		panic(err)
	}
	engine.SetColumnMapper(names.GonicMapper{})

	err = engine.Sync(model.Tables...)
	if err != nil {
		panic(err)
	}
	Engine = engine
}
