package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LinkinStars/go-scaffold/contrib/static_config/file"
	"github.com/LinkinStars/golang-util/gu"
)

var (
	GlobalConf = &Conf{}
	DataPath   = ""
)

const (
	configFilename = "config.yaml"
	defaultConfig  = `WebPort: ":8080"
Username: "admin"
Password: "admin"
Nickname: "ADMIN"
Avatar: "/card/icon/favicon-60.png"
Secret: "your-secret"
SiteName: "DailyCards"
ShowQrcode: true
`
)

type Conf struct {
	WebPort        string
	Username       string
	Password       string
	SiteName       string
	ShowQrcode     bool
	Nickname       string
	Avatar         string
	Secret         string
	HideLinkCorner bool
}

func setDefault() {
	if len(DataPath) == 0 {
		DataPath, _ = os.UserHomeDir()
		DataPath = filepath.Join(DataPath, "daily-cards")
	}
	if err := gu.CreateDirIfNotExist(DataPath); err != nil {
		panic(err)
	}
}

func Init() {
	setDefault()
	conf := filepath.Join(DataPath, configFilename)
	configParser := file.NewStaticConfigParser(conf)
	err := configParser.LoadAndSet(GlobalConf)
	if err != nil {
		panic(err)
	}
}

func WriteConfigFile() (err error) {
	setDefault()
	configFilePath := filepath.Join(DataPath, configFilename)
	if gu.CheckPathIfNotExist(configFilePath) {
		return fmt.Errorf("config file [%s] is already exist", configFilePath)
	}
	f, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	writer := bufio.NewWriter(f)
	if _, err := writer.WriteString(defaultConfig); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}
