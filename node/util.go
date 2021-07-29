package node

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	ChainAddress 	string 	`yaml:"chain-address"`
	ChainPort		int		`yaml:"chain-port"`

	Server struct {
		ContractAddr	string	`yaml:"contract-address"`
	}

	Client struct {
		ClientPrivateAddr	string `yaml:"privateAddress"`
		ClientPublicAddr	string `yaml:"publicAddress"`
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
