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

func WatchMessage() {
	contractAddress := common.HexToAddress(Conf.Server.contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	if err != nil {
        log.Fatal(err)
    }

	for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
			if vLog.Topics[0].String() == trafficTransHash {
				// except source address
                _, err := contractAbi.Unpack("trafficTrans", vLog.Data)
                if err != nil {
                    log.Fatal(err)
                }

				// traffic from source -> to ML

				// TODO: function traffic <-> ML | param: traffic info | return: answer bool
            } else if vLog.Topics[0].String() == voteTransHash {
				// only source address can get it
                _, err := contractAbi.Unpack("voteTrans", vLog.Data)
                if err != nil {
                    log.Fatal(err)
                }
            }
        }
    }
}
