package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var SystemConfig Config

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Mode   string `yaml:"mode"`
	Port   string `yaml:"port"`
	JwtKey string `yaml:"jwtKey"`
}

type Db struct {
	Type     string `yaml:"type"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

func init() {
	var configFile []byte
	var err error

	configFile, err = ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("can not find the config file:", err)
	}

	yaml.Unmarshal(configFile, &SystemConfig)
}

// func getPath() string {
// 	file, _ := exec.LookPath(os.Args[0])
// 	path, _ := filepath.Abs(file)
// 	index := strings.LastIndex(path, string(os.PathSeparator))
// 	path = path[:index]
// 	return path
// }
