package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	laBr "github.com/latoken/bridge-balancer-service/src/service/workers/eth-compatible/abi/bridge/la"
	"github.com/latoken/bridge-balancer-service/src/service/workers/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	jsonrpc "github.com/ybbus/jsonrpc/v2"

	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"
)

// Erc20Worker ...
type Erc20Worker struct {
	provider           string
	chainName          string
	chainID            int64
	destinationChainID string
	logger             *logrus.Entry // logger
	config             *models.WorkerConfig
	client             *ethclient.Client
	contractAddr       common.Address
}

// NewErc20Worker ...
func NewErc20Worker(logger *logrus.Logger, cfg *models.WorkerConfig) *Erc20Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic(fmt.Sprintf("rpc error for %s : %s", cfg.ChainName, err.Error()))
	}

	privKey, err := utils.GetPrivateKey(cfg)
	if err != nil {
		panic(fmt.Sprintf("generate private key error, err=%s", err.Error()))
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("get public key error, err=%s", err.Error()))
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if !bytes.Equal(cfg.WorkerAddr.Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.WorkerAddr, fromAddress,
		))
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		panic("rpc not returning chain id")
	}

	// init token addresses
	return &Erc20Worker{
		chainName:          cfg.ChainName,
		chainID:            chainid.Int64(),
		destinationChainID: cfg.DestinationChainID,
		logger:             logger.WithField("worker", cfg.ChainName),
		provider:           cfg.Provider,
		config:             cfg,
		client:             client,
		contractAddr:       cfg.ContractAddr,
	}
}

// GetChainName returns chain ID
func (w *Erc20Worker) GetChainName() string {
	return w.chainName
}

// GetChainName returns chain ID
func (w *Erc20Worker) GetChainID() string {
	return fmt.Sprint(w.chainID)
}

// returns destinationChainID to be checked
func (w *Erc20Worker) GetDestinationID() string {
	return w.destinationChainID
}

// GetStartHeight returns start blockchain height from config
func (w *Erc20Worker) GetStartHeight() (int64, error) {
	if w.config.StartBlockHeight == 0 {
		return w.GetHeight()
	}
	return w.config.StartBlockHeight, nil
}

// GetConfirmNum returns numbers of blocks after them tx will be confirmed
func (w *Erc20Worker) GetConfirmNum() int64 {
	return w.config.ConfirmNum
}

func (w *Erc20Worker) TransferExtraFee(originChainID, destinationChainID [8]byte, nonce uint64, resourceID [32]byte, receiptAddr string, amount string) (string, error) {
	auth, err := w.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := laBr.NewLaBr(w.contractAddr, w.client)
	if err != nil {
		return "", err
	}
	value, _ := new(big.Int).SetString(amount, 10)
	tx, err := instance.TransferExtraFee(auth, originChainID, destinationChainID, nonce, resourceID, common.HexToAddress(receiptAddr), value)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return tx.Hash().String(), nil

}

// GetStatus returns status of relayer: blockchain; account(address, balance ...)
func (w *Erc20Worker) GetStatus() (*models.WorkerStatus, error) {
	status := &models.WorkerStatus{}
	// set current block height
	height, err := w.GetHeight()
	if err != nil {
		return nil, err
	}
	status.Height = height
	// set worker address
	status.Account.Address = w.config.WorkerAddr.Hex()

	return status, nil
}

// // GetStatus returns status of relayer account(balance eg)
// func (w *Erc20Worker) GetStatus(symbol string) (interface{}, error) {
// 	ethStatus := &EthStatus{}

// 	allowance, err := w.Allowance(symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Allowance = QuoBigInt(allowance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	balance, err := w.Erc20Balance(w.Config.WorkerAddr, symbol)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.Erc20Balance = QuoBigInt(balance, GetBigIntForDecimal(w.Config.ChainDecimal)).String()

// 	ethBalance, err := w.EthBalance(w.Config.WorkerAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethStatus.EthBalance = QuoBigInt(ethBalance, GetBigIntForDecimal(18)).String()

// 	return ethStatus, nil
// }

// GetBlockAndTxs ...
func (w *Erc20Worker) GetBlockAndTxs(height int64) (*models.BlockAndTxLogs, error) {
	clientResp, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		w.logger.Errorln("Error while fetching the block header = ", err)
		return nil, err
	}

	nextHeight := clientResp.Number.Int64()
	if height >= nextHeight {
		return nil, fmt.Errorf("not found")
	} else if height == 0 {
		height = nextHeight - 1
	} else if nextHeight-height >= 100 {
		nextHeight = height + 20
	}

	logs, err := w.getLogs(height, nextHeight)
	if err != nil {
		w.logger.Errorf("while getEvents(block number from %d to %d), err = %v", height, nextHeight, err)
		return nil, err
	}

	return &models.BlockAndTxLogs{
		Height:          nextHeight,
		BlockHash:       clientResp.Hash().String(),
		ParentBlockHash: clientResp.ParentHash.Hex(),
		BlockTime:       int64(clientResp.Time),
		TxLogs:          logs,
	}, nil
}

// GetFetchInterval ...
func (w *Erc20Worker) GetFetchInterval() time.Duration {
	return time.Duration(w.config.FetchInterval) * time.Second
}

// getLogs ...
func (w *Erc20Worker) getLogs(curHeight, nextHeight int64) ([]*storage.TxLog, error) {
	//	topics := [][]common.Hash{{DepositEventHash, ProposalEventHash, ProposalVoteHash}}

	logs, err := w.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(curHeight + 1),
		ToBlock:   big.NewInt(nextHeight),
		// Topics:    topics,
		Addresses: []common.Address{w.contractAddr},
		Topics:    [][]common.Hash{},
	})
	if err != nil {
		w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("get event log error, err=%s", err)
		return nil, err
	}

	models := make([]*storage.TxLog, 0, len(logs))
	for _, log := range logs {
		w.logger.Infof("WORKER(%s) NEW EVENT: %v\n\n", w.chainName, log)
		event, err := w.parseEvent(&log)
		if err != nil {
			w.logger.WithFields(logrus.Fields{"function": "GetLogs()"}).Errorf("parse event log error, err=%s", err)
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog(w.chainName)
		txLog.Chain = w.chainName
		txLog.Height = int64(log.BlockNumber)
		txLog.EventID = log.TxHash.Hex()
		txLog.BlockHash = log.BlockHash.Hex()
		txLog.TxHash = log.TxHash.Hex()
		txLog.Status = storage.TxStatusInit

		models = append(models, txLog)
	}

	return models, nil
}

// GetHeight ..
func (w *Erc20Worker) GetHeight() (int64, error) {
	header, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, nil
	}
	return header.Number.Int64(), nil
}

// GetSentTxStatus ...
func (w *Erc20Worker) GetSentTxStatus(hash string) storage.TxStatus {
	txReceipt, err := w.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		_, isPending, err := w.client.TransactionByHash(context.Background(), common.HexToHash(hash))
		if err != nil {
			if err == ethereum.NotFound {
				return storage.TxSentStatusLost
			}
			return storage.TxSentStatusNotFound
		}
		if isPending {
			return storage.TxSentStatusPending
		}
		return storage.TxSentStatusNotFound
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return storage.TxSentStatusFailed
	}

	return storage.TxSentStatusSuccess
}

func (w *Erc20Worker) GetTxCountLatest() (uint64, error) {
	var result hexutil.Uint64
	rpcClient := jsonrpc.NewClient(w.provider)

	resp, err := rpcClient.Call("eth_getTransactionCount", w.config.WorkerAddr.Hex(), "latest")
	if err != nil {
		return 0, err
	}
	if err := resp.GetObject(&result); err != nil {
		return 0, err
	}
	return uint64(result), nil
}

// GetTransactor ...
func (w *Erc20Worker) getTransactor() (auth *bind.TransactOpts, err error) {
	privateKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return nil, err
	}

	var nonce uint64

	nonce, err = w.client.PendingNonceAt(context.Background(), w.config.WorkerAddr)
	if err != nil {
		return nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(w.chainID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.config.GasLimit) // in units
	auth.GasPrice = w.config.GasPrice

	return auth, nil
}

// EthBalance ...
func (w *Erc20Worker) EthBalance(address common.Address) (*big.Int, error) {
	return w.client.BalanceAt(context.Background(), address, nil)
}

// GetWorkerAddress ...
func (w *Erc20Worker) GetWorkerAddress() string {
	return w.config.WorkerAddr.String()
}

// GetColdWalletAddress ...
func (w *Erc20Worker) GetColdWalletAddress() string {
	return w.config.ColdWalletAddr.String()
}

// IsSameAddress ...
func (w *Erc20Worker) IsSameAddress(addrA string, addrB string) bool {
	return bytes.Equal(common.FromHex(addrA), common.FromHex(addrB))
}

// SendAmount ...
func (w *Erc20Worker) SendAmount(address string, amount *big.Int) (string, error) {
	return "", fmt.Errorf("not implemented") // TODO
}
