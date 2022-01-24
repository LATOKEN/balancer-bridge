package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.nekotal.tech/lachain/crosschain/bridge-backend-service/src/service/workers/utils"
)

func (w *Erc20Worker) CreateMessageHash(amount, recipientAddress, originChainID string) (common.Hash, error) {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	addressTy, _ := abi.NewType("address", "address", nil)
	bytesTy, _ := abi.NewType("bytes", "bytes", nil)

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
		{
			Type: uint256Ty,
		},
	}
	value, _ := new(big.Int).SetString(amount, 10)
	bytes, err := arguments.Pack(
		value,
		common.HexToAddress(recipientAddress),
		[]byte(originChainID),
		big.NewInt(w.signatureNonce),
	)
	if err != nil {
		return common.Hash{}, err
	}

	//increase signature nonce so no signature is same
	w.signatureNonce++

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

	return hexutil.Encode(signature), nil
}
