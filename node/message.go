package node

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"main/node/contract"
)

const (
	voteTransHash    = "0xe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e"
	trafficTransHash = "0x82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca8007577"
)

var (
	votingCount = make(map[uint64]int)
	MLChan      = make(chan contract.TrafficStationupchainTrafficInfo, 20)
)

func WatchMessage() {
	log.Println("[-] Start watch message")
	done := make(chan bool)

	contractAddress := common.HexToAddress(Conf.Server.ContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := Client.SubscribeFilterLogs(context.Background(), query, logs)
	defer sub.Unsubscribe()
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	if err != nil {
		log.Fatal(err)
	}

	// 获取监听event的返回值
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				eventKeccak256 := vLog.Topics[0].String()
				// traffic trans -> ML
				if eventKeccak256 == trafficTransHash {
					ret, err := contractAbi.Unpack("trafficTrans", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					retPack := tidyTrafficToStruct(ret)

					if retPack.SourceAddr.String() != Conf.Client.ClientPublicAddr {
						WriteJsonMessage(Message{
							TypeName: RECEIVEEDTYPE,
							Content: map[string]MessageType{
								"trafficID": MessageType(retPack.TrafficID.Text(10)),
								"sourceAddr": MessageType(retPack.SourceAddr.String()),
								"trafficInfo": MessageType("80,6," + time.Now().Format("15:04:05 2006/01/02") + "9151646,0"),
								"targetDevice": "智能微波炉",
								"targetDeviceType": "智能家居",
							},
						})

						MLChan <- retPack
					}
				} else if eventKeccak256 == voteTransHash {
					ret, err := contractAbi.Unpack("voteTrans", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					retPack := tidyVoteToStruct(ret)
					if retPack.SourceAddr.String() == Conf.Client.ClientPublicAddr {
						votingCount[retPack.TrafficID.Uint64()] += 1

						if votingCount[retPack.TrafficID.Uint64()] == 1 {
							if ok := judgeTrafficFromVote(retPack.TrafficID); ok {
								WriteJsonMessage(Message{
									TypeName: VERIFIEDTYPE,
									Content: nil,
								})

								dealWithDDoSTraffic()
							}
						}
					}
				}
			}
		}
	}()

	go transTrafficInfoToML()

	<-done
}

/*
mlResult := transTrafficInfoToML(retPack)
						// TODO: create message channel as buffer
						SendVote(retPack.TrafficID, retPack.SourceAddr, mlResult)
*/

func transTrafficInfoToML() bool {
	for {
		select {
		case c := <-MLChan:
			WriteJsonMessage(Message{
				TypeName: MLTYPE,
				Content: map[string]MessageType{
					"state":       "get",
					"attackType":  "DDoS",
					"trafficType": "SYN",
				},
			})
			// do something here
			SendVote(c.TrafficID, c.SourceAddr, true)
		}
	}
}

func judgeTrafficFromVote(trafficID *big.Int) bool {
	voteNum, err := pendingVotingNumAt(context.Background(), trafficID)
	if err != nil {
		log.Fatalln("[x] Failed to get vote num!")
	}

	if voteNum.Uint64() == 1 {
		return true
	} else {
		return false
	}
}

func pendingVotingNumAt(ctx context.Context, trafficID *big.Int) (*big.Int, error) {
	auth := &bind.CallOpts{
		From:    common.HexToAddress(Conf.Client.ClientPublicAddr),
		Context: ctx,
	}

	result, err := Instance.GetVoteInfoByID(auth, trafficID)
	if err != nil {
		log.Fatalln("[x] Failed to get vote result!")
		return nil, err
	}

	return result, nil
}

func pendingTrafficIDAt(ctx context.Context, address string) (*big.Int, error) {
	auth := &bind.CallOpts{
		From:    common.HexToAddress(Conf.Client.ClientPublicAddr),
		Context: ctx,
	}

	addr := common.HexToAddress(address) // 需要被查询的地址
	result, err := Instance.PendingTrafficID(auth, addr)
	if err != nil {
		log.Fatalln("[x] Failed to update traffic id!")
		return nil, err
	}

	return result, nil
}

func tidyTrafficToStruct(retrieveValue []interface{}) contract.TrafficStationupchainTrafficInfo {
	mod := contract.TrafficStationupchainTrafficInfo{
		TrafficID:   retrieveValue[0].(*big.Int),
		SourceAddr:  retrieveValue[1].(common.Address),
		TrafficInfo: retrieveValue[2].(string),
	}

	return mod
}

func tidyVoteToStruct(retrieveValue []interface{}) contract.TrafficStationvoteInfo {
	mod := contract.TrafficStationvoteInfo{
		SourceAddr: retrieveValue[0].(common.Address),
		VoteAddr:   retrieveValue[1].(common.Address),
		TrafficID:  retrieveValue[2].(*big.Int),
		State:      retrieveValue[3].(bool),
	}

	return mod
}
