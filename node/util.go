package node

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"main/server/wserver"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	ChainAddress string `yaml:"chain-address"`
	ChainPort    string `yaml:"chain-port"`

	Server struct {
		ContractAddr string `yaml:"contract-address"`
	}

	Client struct {
		ClientPrivateAddr string `yaml:"privateAddress"`
		ClientPublicAddr  string `yaml:"publicAddress"`
	}
}

const (
	configFilePath = "../config/config1.yaml"
)

const (
	DETECTTYPE    = "detect type"
	VERIFIEDTYPE  = "verified type"
	RECEIVEEDTYPE = "received type"
	MLTYPE        = "machine learning"
	VOTETYPE      = "vote type"
	DEALINGTYPE   = "dealing type"
)

var (
	Conf      *ClientConfig
	Connector wserver.WsCat = wserver.WsCat{}
)

func init() {
	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("[x] Read config file error!")
		return
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Fatalln("[x] Faile to unmarshal config file!")
		return
	}

	err = Connector.Connect("ws://" + Conf.ChainAddress + "/websocket/ws")
	if err != nil {
		log.Fatalln("[x] Connect to websocket server error")
		return
	}
}

type MessageType string

type Message struct {
	TypeName string
	Content  map[string]MessageType
}

func WriteMessage(p string) {
	Connector.WriteMessage(p)
}

func WriteJsonMessage(mt Message) {
	jp, err := json.Marshal(&mt)
	if err != nil {
		log.Fatalln("[x] Marshal content error")
	}

	Connector.WriteJsonMessage(jp)
}

func dealWithDDoSTraffic() {
	WriteJsonMessage(Message{
		TypeName: DEALINGTYPE,
		Content: nil,
	})
}
