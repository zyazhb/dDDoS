package node

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	chainAddress 	string 	`yaml:"chain-address"`
	chainPort		int		`yaml:"chain-port"`

	Server struct {
		contractAddr	string	`yaml:"contract-address"`
	}

	Client struct {
		clientPrivateAddr	string `yaml:"privateAddress"`
		clientPublicAddr	string `yaml:"publicAddress"`
	}
}

const (
	configFilePath = "../config/config.yaml"
)

var (
	Conf *ClientConfig
)

func init() {
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("Read config file error!")
		return
	}

	conf := new(ClientConfig)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalln("Faile to unmarshal config file!")
		return
	}

	Conf = conf
}
