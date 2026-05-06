// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBlacklist\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blacklist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimReward\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPendingReward\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_stakeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeMinAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeMaxAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeBlacklist\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRewardRate\",\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeMaxAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeMinAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateStakeLimits\",\"inputs\":[{\"name\":\"_min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userStakes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlacklistAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlacklistRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardClaimed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeLimitsUpdated\",\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unstaked\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Abi *AbiCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Abi *AbiSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Abi.Contract.UPGRADEINTERFACEVERSION(&_Abi.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Abi *AbiCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Abi.Contract.UPGRADEINTERFACEVERSION(&_Abi.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Abi *AbiCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Abi *AbiSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Abi.Contract.Blacklist(&_Abi.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Abi *AbiCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Abi.Contract.Blacklist(&_Abi.CallOpts, arg0)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _user) view returns(uint256)
func (_Abi *AbiCaller) GetPendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getPendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _user) view returns(uint256)
func (_Abi *AbiSession) GetPendingReward(_user common.Address) (*big.Int, error) {
	return _Abi.Contract.GetPendingReward(&_Abi.CallOpts, _user)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _user) view returns(uint256)
func (_Abi *AbiCallerSession) GetPendingReward(_user common.Address) (*big.Int, error) {
	return _Abi.Contract.GetPendingReward(&_Abi.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Abi *AbiCallerSession) Paused() (bool, error) {
	return _Abi.Contract.Paused(&_Abi.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiSession) ProxiableUUID() ([32]byte, error) {
	return _Abi.Contract.ProxiableUUID(&_Abi.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Abi *AbiCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Abi.Contract.ProxiableUUID(&_Abi.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Abi *AbiCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Abi *AbiSession) RewardRate() (*big.Int, error) {
	return _Abi.Contract.RewardRate(&_Abi.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Abi *AbiCallerSession) RewardRate() (*big.Int, error) {
	return _Abi.Contract.RewardRate(&_Abi.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Abi *AbiCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Abi *AbiSession) RewardToken() (common.Address, error) {
	return _Abi.Contract.RewardToken(&_Abi.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Abi *AbiCallerSession) RewardToken() (common.Address, error) {
	return _Abi.Contract.RewardToken(&_Abi.CallOpts)
}

// StakeMaxAmount is a free data retrieval call binding the contract method 0xf7d7a686.
//
// Solidity: function stakeMaxAmount() view returns(uint256)
func (_Abi *AbiCaller) StakeMaxAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "stakeMaxAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeMaxAmount is a free data retrieval call binding the contract method 0xf7d7a686.
//
// Solidity: function stakeMaxAmount() view returns(uint256)
func (_Abi *AbiSession) StakeMaxAmount() (*big.Int, error) {
	return _Abi.Contract.StakeMaxAmount(&_Abi.CallOpts)
}

// StakeMaxAmount is a free data retrieval call binding the contract method 0xf7d7a686.
//
// Solidity: function stakeMaxAmount() view returns(uint256)
func (_Abi *AbiCallerSession) StakeMaxAmount() (*big.Int, error) {
	return _Abi.Contract.StakeMaxAmount(&_Abi.CallOpts)
}

// StakeMinAmount is a free data retrieval call binding the contract method 0x72032061.
//
// Solidity: function stakeMinAmount() view returns(uint256)
func (_Abi *AbiCaller) StakeMinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "stakeMinAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeMinAmount is a free data retrieval call binding the contract method 0x72032061.
//
// Solidity: function stakeMinAmount() view returns(uint256)
func (_Abi *AbiSession) StakeMinAmount() (*big.Int, error) {
	return _Abi.Contract.StakeMinAmount(&_Abi.CallOpts)
}

// StakeMinAmount is a free data retrieval call binding the contract method 0x72032061.
//
// Solidity: function stakeMinAmount() view returns(uint256)
func (_Abi *AbiCallerSession) StakeMinAmount() (*big.Int, error) {
	return _Abi.Contract.StakeMinAmount(&_Abi.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Abi *AbiCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Abi *AbiSession) StakeToken() (common.Address, error) {
	return _Abi.Contract.StakeToken(&_Abi.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Abi *AbiCallerSession) StakeToken() (common.Address, error) {
	return _Abi.Contract.StakeToken(&_Abi.CallOpts)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 lastUpdate)
func (_Abi *AbiCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	LastUpdate *big.Int
}, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "userStakes", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
		LastUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 lastUpdate)
func (_Abi *AbiSession) UserStakes(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	LastUpdate *big.Int
}, error) {
	return _Abi.Contract.UserStakes(&_Abi.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 amount, uint256 rewardDebt, uint256 lastUpdate)
func (_Abi *AbiCallerSession) UserStakes(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	LastUpdate *big.Int
}, error) {
	return _Abi.Contract.UserStakes(&_Abi.CallOpts, arg0)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address _account) returns()
func (_Abi *AbiTransactor) AddBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "addBlacklist", _account)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address _account) returns()
func (_Abi *AbiSession) AddBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.AddBlacklist(&_Abi.TransactOpts, _account)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address _account) returns()
func (_Abi *AbiTransactorSession) AddBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.AddBlacklist(&_Abi.TransactOpts, _account)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Abi *AbiTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Abi *AbiSession) ClaimReward() (*types.Transaction, error) {
	return _Abi.Contract.ClaimReward(&_Abi.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Abi *AbiTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Abi.Contract.ClaimReward(&_Abi.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _stakeToken, address _rewardToken, uint256 _rewardRate, uint256 _stakeMinAmount, uint256 _stakeMaxAmount) returns()
func (_Abi *AbiTransactor) Initialize(opts *bind.TransactOpts, _stakeToken common.Address, _rewardToken common.Address, _rewardRate *big.Int, _stakeMinAmount *big.Int, _stakeMaxAmount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "initialize", _stakeToken, _rewardToken, _rewardRate, _stakeMinAmount, _stakeMaxAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _stakeToken, address _rewardToken, uint256 _rewardRate, uint256 _stakeMinAmount, uint256 _stakeMaxAmount) returns()
func (_Abi *AbiSession) Initialize(_stakeToken common.Address, _rewardToken common.Address, _rewardRate *big.Int, _stakeMinAmount *big.Int, _stakeMaxAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, _stakeToken, _rewardToken, _rewardRate, _stakeMinAmount, _stakeMaxAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd13f90b4.
//
// Solidity: function initialize(address _stakeToken, address _rewardToken, uint256 _rewardRate, uint256 _stakeMinAmount, uint256 _stakeMaxAmount) returns()
func (_Abi *AbiTransactorSession) Initialize(_stakeToken common.Address, _rewardToken common.Address, _rewardRate *big.Int, _stakeMinAmount *big.Int, _stakeMaxAmount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Initialize(&_Abi.TransactOpts, _stakeToken, _rewardToken, _rewardRate, _stakeMinAmount, _stakeMaxAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Abi *AbiTransactorSession) Pause() (*types.Transaction, error) {
	return _Abi.Contract.Pause(&_Abi.TransactOpts)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address _account) returns()
func (_Abi *AbiTransactor) RemoveBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "removeBlacklist", _account)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address _account) returns()
func (_Abi *AbiSession) RemoveBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RemoveBlacklist(&_Abi.TransactOpts, _account)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address _account) returns()
func (_Abi *AbiTransactorSession) RemoveBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Abi.Contract.RemoveBlacklist(&_Abi.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rate) returns()
func (_Abi *AbiTransactor) SetRewardRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setRewardRate", _rate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rate) returns()
func (_Abi *AbiSession) SetRewardRate(_rate *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetRewardRate(&_Abi.TransactOpts, _rate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rate) returns()
func (_Abi *AbiTransactorSession) SetRewardRate(_rate *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SetRewardRate(&_Abi.TransactOpts, _rate)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Abi *AbiTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Abi *AbiSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Stake(&_Abi.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Abi *AbiTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Stake(&_Abi.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Abi *AbiTransactorSession) Unpause() (*types.Transaction, error) {
	return _Abi.Contract.Unpause(&_Abi.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Abi *AbiTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Abi *AbiSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Unstake(&_Abi.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Abi *AbiTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Unstake(&_Abi.TransactOpts, _amount)
}

// UpdateStakeLimits is a paid mutator transaction binding the contract method 0x52c8c18c.
//
// Solidity: function updateStakeLimits(uint256 _min, uint256 _max) returns()
func (_Abi *AbiTransactor) UpdateStakeLimits(opts *bind.TransactOpts, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "updateStakeLimits", _min, _max)
}

// UpdateStakeLimits is a paid mutator transaction binding the contract method 0x52c8c18c.
//
// Solidity: function updateStakeLimits(uint256 _min, uint256 _max) returns()
func (_Abi *AbiSession) UpdateStakeLimits(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.UpdateStakeLimits(&_Abi.TransactOpts, _min, _max)
}

// UpdateStakeLimits is a paid mutator transaction binding the contract method 0x52c8c18c.
//
// Solidity: function updateStakeLimits(uint256 _min, uint256 _max) returns()
func (_Abi *AbiTransactorSession) UpdateStakeLimits(_min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.UpdateStakeLimits(&_Abi.TransactOpts, _min, _max)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeToAndCall(&_Abi.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Abi *AbiTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.UpgradeToAndCall(&_Abi.TransactOpts, newImplementation, data)
}

// AbiBlacklistAddedIterator is returned from FilterBlacklistAdded and is used to iterate over the raw logs and unpacked data for BlacklistAdded events raised by the Abi contract.
type AbiBlacklistAddedIterator struct {
	Event *AbiBlacklistAdded // Event containing the contract specifics and raw log

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
func (it *AbiBlacklistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiBlacklistAdded)
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
		it.Event = new(AbiBlacklistAdded)
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
func (it *AbiBlacklistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiBlacklistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiBlacklistAdded represents a BlacklistAdded event raised by the Abi contract.
type AbiBlacklistAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistAdded is a free log retrieval operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account)
func (_Abi *AbiFilterer) FilterBlacklistAdded(opts *bind.FilterOpts, account []common.Address) (*AbiBlacklistAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "BlacklistAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &AbiBlacklistAddedIterator{contract: _Abi.contract, event: "BlacklistAdded", logs: logs, sub: sub}, nil
}

// WatchBlacklistAdded is a free log subscription operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account)
func (_Abi *AbiFilterer) WatchBlacklistAdded(opts *bind.WatchOpts, sink chan<- *AbiBlacklistAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "BlacklistAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiBlacklistAdded)
				if err := _Abi.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
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

// ParseBlacklistAdded is a log parse operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account)
func (_Abi *AbiFilterer) ParseBlacklistAdded(log types.Log) (*AbiBlacklistAdded, error) {
	event := new(AbiBlacklistAdded)
	if err := _Abi.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiBlacklistRemovedIterator is returned from FilterBlacklistRemoved and is used to iterate over the raw logs and unpacked data for BlacklistRemoved events raised by the Abi contract.
type AbiBlacklistRemovedIterator struct {
	Event *AbiBlacklistRemoved // Event containing the contract specifics and raw log

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
func (it *AbiBlacklistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiBlacklistRemoved)
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
		it.Event = new(AbiBlacklistRemoved)
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
func (it *AbiBlacklistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiBlacklistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiBlacklistRemoved represents a BlacklistRemoved event raised by the Abi contract.
type AbiBlacklistRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistRemoved is a free log retrieval operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account)
func (_Abi *AbiFilterer) FilterBlacklistRemoved(opts *bind.FilterOpts, account []common.Address) (*AbiBlacklistRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "BlacklistRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AbiBlacklistRemovedIterator{contract: _Abi.contract, event: "BlacklistRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistRemoved is a free log subscription operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account)
func (_Abi *AbiFilterer) WatchBlacklistRemoved(opts *bind.WatchOpts, sink chan<- *AbiBlacklistRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "BlacklistRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiBlacklistRemoved)
				if err := _Abi.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
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

// ParseBlacklistRemoved is a log parse operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account)
func (_Abi *AbiFilterer) ParseBlacklistRemoved(log types.Log) (*AbiBlacklistRemoved, error) {
	event := new(AbiBlacklistRemoved)
	if err := _Abi.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Abi contract.
type AbiInitializedIterator struct {
	Event *AbiInitialized // Event containing the contract specifics and raw log

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
func (it *AbiInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiInitialized)
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
		it.Event = new(AbiInitialized)
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
func (it *AbiInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiInitialized represents a Initialized event raised by the Abi contract.
type AbiInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) FilterInitialized(opts *bind.FilterOpts) (*AbiInitializedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AbiInitializedIterator{contract: _Abi.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AbiInitialized) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiInitialized)
				if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Abi *AbiFilterer) ParseInitialized(log types.Log) (*AbiInitialized, error) {
	event := new(AbiInitialized)
	if err := _Abi.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
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
		it.Event = new(AbiOwnershipTransferred)
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
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Abi contract.
type AbiPausedIterator struct {
	Event *AbiPaused // Event containing the contract specifics and raw log

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
func (it *AbiPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiPaused)
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
		it.Event = new(AbiPaused)
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
func (it *AbiPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiPaused represents a Paused event raised by the Abi contract.
type AbiPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) FilterPaused(opts *bind.FilterOpts) (*AbiPausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AbiPausedIterator{contract: _Abi.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AbiPaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiPaused)
				if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Abi *AbiFilterer) ParsePaused(log types.Log) (*AbiPaused, error) {
	event := new(AbiPaused)
	if err := _Abi.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Abi contract.
type AbiRewardClaimedIterator struct {
	Event *AbiRewardClaimed // Event containing the contract specifics and raw log

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
func (it *AbiRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiRewardClaimed)
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
		it.Event = new(AbiRewardClaimed)
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
func (it *AbiRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiRewardClaimed represents a RewardClaimed event raised by the Abi contract.
type AbiRewardClaimed struct {
	User   common.Address
	Reward *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward, uint256 time)
func (_Abi *AbiFilterer) FilterRewardClaimed(opts *bind.FilterOpts, user []common.Address) (*AbiRewardClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiRewardClaimedIterator{contract: _Abi.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward, uint256 time)
func (_Abi *AbiFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *AbiRewardClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "RewardClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiRewardClaimed)
				if err := _Abi.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0xf01da32686223933d8a18a391060918c7f11a3648639edd87ae013e2e2731743.
//
// Solidity: event RewardClaimed(address indexed user, uint256 reward, uint256 time)
func (_Abi *AbiFilterer) ParseRewardClaimed(log types.Log) (*AbiRewardClaimed, error) {
	event := new(AbiRewardClaimed)
	if err := _Abi.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiStakeLimitsUpdatedIterator is returned from FilterStakeLimitsUpdated and is used to iterate over the raw logs and unpacked data for StakeLimitsUpdated events raised by the Abi contract.
type AbiStakeLimitsUpdatedIterator struct {
	Event *AbiStakeLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *AbiStakeLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiStakeLimitsUpdated)
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
		it.Event = new(AbiStakeLimitsUpdated)
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
func (it *AbiStakeLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiStakeLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiStakeLimitsUpdated represents a StakeLimitsUpdated event raised by the Abi contract.
type AbiStakeLimitsUpdated struct {
	Min *big.Int
	Max *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakeLimitsUpdated is a free log retrieval operation binding the contract event 0xebe4bfbb418499c23fa5fa0ea9a38ae0c885402f24ee19601c00738f1231150f.
//
// Solidity: event StakeLimitsUpdated(uint256 min, uint256 max)
func (_Abi *AbiFilterer) FilterStakeLimitsUpdated(opts *bind.FilterOpts) (*AbiStakeLimitsUpdatedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "StakeLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &AbiStakeLimitsUpdatedIterator{contract: _Abi.contract, event: "StakeLimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeLimitsUpdated is a free log subscription operation binding the contract event 0xebe4bfbb418499c23fa5fa0ea9a38ae0c885402f24ee19601c00738f1231150f.
//
// Solidity: event StakeLimitsUpdated(uint256 min, uint256 max)
func (_Abi *AbiFilterer) WatchStakeLimitsUpdated(opts *bind.WatchOpts, sink chan<- *AbiStakeLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "StakeLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiStakeLimitsUpdated)
				if err := _Abi.contract.UnpackLog(event, "StakeLimitsUpdated", log); err != nil {
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

// ParseStakeLimitsUpdated is a log parse operation binding the contract event 0xebe4bfbb418499c23fa5fa0ea9a38ae0c885402f24ee19601c00738f1231150f.
//
// Solidity: event StakeLimitsUpdated(uint256 min, uint256 max)
func (_Abi *AbiFilterer) ParseStakeLimitsUpdated(log types.Log) (*AbiStakeLimitsUpdated, error) {
	event := new(AbiStakeLimitsUpdated)
	if err := _Abi.contract.UnpackLog(event, "StakeLimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Abi contract.
type AbiStakedIterator struct {
	Event *AbiStaked // Event containing the contract specifics and raw log

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
func (it *AbiStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiStaked)
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
		it.Event = new(AbiStaked)
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
func (it *AbiStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiStaked represents a Staked event raised by the Abi contract.
type AbiStaked struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*AbiStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiStakedIterator{contract: _Abi.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AbiStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiStaked)
				if err := _Abi.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) ParseStaked(log types.Log) (*AbiStaked, error) {
	event := new(AbiStaked)
	if err := _Abi.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Abi contract.
type AbiUnpausedIterator struct {
	Event *AbiUnpaused // Event containing the contract specifics and raw log

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
func (it *AbiUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUnpaused)
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
		it.Event = new(AbiUnpaused)
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
func (it *AbiUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUnpaused represents a Unpaused event raised by the Abi contract.
type AbiUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AbiUnpausedIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AbiUnpausedIterator{contract: _Abi.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AbiUnpaused) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUnpaused)
				if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Abi *AbiFilterer) ParseUnpaused(log types.Log) (*AbiUnpaused, error) {
	event := new(AbiUnpaused)
	if err := _Abi.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Abi contract.
type AbiUnstakedIterator struct {
	Event *AbiUnstaked // Event containing the contract specifics and raw log

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
func (it *AbiUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUnstaked)
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
		it.Event = new(AbiUnstaked)
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
func (it *AbiUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUnstaked represents a Unstaked event raised by the Abi contract.
type AbiUnstaked struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*AbiUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &AbiUnstakedIterator{contract: _Abi.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *AbiUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUnstaked)
				if err := _Abi.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Abi *AbiFilterer) ParseUnstaked(log types.Log) (*AbiUnstaked, error) {
	event := new(AbiUnstaked)
	if err := _Abi.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Abi contract.
type AbiUpgradedIterator struct {
	Event *AbiUpgraded // Event containing the contract specifics and raw log

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
func (it *AbiUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUpgraded)
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
		it.Event = new(AbiUpgraded)
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
func (it *AbiUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUpgraded represents a Upgraded event raised by the Abi contract.
type AbiUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AbiUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AbiUpgradedIterator{contract: _Abi.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AbiUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUpgraded)
				if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Abi *AbiFilterer) ParseUpgraded(log types.Log) (*AbiUpgraded, error) {
	event := new(AbiUpgraded)
	if err := _Abi.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
