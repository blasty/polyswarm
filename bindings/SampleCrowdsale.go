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

// SampleCrowdsaleABI is the input ABI used to generate the binding from.
const SampleCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"goalReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasEnded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_endTime\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_goal\",\"type\":\"uint256\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// SampleCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const SampleCrowdsaleBin = `60606040526007805460a060020a60ff0219169055341561001f57600080fd5b60405160c080612815833981016040528080519190602001805191906020018051919060200180519190602001805191906020018051915083905082878787854284101561006c57600080fd5b8383101561007957600080fd5b6000821161008657600080fd5b600160a060020a038116151561009b57600080fd5b6100b064010000000061081661019882021704565b60008054600160a060020a0319908116600160a060020a039384161782556001969096556002949094556004929092556003805490941691161790915581116100f857600080fd5b60065560078054600160a060020a03191633600160a060020a03161790556000811161012357600080fd5b600354600160a060020a03166101376101bd565b600160a060020a039091168152602001604051809103906000f080151561015d57600080fd5b60098054600160a060020a031916600160a060020a03929092169190911790556008558183111561018d57600080fd5b5050505050506101dd565b60006101a26101cd565b604051809103906000f08015156101b857600080fd5b905090565b60405161060c8061163683390190565b604051610bd380611c4283390190565b61144a806101ec6000396000f3006060604052600436106100d75763ffffffff60e060020a6000350416632c4e722e81146100e25780633197cbb614610107578063355274ea1461011a578063401938831461012d5780634042b66f146101405780634bb278f314610153578063521eb2731461016657806378e97925146101955780637d3d6522146101a85780638d4e4083146101cf5780638da5cb5b146101e2578063b5545a3c146101f5578063ec8ac4d814610208578063ecb70fb71461021c578063f2fde38b1461022f578063fbfa77cf1461024e578063fc0c546a14610261575b6100e033610274565b005b34156100ed57600080fd5b6100f561038d565b60405190815260200160405180910390f35b341561011257600080fd5b6100f5610393565b341561012557600080fd5b6100f5610399565b341561013857600080fd5b6100f561039f565b341561014b57600080fd5b6100f56103a5565b341561015e57600080fd5b6100e06103ab565b341561017157600080fd5b61017961046c565b604051600160a060020a03909116815260200160405180910390f35b34156101a057600080fd5b6100f561047b565b34156101b357600080fd5b6101bb610481565b604051901515815260200160405180910390f35b34156101da57600080fd5b6101bb61048c565b34156101ed57600080fd5b6101796104ad565b341561020057600080fd5b6100e06104bc565b6100e0600160a060020a0360043516610274565b341561022757600080fd5b6101bb610554565b341561023a57600080fd5b6100e0600160a060020a0360043516610576565b341561025957600080fd5b610179610611565b341561026c57600080fd5b610179610620565b600080600160a060020a038316151561028c57600080fd5b61029461062f565b151561029f57600080fd5b3491506102ab82610662565b6005549091506102c1908363ffffffff61067f16565b600555600054600160a060020a03166340c10f19848360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561031a57600080fd5b5af1151561032757600080fd5b505050604051805190505082600160a060020a031633600160a060020a03167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18848460405191825260208201526040908101905180910390a3610388610699565b505050565b60045481565b60025481565b60065481565b60085481565b60055481565b60075433600160a060020a039081169116146103c657600080fd5b60075474010000000000000000000000000000000000000000900460ff16156103ee57600080fd5b6103f6610554565b151561040157600080fd5b6104096106fc565b7f6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b1768160405160405180910390a16007805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600354600160a060020a031681565b60015481565b600854600554101590565b60075474010000000000000000000000000000000000000000900460ff1681565b600754600160a060020a031681565b60075474010000000000000000000000000000000000000000900460ff1615156104e557600080fd5b6104ed610481565b156104f757600080fd5b600954600160a060020a031663fa89401a3360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561054757600080fd5b5af1151561038857600080fd5b6006546005546000919010156105686107b3565b806105705750805b91505090565b60075433600160a060020a0390811691161461059157600080fd5b600160a060020a03811615156105a657600080fd5b600754600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600954600160a060020a031681565b600054600160a060020a031681565b60008060065461064a3460055461067f90919063ffffffff16565b111590506106566107bb565b80156105705750919050565b6000610679600454836107eb90919063ffffffff16565b92915050565b60008282018381101561068e57fe5b8091505b5092915050565b600954600160a060020a031663f340fa01343360405160e060020a63ffffffff8516028152600160a060020a0390911660048201526024016000604051808303818588803b15156106e957600080fd5b5af115156106f657600080fd5b50505050565b610704610481565b1561075d57600954600160a060020a03166343d726d66040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561074857600080fd5b5af1151561075557600080fd5b5050506107ad565b600954600160a060020a0316638c52dc416040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561079c57600080fd5b5af115156107a957600080fd5b5050505b6107b15b565b600254421190565b600080600060015442101580156107d457506002544211155b9150503415158180156107e45750805b9250505090565b6000808315156107fe5760009150610692565b5082820282848281151561080e57fe5b041461068e57fe5b600061082061083b565b604051809103906000f080151561083657600080fd5b905090565b604051610bd38061084c833901905600606060405260038054600160a860020a03191633600160a060020a0316179055610ba58061002e6000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100ea57806306fdde0314610111578063095ea7b31461019b57806318160ddd146101bd57806323b872dd146101e2578063313ce5671461020a57806340c10f1914610233578063661884631461025557806370a08231146102775780637d64bcb4146102965780638da5cb5b146102a957806395d89b41146102d8578063a9059cbb146102eb578063d73dd6231461030d578063dd62ed3e1461032f578063f2fde38b14610354575b600080fd5b34156100f557600080fd5b6100fd610375565b604051901515815260200160405180910390f35b341561011c57600080fd5b610124610396565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610160578082015183820152602001610148565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101a657600080fd5b6100fd600160a060020a03600435166024356103cd565b34156101c857600080fd5b6101d0610439565b60405190815260200160405180910390f35b34156101ed57600080fd5b6100fd600160a060020a036004358116906024351660443561043f565b341561021557600080fd5b61021d6105c1565b60405160ff909116815260200160405180910390f35b341561023e57600080fd5b6100fd600160a060020a03600435166024356105c6565b341561026057600080fd5b6100fd600160a060020a03600435166024356106e4565b341561028257600080fd5b6101d0600160a060020a03600435166107de565b34156102a157600080fd5b6100fd6107f9565b34156102b457600080fd5b6102bc6108a6565b604051600160a060020a03909116815260200160405180910390f35b34156102e357600080fd5b6101246108b5565b34156102f657600080fd5b6100fd600160a060020a03600435166024356108ec565b341561031857600080fd5b6100fd600160a060020a03600435166024356109e7565b341561033a57600080fd5b6101d0600160a060020a0360043581169060243516610a8b565b341561035f57600080fd5b610373600160a060020a0360043516610ab6565b005b60035474010000000000000000000000000000000000000000900460ff1681565b60408051908101604052601681527f53616d706c652043726f776473616c6520546f6b656e00000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561045657600080fd5b600160a060020a03841660009081526001602052604090205482111561047b57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156104ae57600080fd5b600160a060020a0384166000908152600160205260409020546104d7908363ffffffff610b5116565b600160a060020a03808616600090815260016020526040808220939093559085168152205461050c908363ffffffff610b6316565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610554908363ffffffff610b5116565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b601281565b60035460009033600160a060020a039081169116146105e457600080fd5b60035474010000000000000000000000000000000000000000900460ff161561060c57600080fd5b60005461061f908363ffffffff610b6316565b6000908155600160a060020a03841681526001602052604090205461064a908363ffffffff610b6316565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561074157600160a060020a033381166000908152600260209081526040808320938816835292905290812055610778565b610751818463ffffffff610b5116565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461081757600080fd5b60035474010000000000000000000000000000000000000000900460ff161561083f57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f5343540000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a038316151561090357600080fd5b600160a060020a03331660009081526001602052604090205482111561092857600080fd5b600160a060020a033316600090815260016020526040902054610951908363ffffffff610b5116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610986908363ffffffff610b6316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610a1f908363ffffffff610b6316565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a03908116911614610ad157600080fd5b600160a060020a0381161515610ae657600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610b5d57fe5b50900390565b600082820183811015610b7257fe5b93925050505600a165627a7a723058205bd5872c788a571627f027d8dc37513f521fc4a39510a73a68d1dbdeb7c26abe0029a165627a7a723058209e51e644ce62d663f8667768eb01e02ccdd4709d32d860b2add48c220630609200296060604052341561000f57600080fd5b60405160208061060c8339810160405280805160008054600160a060020a03191633600160a060020a039081169190911790915590925082161515905061005557600080fd5b60028054600160a060020a031916600160a060020a03929092169190911760a060020a60ff021916905561057e8061008e6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166343d726d6811461009d578063521eb273146100b25780638c52dc41146100e15780638da5cb5b146100f4578063c19d93fb14610107578063cb13cddb1461013e578063f2fde38b1461016f578063f340fa011461018e578063fa89401a146101a2575b600080fd5b34156100a857600080fd5b6100b06101c1565b005b34156100bd57600080fd5b6100c561029c565b604051600160a060020a03909116815260200160405180910390f35b34156100ec57600080fd5b6100b06102ab565b34156100ff57600080fd5b6100c561033c565b341561011257600080fd5b61011a61034b565b6040518082600281111561012a57fe5b60ff16815260200191505060405180910390f35b341561014957600080fd5b61015d600160a060020a036004351661035b565b60405190815260200160405180910390f35b341561017a57600080fd5b6100b0600160a060020a036004351661036d565b6100b0600160a060020a0360043516610408565b34156101ad57600080fd5b6100b0600160a060020a036004351661048c565b60005433600160a060020a039081169116146101dc57600080fd5b60006002805460a060020a900460ff16908111156101f657fe5b1461020057600080fd5b6002805474ff00000000000000000000000000000000000000001916740200000000000000000000000000000000000000001790557f1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a60405160405180910390a1600254600160a060020a039081169030163180156108fc0290604051600060405180830381858888f19350505050151561029a57600080fd5b565b600254600160a060020a031681565b60005433600160a060020a039081169116146102c657600080fd5b60006002805460a060020a900460ff16908111156102e057fe5b146102ea57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a1790557f599d8e5a83cffb867d051598c4d70e805d59802d8081c1c7d6dffc5b6aca2b8960405160405180910390a1565b600054600160a060020a031681565b60025460a060020a900460ff1681565b60016020526000908152604090205481565b60005433600160a060020a0390811691161461038857600080fd5b600160a060020a038116151561039d57600080fd5b600054600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461042357600080fd5b60006002805460a060020a900460ff169081111561043d57fe5b1461044757600080fd5b600160a060020a038116600090815260016020526040902054610470903463ffffffff61053c16565b600160a060020a03909116600090815260016020526040902055565b600060016002805460a060020a900460ff16908111156104a857fe5b146104b257600080fd5b50600160a060020a038116600081815260016020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156104fb57600080fd5b81600160a060020a03167fd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d06518260405190815260200160405180910390a25050565b60008282018381101561054b57fe5b93925050505600a165627a7a723058205827e63f75dd8d46e0696a2eab110918af531ae50445596d1f3a112e4123027e0029606060405260038054600160a860020a03191633600160a060020a0316179055610ba58061002e6000396000f3006060604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100ea57806306fdde0314610111578063095ea7b31461019b57806318160ddd146101bd57806323b872dd146101e2578063313ce5671461020a57806340c10f1914610233578063661884631461025557806370a08231146102775780637d64bcb4146102965780638da5cb5b146102a957806395d89b41146102d8578063a9059cbb146102eb578063d73dd6231461030d578063dd62ed3e1461032f578063f2fde38b14610354575b600080fd5b34156100f557600080fd5b6100fd610375565b604051901515815260200160405180910390f35b341561011c57600080fd5b610124610396565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610160578082015183820152602001610148565b50505050905090810190601f16801561018d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101a657600080fd5b6100fd600160a060020a03600435166024356103cd565b34156101c857600080fd5b6101d0610439565b60405190815260200160405180910390f35b34156101ed57600080fd5b6100fd600160a060020a036004358116906024351660443561043f565b341561021557600080fd5b61021d6105c1565b60405160ff909116815260200160405180910390f35b341561023e57600080fd5b6100fd600160a060020a03600435166024356105c6565b341561026057600080fd5b6100fd600160a060020a03600435166024356106e4565b341561028257600080fd5b6101d0600160a060020a03600435166107de565b34156102a157600080fd5b6100fd6107f9565b34156102b457600080fd5b6102bc6108a6565b604051600160a060020a03909116815260200160405180910390f35b34156102e357600080fd5b6101246108b5565b34156102f657600080fd5b6100fd600160a060020a03600435166024356108ec565b341561031857600080fd5b6100fd600160a060020a03600435166024356109e7565b341561033a57600080fd5b6101d0600160a060020a0360043581169060243516610a8b565b341561035f57600080fd5b610373600160a060020a0360043516610ab6565b005b60035474010000000000000000000000000000000000000000900460ff1681565b60408051908101604052601681527f53616d706c652043726f776473616c6520546f6b656e00000000000000000000602082015281565b600160a060020a03338116600081815260026020908152604080832094871680845294909152808220859055909291907f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b6000600160a060020a038316151561045657600080fd5b600160a060020a03841660009081526001602052604090205482111561047b57600080fd5b600160a060020a03808516600090815260026020908152604080832033909416835292905220548211156104ae57600080fd5b600160a060020a0384166000908152600160205260409020546104d7908363ffffffff610b5116565b600160a060020a03808616600090815260016020526040808220939093559085168152205461050c908363ffffffff610b6316565b600160a060020a03808516600090815260016020908152604080832094909455878316825260028152838220339093168252919091522054610554908363ffffffff610b5116565b600160a060020a03808616600081815260026020908152604080832033861684529091529081902093909355908516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060019392505050565b601281565b60035460009033600160a060020a039081169116146105e457600080fd5b60035474010000000000000000000000000000000000000000900460ff161561060c57600080fd5b60005461061f908363ffffffff610b6316565b6000908155600160a060020a03841681526001602052604090205461064a908363ffffffff610b6316565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a2600160a060020a03831660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a350600192915050565b600160a060020a0333811660009081526002602090815260408083209386168352929052908120548083111561074157600160a060020a033381166000908152600260209081526040808320938816835292905290812055610778565b610751818463ffffffff610b5116565b600160a060020a033381166000908152600260209081526040808320938916835292905220555b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020547f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925915190815260200160405180910390a35060019392505050565b600160a060020a031660009081526001602052604090205490565b60035460009033600160a060020a0390811691161461081757600080fd5b60035474010000000000000000000000000000000000000000900460ff161561083f57600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a150600190565b600354600160a060020a031681565b60408051908101604052600381527f5343540000000000000000000000000000000000000000000000000000000000602082015281565b6000600160a060020a038316151561090357600080fd5b600160a060020a03331660009081526001602052604090205482111561092857600080fd5b600160a060020a033316600090815260016020526040902054610951908363ffffffff610b5116565b600160a060020a033381166000908152600160205260408082209390935590851681522054610986908363ffffffff610b6316565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a350600192915050565b600160a060020a033381166000908152600260209081526040808320938616835292905290812054610a1f908363ffffffff610b6316565b600160a060020a0333811660008181526002602090815260408083209489168084529490915290819020849055919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60035433600160a060020a03908116911614610ad157600080fd5b600160a060020a0381161515610ae657600080fd5b600354600160a060020a0380831691167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610b5d57fe5b50900390565b600082820183811015610b7257fe5b93925050505600a165627a7a723058205bd5872c788a571627f027d8dc37513f521fc4a39510a73a68d1dbdeb7c26abe0029`

// DeploySampleCrowdsale deploys a new Ethereum contract, binding an instance of SampleCrowdsale to it.
func DeploySampleCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _startTime *big.Int, _endTime *big.Int, _rate *big.Int, _goal *big.Int, _cap *big.Int, _wallet common.Address) (common.Address, *types.Transaction, *SampleCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SampleCrowdsaleBin), backend, _startTime, _endTime, _rate, _goal, _cap, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SampleCrowdsale{SampleCrowdsaleCaller: SampleCrowdsaleCaller{contract: contract}, SampleCrowdsaleTransactor: SampleCrowdsaleTransactor{contract: contract}, SampleCrowdsaleFilterer: SampleCrowdsaleFilterer{contract: contract}}, nil
}

// SampleCrowdsale is an auto generated Go binding around an Ethereum contract.
type SampleCrowdsale struct {
	SampleCrowdsaleCaller     // Read-only binding to the contract
	SampleCrowdsaleTransactor // Write-only binding to the contract
	SampleCrowdsaleFilterer   // Log filterer for contract events
}

// SampleCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SampleCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SampleCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SampleCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SampleCrowdsaleSession struct {
	Contract     *SampleCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SampleCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SampleCrowdsaleCallerSession struct {
	Contract *SampleCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SampleCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SampleCrowdsaleTransactorSession struct {
	Contract     *SampleCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SampleCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SampleCrowdsaleRaw struct {
	Contract *SampleCrowdsale // Generic contract binding to access the raw methods on
}

// SampleCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SampleCrowdsaleCallerRaw struct {
	Contract *SampleCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// SampleCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SampleCrowdsaleTransactorRaw struct {
	Contract *SampleCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSampleCrowdsale creates a new instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsale(address common.Address, backend bind.ContractBackend) (*SampleCrowdsale, error) {
	contract, err := bindSampleCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsale{SampleCrowdsaleCaller: SampleCrowdsaleCaller{contract: contract}, SampleCrowdsaleTransactor: SampleCrowdsaleTransactor{contract: contract}, SampleCrowdsaleFilterer: SampleCrowdsaleFilterer{contract: contract}}, nil
}

// NewSampleCrowdsaleCaller creates a new read-only instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*SampleCrowdsaleCaller, error) {
	contract, err := bindSampleCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleCaller{contract: contract}, nil
}

// NewSampleCrowdsaleTransactor creates a new write-only instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SampleCrowdsaleTransactor, error) {
	contract, err := bindSampleCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTransactor{contract: contract}, nil
}

// NewSampleCrowdsaleFilterer creates a new log filterer instance of SampleCrowdsale, bound to a specific deployed contract.
func NewSampleCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SampleCrowdsaleFilterer, error) {
	contract, err := bindSampleCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleFilterer{contract: contract}, nil
}

// bindSampleCrowdsale binds a generic wrapper to an already deployed contract.
func bindSampleCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SampleCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsale.Contract.SampleCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.SampleCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsale *SampleCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.SampleCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SampleCrowdsale *SampleCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SampleCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SampleCrowdsale *SampleCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SampleCrowdsale *SampleCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Cap() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Cap(&_SampleCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Cap(&_SampleCrowdsale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "endTime")
	return *ret0, err
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) EndTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.EndTime(&_SampleCrowdsale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) EndTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.EndTime(&_SampleCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Goal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "goal")
	return *ret0, err
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Goal() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Goal(&_SampleCrowdsale.CallOpts)
}

// Goal is a free data retrieval call binding the contract method 0x40193883.
//
// Solidity: function goal() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Goal() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Goal(&_SampleCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) GoalReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "goalReached")
	return *ret0, err
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) GoalReached() (bool, error) {
	return _SampleCrowdsale.Contract.GoalReached(&_SampleCrowdsale.CallOpts)
}

// GoalReached is a free data retrieval call binding the contract method 0x7d3d6522.
//
// Solidity: function goalReached() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) GoalReached() (bool, error) {
	return _SampleCrowdsale.Contract.GoalReached(&_SampleCrowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) HasEnded(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "hasEnded")
	return *ret0, err
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) HasEnded() (bool, error) {
	return _SampleCrowdsale.Contract.HasEnded(&_SampleCrowdsale.CallOpts)
}

// HasEnded is a free data retrieval call binding the contract method 0xecb70fb7.
//
// Solidity: function hasEnded() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) HasEnded() (bool, error) {
	return _SampleCrowdsale.Contract.HasEnded(&_SampleCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleSession) IsFinalized() (bool, error) {
	return _SampleCrowdsale.Contract.IsFinalized(&_SampleCrowdsale.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) IsFinalized() (bool, error) {
	return _SampleCrowdsale.Contract.IsFinalized(&_SampleCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Owner() (common.Address, error) {
	return _SampleCrowdsale.Contract.Owner(&_SampleCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _SampleCrowdsale.Contract.Owner(&_SampleCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) Rate() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Rate(&_SampleCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _SampleCrowdsale.Contract.Rate(&_SampleCrowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) StartTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.StartTime(&_SampleCrowdsale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) StartTime() (*big.Int, error) {
	return _SampleCrowdsale.Contract.StartTime(&_SampleCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Token() (common.Address, error) {
	return _SampleCrowdsale.Contract.Token(&_SampleCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Token() (common.Address, error) {
	return _SampleCrowdsale.Contract.Token(&_SampleCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "vault")
	return *ret0, err
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Vault() (common.Address, error) {
	return _SampleCrowdsale.Contract.Vault(&_SampleCrowdsale.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Vault() (common.Address, error) {
	return _SampleCrowdsale.Contract.Vault(&_SampleCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleSession) Wallet() (common.Address, error) {
	return _SampleCrowdsale.Contract.Wallet(&_SampleCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _SampleCrowdsale.Contract.Wallet(&_SampleCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SampleCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _SampleCrowdsale.Contract.WeiRaised(&_SampleCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_SampleCrowdsale *SampleCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _SampleCrowdsale.Contract.WeiRaised(&_SampleCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.BuyTokens(&_SampleCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(beneficiary address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.BuyTokens(&_SampleCrowdsale.TransactOpts, beneficiary)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) ClaimRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "claimRefund")
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) ClaimRefund() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.ClaimRefund(&_SampleCrowdsale.TransactOpts)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xb5545a3c.
//
// Solidity: function claimRefund() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) ClaimRefund() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.ClaimRefund(&_SampleCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) Finalize() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.Finalize(&_SampleCrowdsale.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) Finalize() (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.Finalize(&_SampleCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.TransferOwnership(&_SampleCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SampleCrowdsale *SampleCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SampleCrowdsale.Contract.TransferOwnership(&_SampleCrowdsale.TransactOpts, newOwner)
}

// SampleCrowdsaleFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the SampleCrowdsale contract.
type SampleCrowdsaleFinalizedIterator struct {
	Event *SampleCrowdsaleFinalized // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleFinalized)
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
		it.Event = new(SampleCrowdsaleFinalized)
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
func (it *SampleCrowdsaleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleFinalized represents a Finalized event raised by the SampleCrowdsale contract.
type SampleCrowdsaleFinalized struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_SampleCrowdsale *SampleCrowdsaleFilterer) FilterFinalized(opts *bind.FilterOpts) (*SampleCrowdsaleFinalizedIterator, error) {

	logs, sub, err := _SampleCrowdsale.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleFinalizedIterator{contract: _SampleCrowdsale.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x6823b073d48d6e3a7d385eeb601452d680e74bb46afe3255a7d778f3a9b17681.
//
// Solidity: event Finalized()
func (_SampleCrowdsale *SampleCrowdsaleFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleFinalized) (event.Subscription, error) {

	logs, sub, err := _SampleCrowdsale.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleFinalized)
				if err := _SampleCrowdsale.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// SampleCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SampleCrowdsale contract.
type SampleCrowdsaleOwnershipTransferredIterator struct {
	Event *SampleCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleOwnershipTransferred)
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
		it.Event = new(SampleCrowdsaleOwnershipTransferred)
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
func (it *SampleCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the SampleCrowdsale contract.
type SampleCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SampleCrowdsale *SampleCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SampleCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SampleCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleOwnershipTransferredIterator{contract: _SampleCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SampleCrowdsale *SampleCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SampleCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleOwnershipTransferred)
				if err := _SampleCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SampleCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the SampleCrowdsale contract.
type SampleCrowdsaleTokenPurchaseIterator struct {
	Event *SampleCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *SampleCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SampleCrowdsaleTokenPurchase)
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
		it.Event = new(SampleCrowdsaleTokenPurchase)
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
func (it *SampleCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SampleCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SampleCrowdsaleTokenPurchase represents a TokenPurchase event raised by the SampleCrowdsale contract.
type SampleCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_SampleCrowdsale *SampleCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*SampleCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SampleCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SampleCrowdsaleTokenPurchaseIterator{contract: _SampleCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: event TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_SampleCrowdsale *SampleCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *SampleCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SampleCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SampleCrowdsaleTokenPurchase)
				if err := _SampleCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
