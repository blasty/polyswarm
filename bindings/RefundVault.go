// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RefundVaultABI is the input ABI used to generate the binding from.
const RefundVaultABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableRefunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposited\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RefundsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"weiAmount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RefundVaultBin is the compiled bytecode used for deploying new contracts.
const RefundVaultBin = `6060604052341561000f57600080fd5b60405160208061060c8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60028054600160a060020a031916600160a060020a03929092169190911760a060020a60ff021916905561057e8061008e6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166343d726d6811461009d578063521eb273146100b25780638c52dc41146100e15780638da5cb5b146100f4578063c19d93fb14610107578063cb13cddb1461013e578063f2fde38b1461016f578063f340fa011461018e578063fa89401a146101a2575b600080fd5b34156100a857600080fd5b6100b06101c1565b005b34156100bd57600080fd5b6100c561029c565b604051600160a060020a03909116815260200160405180910390f35b34156100ec57600080fd5b6100b06102ab565b34156100ff57600080fd5b6100c561033c565b341561011257600080fd5b61011a61034b565b6040518082600281111561012a57fe5b60ff16815260200191505060405180910390f35b341561014957600080fd5b61015d600160a060020a036004351661035b565b60405190815260200160405180910390f35b341561017a57600080fd5b6100b0600160a060020a036004351661036d565b6100b0600160a060020a0360043516610408565b34156101ad57600080fd5b6100b0600160a060020a036004351661048c565b60005433600160a060020a039081169116146101dc57600080fd5b60006002805460a060020a900460ff16908111156101f657fe5b1461020057600080fd5b6002805474ff00000000000000000000000000000000000000001916740200000000000000000000000000000000000000001790557f1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a60405160405180910390a1600254600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561029a57600080fd5b565b600254600160a060020a031681565b60005433600160a060020a039081169116146102c657600080fd5b60006002805460a060020a900460ff16908111156102e057fe5b146102ea57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a1790557f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8960405160405180910390a1565b600054600160a060020a031681565b60025460a060020a900460ff1681565b60016020526000908152604090205481565b60005433600160a060020a0390811691161461038857600080fd5b600160a060020a038116151561039d57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461042357600080fd5b60006002805460a060020a900460ff169081111561043d57fe5b1461044757600080fd5b600160a060020a038116600090815260016020526040902054610470903463ffffffff61053c16565b600160a060020a03909116600090815260016020526040902055565b600060016002805460a060020a900460ff16908111156104a857fe5b146104b257600080fd5b50600160a060020a038116600081815260016020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156104fb57600080fd5b81600160a060020a03167fd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d06518260405190815260200160405180910390a25050565b60008282018381101561054b57fe5b93925050505600a165627a7a723058205827e63f75dd8d46e0696a2eab110918af531ae50445596d1f3a112e4123027e0029`

// DeployRefundVault deploys a new Ethereum contract, binding an instance of RefundVault to it.
func DeployRefundVault(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *RefundVault, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundVaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RefundVaultBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RefundVault{RefundVaultCaller: RefundVaultCaller{contract: contract}, RefundVaultTransactor: RefundVaultTransactor{contract: contract}, RefundVaultFilterer: RefundVaultFilterer{contract: contract}}, nil
}

// RefundVault is an auto generated Go binding around an Ethereum contract.
type RefundVault struct {
	RefundVaultCaller     // Read-only binding to the contract
	RefundVaultTransactor // Write-only binding to the contract
	RefundVaultFilterer   // Log filterer for contract events
}

// RefundVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefundVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefundVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefundVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefundVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefundVaultSession struct {
	Contract     *RefundVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefundVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefundVaultCallerSession struct {
	Contract *RefundVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RefundVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefundVaultTransactorSession struct {
	Contract     *RefundVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RefundVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefundVaultRaw struct {
	Contract *RefundVault // Generic contract binding to access the raw methods on
}

// RefundVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefundVaultCallerRaw struct {
	Contract *RefundVaultCaller // Generic read-only contract binding to access the raw methods on
}

// RefundVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefundVaultTransactorRaw struct {
	Contract *RefundVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefundVault creates a new instance of RefundVault, bound to a specific deployed contract.
func NewRefundVault(address common.Address, backend bind.ContractBackend) (*RefundVault, error) {
	contract, err := bindRefundVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RefundVault{RefundVaultCaller: RefundVaultCaller{contract: contract}, RefundVaultTransactor: RefundVaultTransactor{contract: contract}, RefundVaultFilterer: RefundVaultFilterer{contract: contract}}, nil
}

// NewRefundVaultCaller creates a new read-only instance of RefundVault, bound to a specific deployed contract.
func NewRefundVaultCaller(address common.Address, caller bind.ContractCaller) (*RefundVaultCaller, error) {
	contract, err := bindRefundVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefundVaultCaller{contract: contract}, nil
}

// NewRefundVaultTransactor creates a new write-only instance of RefundVault, bound to a specific deployed contract.
func NewRefundVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*RefundVaultTransactor, error) {
	contract, err := bindRefundVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefundVaultTransactor{contract: contract}, nil
}

// NewRefundVaultFilterer creates a new log filterer instance of RefundVault, bound to a specific deployed contract.
func NewRefundVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*RefundVaultFilterer, error) {
	contract, err := bindRefundVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefundVaultFilterer{contract: contract}, nil
}

// bindRefundVault binds a generic wrapper to an already deployed contract.
func bindRefundVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RefundVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundVault *RefundVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundVault.Contract.RefundVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundVault *RefundVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundVault.Contract.RefundVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundVault *RefundVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundVault.Contract.RefundVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RefundVault *RefundVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RefundVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RefundVault *RefundVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RefundVault *RefundVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RefundVault.Contract.contract.Transact(opts, method, params...)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited( address) constant returns(uint256)
func (_RefundVault *RefundVaultCaller) Deposited(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RefundVault.contract.Call(opts, out, "deposited", arg0)
	return *ret0, err
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited( address) constant returns(uint256)
func (_RefundVault *RefundVaultSession) Deposited(arg0 common.Address) (*big.Int, error) {
	return _RefundVault.Contract.Deposited(&_RefundVault.CallOpts, arg0)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited( address) constant returns(uint256)
func (_RefundVault *RefundVaultCallerSession) Deposited(arg0 common.Address) (*big.Int, error) {
	return _RefundVault.Contract.Deposited(&_RefundVault.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundVault *RefundVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundVault.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundVault *RefundVaultSession) Owner() (common.Address, error) {
	return _RefundVault.Contract.Owner(&_RefundVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RefundVault *RefundVaultCallerSession) Owner() (common.Address, error) {
	return _RefundVault.Contract.Owner(&_RefundVault.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RefundVault *RefundVaultCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RefundVault.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RefundVault *RefundVaultSession) State() (uint8, error) {
	return _RefundVault.Contract.State(&_RefundVault.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_RefundVault *RefundVaultCallerSession) State() (uint8, error) {
	return _RefundVault.Contract.State(&_RefundVault.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundVault *RefundVaultCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RefundVault.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundVault *RefundVaultSession) Wallet() (common.Address, error) {
	return _RefundVault.Contract.Wallet(&_RefundVault.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_RefundVault *RefundVaultCallerSession) Wallet() (common.Address, error) {
	return _RefundVault.Contract.Wallet(&_RefundVault.CallOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundVault *RefundVaultTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundVault.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundVault *RefundVaultSession) Close() (*types.Transaction, error) {
	return _RefundVault.Contract.Close(&_RefundVault.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RefundVault *RefundVaultTransactorSession) Close() (*types.Transaction, error) {
	return _RefundVault.Contract.Close(&_RefundVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(investor address) returns()
func (_RefundVault *RefundVaultTransactor) Deposit(opts *bind.TransactOpts, investor common.Address) (*types.Transaction, error) {
	return _RefundVault.contract.Transact(opts, "deposit", investor)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(investor address) returns()
func (_RefundVault *RefundVaultSession) Deposit(investor common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.Deposit(&_RefundVault.TransactOpts, investor)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(investor address) returns()
func (_RefundVault *RefundVaultTransactorSession) Deposit(investor common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.Deposit(&_RefundVault.TransactOpts, investor)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundVault *RefundVaultTransactor) EnableRefunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RefundVault.contract.Transact(opts, "enableRefunds")
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundVault *RefundVaultSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundVault.Contract.EnableRefunds(&_RefundVault.TransactOpts)
}

// EnableRefunds is a paid mutator transaction binding the contract method 0x8c52dc41.
//
// Solidity: function enableRefunds() returns()
func (_RefundVault *RefundVaultTransactorSession) EnableRefunds() (*types.Transaction, error) {
	return _RefundVault.Contract.EnableRefunds(&_RefundVault.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(investor address) returns()
func (_RefundVault *RefundVaultTransactor) Refund(opts *bind.TransactOpts, investor common.Address) (*types.Transaction, error) {
	return _RefundVault.contract.Transact(opts, "refund", investor)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(investor address) returns()
func (_RefundVault *RefundVaultSession) Refund(investor common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.Refund(&_RefundVault.TransactOpts, investor)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(investor address) returns()
func (_RefundVault *RefundVaultTransactorSession) Refund(investor common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.Refund(&_RefundVault.TransactOpts, investor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundVault *RefundVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RefundVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundVault *RefundVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.TransferOwnership(&_RefundVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_RefundVault *RefundVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RefundVault.Contract.TransferOwnership(&_RefundVault.TransactOpts, newOwner)
}

// RefundVaultClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the RefundVault contract.
type RefundVaultClosedIterator struct {
	Event *RefundVaultClosed // Event containing the contract specifics and raw log

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
func (it *RefundVaultClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundVaultClosed)
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
		it.Event = new(RefundVaultClosed)
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
func (it *RefundVaultClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundVaultClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundVaultClosed represents a Closed event raised by the RefundVault contract.
type RefundVaultClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_RefundVault *RefundVaultFilterer) FilterClosed(opts *bind.FilterOpts) (*RefundVaultClosedIterator, error) {

	logs, sub, err := _RefundVault.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &RefundVaultClosedIterator{contract: _RefundVault.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_RefundVault *RefundVaultFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *RefundVaultClosed) (event.Subscription, error) {

	logs, sub, err := _RefundVault.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundVaultClosed)
				if err := _RefundVault.contract.UnpackLog(event, "Closed", log); err != nil {
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

// RefundVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RefundVault contract.
type RefundVaultOwnershipTransferredIterator struct {
	Event *RefundVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RefundVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundVaultOwnershipTransferred)
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
		it.Event = new(RefundVaultOwnershipTransferred)
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
func (it *RefundVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundVaultOwnershipTransferred represents a OwnershipTransferred event raised by the RefundVault contract.
type RefundVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RefundVault *RefundVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RefundVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RefundVaultOwnershipTransferredIterator{contract: _RefundVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RefundVault *RefundVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RefundVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RefundVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundVaultOwnershipTransferred)
				if err := _RefundVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RefundVaultRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the RefundVault contract.
type RefundVaultRefundedIterator struct {
	Event *RefundVaultRefunded // Event containing the contract specifics and raw log

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
func (it *RefundVaultRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundVaultRefunded)
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
		it.Event = new(RefundVaultRefunded)
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
func (it *RefundVaultRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundVaultRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundVaultRefunded represents a Refunded event raised by the RefundVault contract.
type RefundVaultRefunded struct {
	Beneficiary common.Address
	WeiAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(beneficiary indexed address, weiAmount uint256)
func (_RefundVault *RefundVaultFilterer) FilterRefunded(opts *bind.FilterOpts, beneficiary []common.Address) (*RefundVaultRefundedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundVault.contract.FilterLogs(opts, "Refunded", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &RefundVaultRefundedIterator{contract: _RefundVault.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(beneficiary indexed address, weiAmount uint256)
func (_RefundVault *RefundVaultFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *RefundVaultRefunded, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _RefundVault.contract.WatchLogs(opts, "Refunded", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundVaultRefunded)
				if err := _RefundVault.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// RefundVaultRefundsEnabledIterator is returned from FilterRefundsEnabled and is used to iterate over the raw logs and unpacked data for RefundsEnabled events raised by the RefundVault contract.
type RefundVaultRefundsEnabledIterator struct {
	Event *RefundVaultRefundsEnabled // Event containing the contract specifics and raw log

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
func (it *RefundVaultRefundsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefundVaultRefundsEnabled)
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
		it.Event = new(RefundVaultRefundsEnabled)
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
func (it *RefundVaultRefundsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefundVaultRefundsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefundVaultRefundsEnabled represents a RefundsEnabled event raised by the RefundVault contract.
type RefundVaultRefundsEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefundsEnabled is a free log retrieval operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundVault *RefundVaultFilterer) FilterRefundsEnabled(opts *bind.FilterOpts) (*RefundVaultRefundsEnabledIterator, error) {

	logs, sub, err := _RefundVault.contract.FilterLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return &RefundVaultRefundsEnabledIterator{contract: _RefundVault.contract, event: "RefundsEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundsEnabled is a free log subscription operation binding the contract event 0x599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b89.
//
// Solidity: event RefundsEnabled()
func (_RefundVault *RefundVaultFilterer) WatchRefundsEnabled(opts *bind.WatchOpts, sink chan<- *RefundVaultRefundsEnabled) (event.Subscription, error) {

	logs, sub, err := _RefundVault.contract.WatchLogs(opts, "RefundsEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefundVaultRefundsEnabled)
				if err := _RefundVault.contract.UnpackLog(event, "RefundsEnabled", log); err != nil {
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
