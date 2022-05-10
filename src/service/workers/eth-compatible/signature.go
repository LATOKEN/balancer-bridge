package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/latoken/bridge-balancer-service/src/service/workers/utils"
)

func (w *Erc20Worker) CreateMessageHash(amount, recipientAddress, destinationChainID string) (common.Hash, error) {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	bytesTy, _ := abi.NewType("bytes8", "bytes8", nil)

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: bytesTy,
		},
	}
	value, _ := new(big.Int).SetString(amount, 10)
	bytes, err := arguments.Pack(
		value,
		common.HexToAddress(recipientAddress),
		utils.StringToBytes8(destinationChainID),
	)
	if err != nil {
		return common.Hash{}, err
	}
	messageHash := crypto.Keccak256Hash(bytes)
	return messageHash, nil
}

func (w *Erc20Worker) CreateSignature(messageHash common.Hash) (string, error) {
	privKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return "", err
	}
	signature, er := crypto.Sign(messageHash.Bytes(), privKey)
	if er != nil {
		return "", er
	}
	signature[64] = signature[64] + 35 + byte(w.chainID)*2
	return hexutil.Encode(signature), nil
}
