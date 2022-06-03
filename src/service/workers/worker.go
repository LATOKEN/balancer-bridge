package workers

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"
)

// IWorker ...
type IWorker interface {
	// GetChain returns unique name of the chain(like LA, ETH and etc)
	GetChainID() string
	GetChainName() string
	GetDestinationID() string
	// GetWokrerAddress returns worker address
	GetWorkerAddress() string
	// GetStartHeight returns blockchain start height for watcher
	GetStartHeight() (int64, error)
	// GetConfirmNum returns numbers of blocks after them tx will be confirmed
	GetConfirmNum() int64
	// GetHeight returns current height of chain
	GetHeight() (int64, error)
	// GetBlockAndTxs returns block info and txs included in this block
	GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error)
	// GetFetchInterval returns fetch interval of the chain like average blocking time, it is used in observer
	GetFetchInterval() time.Duration
	// GetSentTxStatus returns status of tx sent
	GetSentTxStatus(hash string) storage.TxStatus
	// IsSameAddress returns is addrA the same with addrB
	IsSameAddress(addrA string, addrB string) bool
	// GetStatus returns status of relayer: blockchain; account(address, balance ...)
	GetStatus() (*models.WorkerStatus, error)

	//TransferExtraFee to be called on lachain side to transfer
	TransferExtraFee(originChainID, destinationChainID [8]byte, nonce uint64, resourceID [32]byte, receiptAddr string, amount string) (string, error)

	CreateMessageHash(amount, recipientAddress, originChainID string) (common.Hash, error)

	CreateSignature(messageHash common.Hash, chainId string) (string, error)
}
