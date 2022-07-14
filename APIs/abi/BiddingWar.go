// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeExtension\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GameEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GameNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GameNotPlaying\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBid\",\"type\":\"uint256\"}],\"name\":\"LowerBid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"NotElapsedToBid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"NotElapsedToEnd\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"EndGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"StartGame\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBidEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultTimeExtension\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPlaying\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTimeElapsedToBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTimeElapsedToEnd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidEpoch\",\"type\":\"uint256\"}],\"name\":\"setDefaultBidEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeExtension\",\"type\":\"uint256\"}],\"name\":\"setDefaultTimeExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"playingPeriod\",\"type\":\"uint256\"}],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Store *StoreCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Store *StoreSession) Commission() (*big.Int, error) {
	return _Store.Contract.Commission(&_Store.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Store *StoreCallerSession) Commission() (*big.Int, error) {
	return _Store.Contract.Commission(&_Store.CallOpts)
}

// DefaultBidEpoch is a free data retrieval call binding the contract method 0xea13b9f9.
//
// Solidity: function defaultBidEpoch() view returns(uint256)
func (_Store *StoreCaller) DefaultBidEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "defaultBidEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultBidEpoch is a free data retrieval call binding the contract method 0xea13b9f9.
//
// Solidity: function defaultBidEpoch() view returns(uint256)
func (_Store *StoreSession) DefaultBidEpoch() (*big.Int, error) {
	return _Store.Contract.DefaultBidEpoch(&_Store.CallOpts)
}

// DefaultBidEpoch is a free data retrieval call binding the contract method 0xea13b9f9.
//
// Solidity: function defaultBidEpoch() view returns(uint256)
func (_Store *StoreCallerSession) DefaultBidEpoch() (*big.Int, error) {
	return _Store.Contract.DefaultBidEpoch(&_Store.CallOpts)
}

// DefaultTimeExtension is a free data retrieval call binding the contract method 0xe9eace91.
//
// Solidity: function defaultTimeExtension() view returns(uint256)
func (_Store *StoreCaller) DefaultTimeExtension(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "defaultTimeExtension")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultTimeExtension is a free data retrieval call binding the contract method 0xe9eace91.
//
// Solidity: function defaultTimeExtension() view returns(uint256)
func (_Store *StoreSession) DefaultTimeExtension() (*big.Int, error) {
	return _Store.Contract.DefaultTimeExtension(&_Store.CallOpts)
}

// DefaultTimeExtension is a free data retrieval call binding the contract method 0xe9eace91.
//
// Solidity: function defaultTimeExtension() view returns(uint256)
func (_Store *StoreCallerSession) DefaultTimeExtension() (*big.Int, error) {
	return _Store.Contract.DefaultTimeExtension(&_Store.CallOpts)
}

// GameInfo is a free data retrieval call binding the contract method 0xdb73bfce.
//
// Solidity: function gameInfo() view returns(uint256, uint256, uint256, uint256, uint256, address, uint256, bool, uint256)
func (_Store *StoreCaller) GameInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "gameInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(bool)).(*bool)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, err

}

// GameInfo is a free data retrieval call binding the contract method 0xdb73bfce.
//
// Solidity: function gameInfo() view returns(uint256, uint256, uint256, uint256, uint256, address, uint256, bool, uint256)
func (_Store *StoreSession) GameInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, bool, *big.Int, error) {
	return _Store.Contract.GameInfo(&_Store.CallOpts)
}

// GameInfo is a free data retrieval call binding the contract method 0xdb73bfce.
//
// Solidity: function gameInfo() view returns(uint256, uint256, uint256, uint256, uint256, address, uint256, bool, uint256)
func (_Store *StoreCallerSession) GameInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, bool, *big.Int, error) {
	return _Store.Contract.GameInfo(&_Store.CallOpts)
}

// IsPlaying is a free data retrieval call binding the contract method 0x55113fd3.
//
// Solidity: function isPlaying() view returns(bool)
func (_Store *StoreCaller) IsPlaying(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isPlaying")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPlaying is a free data retrieval call binding the contract method 0x55113fd3.
//
// Solidity: function isPlaying() view returns(bool)
func (_Store *StoreSession) IsPlaying() (bool, error) {
	return _Store.Contract.IsPlaying(&_Store.CallOpts)
}

// IsPlaying is a free data retrieval call binding the contract method 0x55113fd3.
//
// Solidity: function isPlaying() view returns(bool)
func (_Store *StoreCallerSession) IsPlaying() (bool, error) {
	return _Store.Contract.IsPlaying(&_Store.CallOpts)
}

// IsTimeElapsedToBid is a free data retrieval call binding the contract method 0x6e31ed14.
//
// Solidity: function isTimeElapsedToBid() view returns(bool)
func (_Store *StoreCaller) IsTimeElapsedToBid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isTimeElapsedToBid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimeElapsedToBid is a free data retrieval call binding the contract method 0x6e31ed14.
//
// Solidity: function isTimeElapsedToBid() view returns(bool)
func (_Store *StoreSession) IsTimeElapsedToBid() (bool, error) {
	return _Store.Contract.IsTimeElapsedToBid(&_Store.CallOpts)
}

// IsTimeElapsedToBid is a free data retrieval call binding the contract method 0x6e31ed14.
//
// Solidity: function isTimeElapsedToBid() view returns(bool)
func (_Store *StoreCallerSession) IsTimeElapsedToBid() (bool, error) {
	return _Store.Contract.IsTimeElapsedToBid(&_Store.CallOpts)
}

// IsTimeElapsedToEnd is a free data retrieval call binding the contract method 0xd804d6af.
//
// Solidity: function isTimeElapsedToEnd() view returns(bool)
func (_Store *StoreCaller) IsTimeElapsedToEnd(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isTimeElapsedToEnd")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimeElapsedToEnd is a free data retrieval call binding the contract method 0xd804d6af.
//
// Solidity: function isTimeElapsedToEnd() view returns(bool)
func (_Store *StoreSession) IsTimeElapsedToEnd() (bool, error) {
	return _Store.Contract.IsTimeElapsedToEnd(&_Store.CallOpts)
}

// IsTimeElapsedToEnd is a free data retrieval call binding the contract method 0xd804d6af.
//
// Solidity: function isTimeElapsedToEnd() view returns(bool)
func (_Store *StoreCallerSession) IsTimeElapsedToEnd() (bool, error) {
	return _Store.Contract.IsTimeElapsedToEnd(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Store *StoreTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Store *StoreSession) Bid() (*types.Transaction, error) {
	return _Store.Contract.Bid(&_Store.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Store *StoreTransactorSession) Bid() (*types.Transaction, error) {
	return _Store.Contract.Bid(&_Store.TransactOpts)
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Store *StoreTransactor) EndGame(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "endGame")
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Store *StoreSession) EndGame() (*types.Transaction, error) {
	return _Store.Contract.EndGame(&_Store.TransactOpts)
}

// EndGame is a paid mutator transaction binding the contract method 0x6cbc2ded.
//
// Solidity: function endGame() returns()
func (_Store *StoreTransactorSession) EndGame() (*types.Transaction, error) {
	return _Store.Contract.EndGame(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// SetDefaultBidEpoch is a paid mutator transaction binding the contract method 0x0647e025.
//
// Solidity: function setDefaultBidEpoch(uint256 bidEpoch) returns()
func (_Store *StoreTransactor) SetDefaultBidEpoch(opts *bind.TransactOpts, bidEpoch *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setDefaultBidEpoch", bidEpoch)
}

// SetDefaultBidEpoch is a paid mutator transaction binding the contract method 0x0647e025.
//
// Solidity: function setDefaultBidEpoch(uint256 bidEpoch) returns()
func (_Store *StoreSession) SetDefaultBidEpoch(bidEpoch *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetDefaultBidEpoch(&_Store.TransactOpts, bidEpoch)
}

// SetDefaultBidEpoch is a paid mutator transaction binding the contract method 0x0647e025.
//
// Solidity: function setDefaultBidEpoch(uint256 bidEpoch) returns()
func (_Store *StoreTransactorSession) SetDefaultBidEpoch(bidEpoch *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetDefaultBidEpoch(&_Store.TransactOpts, bidEpoch)
}

// SetDefaultTimeExtension is a paid mutator transaction binding the contract method 0xa1575c7f.
//
// Solidity: function setDefaultTimeExtension(uint256 timeExtension) returns()
func (_Store *StoreTransactor) SetDefaultTimeExtension(opts *bind.TransactOpts, timeExtension *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setDefaultTimeExtension", timeExtension)
}

// SetDefaultTimeExtension is a paid mutator transaction binding the contract method 0xa1575c7f.
//
// Solidity: function setDefaultTimeExtension(uint256 timeExtension) returns()
func (_Store *StoreSession) SetDefaultTimeExtension(timeExtension *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetDefaultTimeExtension(&_Store.TransactOpts, timeExtension)
}

// SetDefaultTimeExtension is a paid mutator transaction binding the contract method 0xa1575c7f.
//
// Solidity: function setDefaultTimeExtension(uint256 timeExtension) returns()
func (_Store *StoreTransactorSession) SetDefaultTimeExtension(timeExtension *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetDefaultTimeExtension(&_Store.TransactOpts, timeExtension)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 playingPeriod) returns()
func (_Store *StoreTransactor) StartGame(opts *bind.TransactOpts, playingPeriod *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "startGame", playingPeriod)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 playingPeriod) returns()
func (_Store *StoreSession) StartGame(playingPeriod *big.Int) (*types.Transaction, error) {
	return _Store.Contract.StartGame(&_Store.TransactOpts, playingPeriod)
}

// StartGame is a paid mutator transaction binding the contract method 0xe5ed1d59.
//
// Solidity: function startGame(uint256 playingPeriod) returns()
func (_Store *StoreTransactorSession) StartGame(playingPeriod *big.Int) (*types.Transaction, error) {
	return _Store.Contract.StartGame(&_Store.TransactOpts, playingPeriod)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address sender) returns()
func (_Store *StoreTransactor) Withdraw(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "withdraw", sender)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address sender) returns()
func (_Store *StoreSession) Withdraw(sender common.Address) (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts, sender)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address sender) returns()
func (_Store *StoreTransactorSession) Withdraw(sender common.Address) (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts, sender)
}

// StoreBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Store contract.
type StoreBidIterator struct {
	Event *StoreBid // Event containing the contract specifics and raw log

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
func (it *StoreBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBid)
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
		it.Event = new(StoreBid)
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
func (it *StoreBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBid represents a Bid event raised by the Store contract.
type StoreBid struct {
	Bidder   common.Address
	BidPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 bidPrice)
func (_Store *StoreFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*StoreBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &StoreBidIterator{contract: _Store.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 bidPrice)
func (_Store *StoreFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *StoreBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBid)
				if err := _Store.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 bidPrice)
func (_Store *StoreFilterer) ParseBid(log types.Log) (*StoreBid, error) {
	event := new(StoreBid)
	if err := _Store.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreEndGameIterator is returned from FilterEndGame and is used to iterate over the raw logs and unpacked data for EndGame events raised by the Store contract.
type StoreEndGameIterator struct {
	Event *StoreEndGame // Event containing the contract specifics and raw log

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
func (it *StoreEndGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreEndGame)
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
		it.Event = new(StoreEndGame)
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
func (it *StoreEndGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreEndGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreEndGame represents a EndGame event raised by the Store contract.
type StoreEndGame struct {
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEndGame is a free log retrieval operation binding the contract event 0xef49d593dd1d07d53ec845a9c6b11933e44043587cb2f02cb256be5788aea420.
//
// Solidity: event EndGame(address winner, uint256 prizeAmount)
func (_Store *StoreFilterer) FilterEndGame(opts *bind.FilterOpts) (*StoreEndGameIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "EndGame")
	if err != nil {
		return nil, err
	}
	return &StoreEndGameIterator{contract: _Store.contract, event: "EndGame", logs: logs, sub: sub}, nil
}

// WatchEndGame is a free log subscription operation binding the contract event 0xef49d593dd1d07d53ec845a9c6b11933e44043587cb2f02cb256be5788aea420.
//
// Solidity: event EndGame(address winner, uint256 prizeAmount)
func (_Store *StoreFilterer) WatchEndGame(opts *bind.WatchOpts, sink chan<- *StoreEndGame) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "EndGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreEndGame)
				if err := _Store.contract.UnpackLog(event, "EndGame", log); err != nil {
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

// ParseEndGame is a log parse operation binding the contract event 0xef49d593dd1d07d53ec845a9c6b11933e44043587cb2f02cb256be5788aea420.
//
// Solidity: event EndGame(address winner, uint256 prizeAmount)
func (_Store *StoreFilterer) ParseEndGame(log types.Log) (*StoreEndGame, error) {
	event := new(StoreEndGame)
	if err := _Store.contract.UnpackLog(event, "EndGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
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
		it.Event = new(StoreOwnershipTransferred)
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
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Store *StoreFilterer) ParseOwnershipTransferred(log types.Log) (*StoreOwnershipTransferred, error) {
	event := new(StoreOwnershipTransferred)
	if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStartGameIterator is returned from FilterStartGame and is used to iterate over the raw logs and unpacked data for StartGame events raised by the Store contract.
type StoreStartGameIterator struct {
	Event *StoreStartGame // Event containing the contract specifics and raw log

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
func (it *StoreStartGameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStartGame)
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
		it.Event = new(StoreStartGame)
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
func (it *StoreStartGameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStartGameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStartGame represents a StartGame event raised by the Store contract.
type StoreStartGame struct {
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartGame is a free log retrieval operation binding the contract event 0x12333b322723ec26360eb708b341801ecca576387b55caa77d6e408107d2c70f.
//
// Solidity: event StartGame(uint256 startedAt)
func (_Store *StoreFilterer) FilterStartGame(opts *bind.FilterOpts) (*StoreStartGameIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return &StoreStartGameIterator{contract: _Store.contract, event: "StartGame", logs: logs, sub: sub}, nil
}

// WatchStartGame is a free log subscription operation binding the contract event 0x12333b322723ec26360eb708b341801ecca576387b55caa77d6e408107d2c70f.
//
// Solidity: event StartGame(uint256 startedAt)
func (_Store *StoreFilterer) WatchStartGame(opts *bind.WatchOpts, sink chan<- *StoreStartGame) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StartGame")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStartGame)
				if err := _Store.contract.UnpackLog(event, "StartGame", log); err != nil {
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

// ParseStartGame is a log parse operation binding the contract event 0x12333b322723ec26360eb708b341801ecca576387b55caa77d6e408107d2c70f.
//
// Solidity: event StartGame(uint256 startedAt)
func (_Store *StoreFilterer) ParseStartGame(log types.Log) (*StoreStartGame, error) {
	event := new(StoreStartGame)
	if err := _Store.contract.UnpackLog(event, "StartGame", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address sender, uint256 amount)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts) (*StoreTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address sender, uint256 amount)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address sender, uint256 amount)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
