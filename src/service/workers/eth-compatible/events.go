package eth

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/latoken/bridge-balancer-service/src/service/storage"
	ethBr "github.com/latoken/bridge-balancer-service/src/service/workers/eth-compatible/abi/bridge/eth"
	laBr "github.com/latoken/bridge-balancer-service/src/service/workers/eth-compatible/abi/bridge/la"
	"github.com/latoken/bridge-balancer-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

const (
	ExtraFeeSuppliedEvent = "ExtraFeeSupplied"
)

var (
	ExtraFeeEventHash = common.HexToHash("0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db")
)

// ExtraFeeSupplied represents a ExtraFeeSupplied event raised by the Bridge.sol contract.
type ExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResourceID         [32]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Status             uint8
	Params             []byte
	Raw                types.Log // Blockchain specific contextual infos
}

func ParseExtraFeeSupplied(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ExtraFeeSupplied
	if err := abi.UnpackIntoInterface(&ev, ExtraFeeSuppliedEvent, log.Data); err != nil {
		return nil, err
	}

	fmt.Printf("ExtraFeeSupplied\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("amount: %s\n", ev.Amount.String())
	fmt.Printf("resourceID: %s\n", common.Bytes2Hex(ev.ResourceID[:]))
	fmt.Printf("nonce: %d\n", ev.DepositNonce)
	fmt.Printf("status: %d\n", ev.Status)
	fmt.Printf("params: %s\n", common.Bytes2Hex(ev.Params[:]))

	return ev, nil
}

// ToTxLog ...
func (ev ExtraFeeSupplied) ToTxLog(chain string) *storage.TxLog {
	txlog := &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeFeeTransfer,
		ReceiverAddr:       ev.RecipientAddress.String(),
		InAmount:           ev.Amount.String(),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		SwapID:             utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)),
		DepositNonce:       ev.DepositNonce,
		Data:               common.Bytes2Hex(ev.Params[:]),
		EventStatus:        storage.EventStatusFeeTransferInit,
	}
	if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeFeeTransferConfirm
		txlog.EventStatus = storage.EventStatusFeeTransferConfirmed
	} else if ev.Status == uint8(4) {
		txlog.TxType = storage.TxTypeFeeReversal
		txlog.EventStatus = storage.EventStatusFeeTransferReversed
	}
	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ExtraFeeEventHash[:]) {
		var contractAbi abi.ABI
		if w.GetChainName() != "LA" {
			contractAbi, _ = abi.JSON(strings.NewReader(ethBr.EthBrABI))
		} else {
			contractAbi, _ = abi.JSON(strings.NewReader(laBr.LaBrABI))
		}
		return ParseExtraFeeSupplied(&contractAbi, log)
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
