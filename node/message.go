package node

import (
	"context"
	"log"
	"strings"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"main/node/contract"
)

const (
	voteTransHash	 = "0xe6d2d2c21284e671da4a773c628d519f81d37af2ae3950c34d207f567bb08f4e"
	trafficTransHash = "0x82a0cd03c34dff9e0879fd6d40cea94cdf58b4b4bbf4b2c8a3ebb91ca8007577"
)

type UpchainTrafficInfo struct {
	TrafficID *big.Int
	SourceAddr common.Address
	TrafficInfo string
}

type VoteInfo struct {
	SourceAddr common.Address
	VoteAddr common.Address
	TrafficID *big.Int
	State bool
}

var (
	votingCount map[uint64]int
)

func WatchMessage() {
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
				if vLog.Topics[0].String() == trafficTransHash {
					ret, err := contractAbi.Unpack("trafficTrans", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					retPack := tidyTrafficToStruct(ret)

					if (retPack.SourceAddr.String() != Conf.Client.ClientPrivateAddr) {
						mlResult := transTrafficInfoToML(retPack)
						// TODO: create message channel as buffer
						SendVote(retPack.TrafficID, retPack.SourceAddr, mlResult)
					}
				} else if vLog.Topics[0].String() == voteTransHash {
					ret, err := contractAbi.Unpack("voteTrans", vLog.Data)
					if err != nil {
						log.Fatal(err)
					}
					retPack := tidyVoteToStruct(ret)
					if (retPack.SourceAddr.String() == Conf.Client.ClientPublicAddr) {
						votingCount[retPack.TrafficID.Uint64()] += 1

						if (votingCount[retPack.TrafficID.Uint64()] == 1) {
							judgeTrafficFromVote(retPack.TrafficID)
						}
					}
				}
			}
		}
	}()

	<-done
}

func transTrafficInfoToML(info contract.TrafficStationupchainTrafficInfo) bool {
	return true
}

func judgeTrafficFromVote(trafficID *big.Int) {
	voteNum, err := pendingVotingNumAt(context.Background(), trafficID)
	if err != nil {
		log.Fatalln("[x] Failed to get vote num!")
	}

	// TODO: 这里硬编码的数量要改成动态获取peer的数量
	if voteNum.Uint64() == 1 {
		// TODO: 这里的话一个是iptables规则处理，一个是发送预警 -> 新建event进行处理
	} else {
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
