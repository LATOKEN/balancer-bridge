package app

import "github.com/ethereum/go-ethereum/common"

// ErrorMsg ...
type ErrorMsg struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

type SigAndHash struct {
	Hash      common.Hash `json:"hash"`
	Signature string      `json:"signature"`
}

func createNewError(err, msg string) ErrorMsg {
	return ErrorMsg{
		Type:  err,
		Error: msg,
	}
}
