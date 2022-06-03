package eth

import (
	"encoding/binary"
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
		utils.StringToBytes8LeftPad(destinationChainID),
	)
	if err != nil {
		return common.Hash{}, err
	}
	messageHash := crypto.Keccak256Hash(bytes)
	return messageHash, nil
}

func (w *Erc20Worker) CreateSignature(messageHash common.Hash, destinationChainID string) (string, error) {
	privKey, err := utils.GetPrivateKey(w.config)
	if err != nil {
		return "", err
	}
	signature, er := crypto.Sign(messageHash.Bytes(), privKey)
	if er != nil {
		return "", er
	}
	if destinationChainID == w.config.DestinationChainID {
		return hexutil.Encode(signature), nil
	} else {
		if w.chainID < 110 {
			signature[64] = byte(w.chainID)*2 + 35 + signature[64]
			return hexutil.Encode(signature), nil
		}
		result := make([]byte, 66)
		encodedRecId := uint32(w.chainID)*2 + 35 + uint32(signature[64])
		recIdBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(recIdBytes, encodedRecId)
		copy(result[0:64], signature[0:64])
		// only 2 first bytes contains non-zero value because chainId is byte, so it is less than 256
		result[64] = recIdBytes[1]
		result[65] = recIdBytes[0]
		return hexutil.Encode(result), nil
	}

}
