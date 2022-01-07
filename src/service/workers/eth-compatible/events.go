package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/storage"
	ethBr "gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/eth-compatible/abi/bridge/eth"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	ExtraFeeEvent = "ExtraFeeSupplied"
)

var (
	ExtraFeeEventHash = common.HexToHash("0xac4619eda7a5583d586072607ccbcff24908b067d39a59e2ebcdf3509aa65d2e")
)

// ExtraFeeSupplied represents a ExtraFeeSupplied event raised by the Bridge.sol contract.
type ExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

func ParseLAExtraFeeSupplied(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ExtraFeeSupplied
	if err := abi.UnpackIntoInterface(&ev, ExtraFeeEvent, log.Data); err != nil {
		return nil, err
	}

	fmt.Printf("ExtraFeeSupplied\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("amount: ", ev.Amount.String())

	return ev, nil
}

// !!! TODO !!!

// ToTxLog ...
func (ev ExtraFeeSupplied) ToTxLog(chain string) *storage.TxLog {
	txlog := &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeFeeTransfer,
		ReceiverAddr:       ev.RecipientAddress.String(),
		InAmount:           ev.Amount.String(),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
	}
	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ExtraFeeEventHash[:]) {
		if w.GetChainName() != storage.LaChain {
			abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
			return ParseLAExtraFeeSupplied(&abi, log)
		}
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
