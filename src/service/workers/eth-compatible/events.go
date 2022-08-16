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
	ExtraFeeSuppliedEvent    = "ExtraFeeSupplied"
	ExtraFeeTransferredEvent = "ExtraFeeTransferred"
)

var (
	ExtraFeeSuppliedEventHash = common.HexToHash("0x525223e7c9e63747e47dd4558940766054da3d0378f4006848d2a201545f55a4")
	ExtraFeeTransferredHash   = common.HexToHash("0x11d9d6d82ced6158185f9c4a4ab3c7768ed3c14cbd759491ca5d2fb42b7935fd")
)

// ExtraFeeSupplied represents a ExtraFeeSupplied event raised by the Bridge.sol contract.
type ExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResourceID         [32]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

type ExtraFeeTransferred struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResouceID          [32]byte
	Recipient          common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

func ParseETHExtraFeeSupplied(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
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

	return ev, nil
}

func ParseLAExtraFeeSupplied(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ExtraFeeTransferred
	if err := abi.UnpackIntoInterface(&ev, ExtraFeeTransferredEvent, log.Data); err != nil {
		return nil, err
	}

	fmt.Printf("ExtraFeeSupplied\n")
	fmt.Printf("origin chain ID: 0x%s\n", common.Bytes2Hex(ev.OriginChainID[:]))
	fmt.Printf("destination chain ID: 0x%s\n", common.Bytes2Hex(ev.DestinationChainID[:]))
	fmt.Printf("amount: %s\n", ev.Amount.String())
	fmt.Printf("resourceID: %s\n", common.Bytes2Hex(ev.ResouceID[:]))
	fmt.Printf("nonce: %d\n", ev.DepositNonce)

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
		EventStatus:        storage.EventStatusFeeTransferInitConfrimed,
	}
	return txlog
}

func (ev ExtraFeeTransferred) ToTxLog(chain string) *storage.TxLog {
	txlog := &storage.TxLog{
		Chain:              chain,
		TxType:             storage.TxTypeFeeTransfer,
		ReceiverAddr:       ev.Recipient.String(),
		InAmount:           ev.Amount.String(),
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		ResourceID:         common.Bytes2Hex(ev.ResouceID[:]),
		SwapID:             utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)),
		DepositNonce:       ev.DepositNonce,
		EventStatus:        storage.EventStatusFeeTransferConfirmed,
	}
	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ExtraFeeSuppliedEventHash[:]) {
		if w.GetChainName() != "LA" {
			abi, _ := abi.JSON(strings.NewReader(ethBr.EthBrABI))
			return ParseETHExtraFeeSupplied(&abi, log)
		}
	}
	if bytes.Equal(log.Topics[0][:], ExtraFeeTransferredHash[:]) {
		if w.GetChainName() == "LA" {
			abi, _ := abi.JSON(strings.NewReader(laBr.LaBrABI))
			return ParseLAExtraFeeSupplied(&abi, log)
		}
	}
	return nil, nil
}

// ContractEvent ...
type ContractEvent interface {
	ToTxLog(chain string) *storage.TxLog
}
