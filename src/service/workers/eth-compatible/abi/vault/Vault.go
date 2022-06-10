// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_strategy\",\"internalType\":\"contractIStrategy\"},{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"_approvalDelay\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewStratCandidate\",\"inputs\":[{\"type\":\"address\",\"name\":\"implementation\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpgradeStrat\",\"inputs\":[{\"type\":\"address\",\"name\":\"implementation\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"allowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"approvalDelay\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"available\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balance\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"decimals\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"decreaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"deposit\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"depositAll\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"earn\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPricePerFullShare\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"inCaseTokensGetStuck\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"increaseAllowance\",\"inputs\":[{\"type\":\"address\",\"name\":\"spender\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"addedValue\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"proposeStrat\",\"inputs\":[{\"type\":\"address\",\"name\":\"_implementation\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"implementation\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"proposedTime\",\"internalType\":\"uint256\"}],\"name\":\"stratCandidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIStrategy\"}],\"name\":\"strategy\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"upgradeStrat\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIERC20\"}],\"name\":\"want\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdraw\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_shares\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"withdrawAll\",\"inputs\":[]}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vault *VaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vault *VaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vault.Contract.Allowance(&_Vault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Vault *VaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vault.Contract.Allowance(&_Vault.CallOpts, owner, spender)
}

// ApprovalDelay is a free data retrieval call binding the contract method 0xe2d1e75c.
//
// Solidity: function approvalDelay() view returns(uint256)
func (_Vault *VaultCaller) ApprovalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "approvalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovalDelay is a free data retrieval call binding the contract method 0xe2d1e75c.
//
// Solidity: function approvalDelay() view returns(uint256)
func (_Vault *VaultSession) ApprovalDelay() (*big.Int, error) {
	return _Vault.Contract.ApprovalDelay(&_Vault.CallOpts)
}

// ApprovalDelay is a free data retrieval call binding the contract method 0xe2d1e75c.
//
// Solidity: function approvalDelay() view returns(uint256)
func (_Vault *VaultCallerSession) ApprovalDelay() (*big.Int, error) {
	return _Vault.Contract.ApprovalDelay(&_Vault.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Vault *VaultCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Vault *VaultSession) Available() (*big.Int, error) {
	return _Vault.Contract.Available(&_Vault.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Vault *VaultCallerSession) Available() (*big.Int, error) {
	return _Vault.Contract.Available(&_Vault.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Vault *VaultCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Vault *VaultSession) Balance() (*big.Int, error) {
	return _Vault.Contract.Balance(&_Vault.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Vault *VaultCallerSession) Balance() (*big.Int, error) {
	return _Vault.Contract.Balance(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vault *VaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vault *VaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vault *VaultSession) Decimals() (uint8, error) {
	return _Vault.Contract.Decimals(&_Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vault *VaultCallerSession) Decimals() (uint8, error) {
	return _Vault.Contract.Decimals(&_Vault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Vault *VaultCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Vault *VaultSession) GetPricePerFullShare() (*big.Int, error) {
	return _Vault.Contract.GetPricePerFullShare(&_Vault.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Vault *VaultCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Vault.Contract.GetPricePerFullShare(&_Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vault *VaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vault *VaultSession) Name() (string, error) {
	return _Vault.Contract.Name(&_Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vault *VaultCallerSession) Name() (string, error) {
	return _Vault.Contract.Name(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultCallerSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// StratCandidate is a free data retrieval call binding the contract method 0x76dfabb8.
//
// Solidity: function stratCandidate() view returns(address implementation, uint256 proposedTime)
func (_Vault *VaultCaller) StratCandidate(opts *bind.CallOpts) (struct {
	Implementation common.Address
	ProposedTime   *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stratCandidate")

	outstruct := new(struct {
		Implementation common.Address
		ProposedTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Implementation = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProposedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StratCandidate is a free data retrieval call binding the contract method 0x76dfabb8.
//
// Solidity: function stratCandidate() view returns(address implementation, uint256 proposedTime)
func (_Vault *VaultSession) StratCandidate() (struct {
	Implementation common.Address
	ProposedTime   *big.Int
}, error) {
	return _Vault.Contract.StratCandidate(&_Vault.CallOpts)
}

// StratCandidate is a free data retrieval call binding the contract method 0x76dfabb8.
//
// Solidity: function stratCandidate() view returns(address implementation, uint256 proposedTime)
func (_Vault *VaultCallerSession) StratCandidate() (struct {
	Implementation common.Address
	ProposedTime   *big.Int
}, error) {
	return _Vault.Contract.StratCandidate(&_Vault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Vault *VaultCaller) Strategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "strategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Vault *VaultSession) Strategy() (common.Address, error) {
	return _Vault.Contract.Strategy(&_Vault.CallOpts)
}

// Strategy is a free data retrieval call binding the contract method 0xa8c62e76.
//
// Solidity: function strategy() view returns(address)
func (_Vault *VaultCallerSession) Strategy() (common.Address, error) {
	return _Vault.Contract.Strategy(&_Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vault *VaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vault *VaultSession) Symbol() (string, error) {
	return _Vault.Contract.Symbol(&_Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vault *VaultCallerSession) Symbol() (string, error) {
	return _Vault.Contract.Symbol(&_Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultSession) TotalSupply() (*big.Int, error) {
	return _Vault.Contract.TotalSupply(&_Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultCallerSession) TotalSupply() (*big.Int, error) {
	return _Vault.Contract.TotalSupply(&_Vault.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_Vault *VaultCaller) Want(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "want")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_Vault *VaultSession) Want() (common.Address, error) {
	return _Vault.Contract.Want(&_Vault.CallOpts)
}

// Want is a free data retrieval call binding the contract method 0x1f1fcd51.
//
// Solidity: function want() view returns(address)
func (_Vault *VaultCallerSession) Want() (common.Address, error) {
	return _Vault.Contract.Want(&_Vault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vault *VaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vault *VaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Approve(&_Vault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Vault *VaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Approve(&_Vault.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vault *VaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vault *VaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.DecreaseAllowance(&_Vault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Vault *VaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.DecreaseAllowance(&_Vault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vault *VaultTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, _amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Vault *VaultTransactor) DepositAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositAll")
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Vault *VaultSession) DepositAll() (*types.Transaction, error) {
	return _Vault.Contract.DepositAll(&_Vault.TransactOpts)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Vault *VaultTransactorSession) DepositAll() (*types.Transaction, error) {
	return _Vault.Contract.DepositAll(&_Vault.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Vault *VaultTransactor) Earn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "earn")
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Vault *VaultSession) Earn() (*types.Transaction, error) {
	return _Vault.Contract.Earn(&_Vault.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Vault *VaultTransactorSession) Earn() (*types.Transaction, error) {
	return _Vault.Contract.Earn(&_Vault.TransactOpts)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Vault *VaultTransactor) InCaseTokensGetStuck(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "inCaseTokensGetStuck", _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Vault *VaultSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.InCaseTokensGetStuck(&_Vault.TransactOpts, _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_Vault *VaultTransactorSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.InCaseTokensGetStuck(&_Vault.TransactOpts, _token)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vault *VaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vault *VaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.IncreaseAllowance(&_Vault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Vault *VaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.IncreaseAllowance(&_Vault.TransactOpts, spender, addedValue)
}

// ProposeStrat is a paid mutator transaction binding the contract method 0x5b12ff9b.
//
// Solidity: function proposeStrat(address _implementation) returns()
func (_Vault *VaultTransactor) ProposeStrat(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "proposeStrat", _implementation)
}

// ProposeStrat is a paid mutator transaction binding the contract method 0x5b12ff9b.
//
// Solidity: function proposeStrat(address _implementation) returns()
func (_Vault *VaultSession) ProposeStrat(_implementation common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ProposeStrat(&_Vault.TransactOpts, _implementation)
}

// ProposeStrat is a paid mutator transaction binding the contract method 0x5b12ff9b.
//
// Solidity: function proposeStrat(address _implementation) returns()
func (_Vault *VaultTransactorSession) ProposeStrat(_implementation common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ProposeStrat(&_Vault.TransactOpts, _implementation)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vault.Contract.RenounceOwnership(&_Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vault.Contract.RenounceOwnership(&_Vault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Vault *VaultSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Vault *VaultTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Vault *VaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Vault *VaultSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferFrom(&_Vault.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Vault *VaultTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferFrom(&_Vault.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vault.Contract.TransferOwnership(&_Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vault.Contract.TransferOwnership(&_Vault.TransactOpts, newOwner)
}

// UpgradeStrat is a paid mutator transaction binding the contract method 0xe6685244.
//
// Solidity: function upgradeStrat() returns()
func (_Vault *VaultTransactor) UpgradeStrat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "upgradeStrat")
}

// UpgradeStrat is a paid mutator transaction binding the contract method 0xe6685244.
//
// Solidity: function upgradeStrat() returns()
func (_Vault *VaultSession) UpgradeStrat() (*types.Transaction, error) {
	return _Vault.Contract.UpgradeStrat(&_Vault.TransactOpts)
}

// UpgradeStrat is a paid mutator transaction binding the contract method 0xe6685244.
//
// Solidity: function upgradeStrat() returns()
func (_Vault *VaultTransactorSession) UpgradeStrat() (*types.Transaction, error) {
	return _Vault.Contract.UpgradeStrat(&_Vault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Vault *VaultSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Vault *VaultTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Vault *VaultTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Vault *VaultSession) WithdrawAll() (*types.Transaction, error) {
	return _Vault.Contract.WithdrawAll(&_Vault.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Vault *VaultTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Vault.Contract.WithdrawAll(&_Vault.TransactOpts)
}

// VaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vault contract.
type VaultApprovalIterator struct {
	Event *VaultApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultApproval represents a Approval event raised by the Vault contract.
type VaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vault *VaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultApprovalIterator{contract: _Vault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vault *VaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultApproval)
				if err := _Vault.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vault *VaultFilterer) ParseApproval(log types.Log) (*VaultApproval, error) {
	event := new(VaultApproval)
	if err := _Vault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultNewStratCandidateIterator is returned from FilterNewStratCandidate and is used to iterate over the raw logs and unpacked data for NewStratCandidate events raised by the Vault contract.
type VaultNewStratCandidateIterator struct {
	Event *VaultNewStratCandidate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultNewStratCandidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNewStratCandidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultNewStratCandidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultNewStratCandidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNewStratCandidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNewStratCandidate represents a NewStratCandidate event raised by the Vault contract.
type VaultNewStratCandidate struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewStratCandidate is a free log retrieval operation binding the contract event 0x1aae2ec5647db56da2d513de40528ba3565c6057525637050660c4323bbac7df.
//
// Solidity: event NewStratCandidate(address implementation)
func (_Vault *VaultFilterer) FilterNewStratCandidate(opts *bind.FilterOpts) (*VaultNewStratCandidateIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NewStratCandidate")
	if err != nil {
		return nil, err
	}
	return &VaultNewStratCandidateIterator{contract: _Vault.contract, event: "NewStratCandidate", logs: logs, sub: sub}, nil
}

// WatchNewStratCandidate is a free log subscription operation binding the contract event 0x1aae2ec5647db56da2d513de40528ba3565c6057525637050660c4323bbac7df.
//
// Solidity: event NewStratCandidate(address implementation)
func (_Vault *VaultFilterer) WatchNewStratCandidate(opts *bind.WatchOpts, sink chan<- *VaultNewStratCandidate) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NewStratCandidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNewStratCandidate)
				if err := _Vault.contract.UnpackLog(event, "NewStratCandidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewStratCandidate is a log parse operation binding the contract event 0x1aae2ec5647db56da2d513de40528ba3565c6057525637050660c4323bbac7df.
//
// Solidity: event NewStratCandidate(address implementation)
func (_Vault *VaultFilterer) ParseNewStratCandidate(log types.Log) (*VaultNewStratCandidate, error) {
	event := new(VaultNewStratCandidate)
	if err := _Vault.contract.UnpackLog(event, "NewStratCandidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vault contract.
type VaultOwnershipTransferredIterator struct {
	Event *VaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultOwnershipTransferred represents a OwnershipTransferred event raised by the Vault contract.
type VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VaultOwnershipTransferredIterator{contract: _Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultOwnershipTransferred)
				if err := _Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) ParseOwnershipTransferred(log types.Log) (*VaultOwnershipTransferred, error) {
	event := new(VaultOwnershipTransferred)
	if err := _Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vault contract.
type VaultTransferIterator struct {
	Event *VaultTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTransfer represents a Transfer event raised by the Vault contract.
type VaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vault *VaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultTransferIterator{contract: _Vault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vault *VaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTransfer)
				if err := _Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vault *VaultFilterer) ParseTransfer(log types.Log) (*VaultTransfer, error) {
	event := new(VaultTransfer)
	if err := _Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpgradeStratIterator is returned from FilterUpgradeStrat and is used to iterate over the raw logs and unpacked data for UpgradeStrat events raised by the Vault contract.
type VaultUpgradeStratIterator struct {
	Event *VaultUpgradeStrat // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpgradeStratIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpgradeStrat)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpgradeStrat)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpgradeStratIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpgradeStratIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpgradeStrat represents a UpgradeStrat event raised by the Vault contract.
type VaultUpgradeStrat struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgradeStrat is a free log retrieval operation binding the contract event 0x7f37d440e85aba7fbf641c4bda5ca4ef669a80bffaacde2aa8d9feb1b048c82c.
//
// Solidity: event UpgradeStrat(address implementation)
func (_Vault *VaultFilterer) FilterUpgradeStrat(opts *bind.FilterOpts) (*VaultUpgradeStratIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpgradeStrat")
	if err != nil {
		return nil, err
	}
	return &VaultUpgradeStratIterator{contract: _Vault.contract, event: "UpgradeStrat", logs: logs, sub: sub}, nil
}

// WatchUpgradeStrat is a free log subscription operation binding the contract event 0x7f37d440e85aba7fbf641c4bda5ca4ef669a80bffaacde2aa8d9feb1b048c82c.
//
// Solidity: event UpgradeStrat(address implementation)
func (_Vault *VaultFilterer) WatchUpgradeStrat(opts *bind.WatchOpts, sink chan<- *VaultUpgradeStrat) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpgradeStrat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpgradeStrat)
				if err := _Vault.contract.UnpackLog(event, "UpgradeStrat", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgradeStrat is a log parse operation binding the contract event 0x7f37d440e85aba7fbf641c4bda5ca4ef669a80bffaacde2aa8d9feb1b048c82c.
//
// Solidity: event UpgradeStrat(address implementation)
func (_Vault *VaultFilterer) ParseUpgradeStrat(log types.Log) (*VaultUpgradeStrat, error) {
	event := new(VaultUpgradeStrat)
	if err := _Vault.contract.UnpackLog(event, "UpgradeStrat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
