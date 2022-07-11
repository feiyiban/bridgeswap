// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgev1

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

// Bridgev1MetaData contains all meta data concerning the Bridgev1 contract.
var Bridgev1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MAPTransferIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromChain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MAPTransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selfChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toChain\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_receiver\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toChainId\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200262c3803806200262c833981810160405281019062000037919062000261565b6200004c6000801b33620000a760201b60201c565b6200007e7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0833620000a760201b60201c565b6000600260006101000a81548160ff021916908315150217905550806003819055505062000293565b620000b98282620000bd60201b60201c565b5050565b620000cf8282620001ae60201b60201c565b620001aa57600180600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506200014f6200021960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600033905090565b600080fd5b6000819050919050565b6200023b8162000226565b81146200024757600080fd5b50565b6000815190506200025b8162000230565b92915050565b6000602082840312156200027a576200027962000221565b5b60006200028a848285016200024a565b91505092915050565b61238980620002a36000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063a217fddf116100a2578063d547741f11610071578063d547741f146102a5578063d9caed12146102c1578063ec87621c146102dd578063ef146cf6146102fb578063fe4b84df1461031757610116565b8063a217fddf14610243578063ac18de4314610261578063cc9e3e891461027d578063d431b1ac1461029b57610116565b806336568abe116100e957806336568abe146101b35780635c975abb146101cf578063848cb5c6146101ed5780638c7cb280146101f757806391d148541461021357610116565b806301ffc9a71461011b578063248a9ca31461014b5780632d06177a1461017b5780632f2ff15d14610197575b600080fd5b6101356004803603810190610130919061146a565b610333565b60405161014291906114b2565b60405180910390f35b61016560048036038101906101609190611503565b6103ad565b604051610172919061153f565b60405180910390f35b610195600480360381019061019091906115b8565b6103cd565b005b6101b160048036038101906101ac91906115e5565b610408565b005b6101cd60048036038101906101c891906115e5565b610429565b005b6101d76104ac565b6040516101e491906114b2565b60405180910390f35b6101f56104c3565b005b610211600480360381019061020c91906117a1565b610536565b005b61022d600480360381019061022891906115e5565b610689565b60405161023a91906114b2565b60405180910390f35b61024b6106f4565b604051610258919061153f565b60405180910390f35b61027b600480360381019061027691906115b8565b6106fb565b005b610285610736565b6040516102929190611833565b60405180910390f35b6102a361073c565b005b6102bf60048036038101906102ba91906115e5565b6107af565b005b6102db60048036038101906102d6919061188c565b6107d0565b005b6102e56108bd565b6040516102f2919061153f565b60405180910390f35b610315600480360381019061031091906118df565b6108e1565b005b610331600480360381019061032c919061195a565b610974565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103a657506103a582610a08565b5b9050919050565b600060016000838152602001908152602001600020600101549050919050565b6000801b6103da81610a72565b6104047f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0883610a86565b5050565b610411826103ad565b61041a81610a72565b6104248383610a94565b505050565b610431610b74565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461049e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049590611a0a565b60405180910390fd5b6104a88282610b7c565b5050565b6000600260009054906101000a900460ff16905090565b6104ed7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0833610689565b61052c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052390611a76565b60405180910390fd5b610534610c5e565b565b61053e6104ac565b1561057e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057590611ae2565b60405180910390fd5b818473ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016105b89190611b11565b602060405180830381865afa1580156105d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f99190611b41565b101561063a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063190611bba565b60405180910390fd5b61064684333085610d00565b806003547fc2cf97cb20e3fec72bbd1515ad10ad56582e8085858049990a34a1abe8fbee3b858560405161067b929190611c51565b60405180910390a350505050565b60006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b6000801b61070881610a72565b6107327f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0883610b7c565b5050565b60035481565b6107667f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0833610689565b6107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c90611a76565b60405180910390fd5b6107ad610e39565b565b6107b8826103ad565b6107c181610a72565b6107cb8383610b7c565b505050565b6107fa7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0833610689565b610839576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083090611a76565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610874929190611ce0565b6020604051808303816000875af1158015610893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b79190611d35565b50505050565b7f241ecf16d79d0f8dbfb92cbc07fe17840425976cf0667f022fe9877caa831b0881565b6108e96104ac565b15610929576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092090611ae2565b60405180910390fd5b610934858585610edc565b80827f480c1cd6cfcaa4162f54fd97792200e72dbb4bd1a21259f882414f0d9188941d856040516109659190611833565b60405180910390a35050505050565b60006109806001611012565b905080156109a4576001600060016101000a81548160ff0219169083151502179055505b816003819055508015610a045760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516109fb9190611daa565b60405180910390a15b5050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610a8381610a7e610b74565b611102565b50565b610a908282610a94565b5050565b610a9e8282610689565b610b7057600180600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610b15610b74565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b610b868282610689565b15610c5a5760006001600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610bff610b74565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b610c666104ac565b610ca5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9c90611e11565b60405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610ce9610b74565b604051610cf69190611b11565b60405180910390a1565b6000808573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610d3493929190611e31565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610d829190611eaf565b6000604051808303816000865af19150503d8060008114610dbf576040519150601f19603f3d011682016040523d82523d6000602084013e610dc4565b606091505b5091509150818015610df25750600081511480610df1575080806020019051810190610df09190611d35565b5b5b610e31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2890611f38565b60405180910390fd5b505050505050565b610e416104ac565b15610e81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7890611ae2565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610ec5610b74565b604051610ed29190611b11565b60405180910390a1565b6000808473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401610f0e929190611f58565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610f5c9190611eaf565b6000604051808303816000865af19150503d8060008114610f99576040519150601f19603f3d011682016040523d82523d6000602084013e610f9e565b606091505b5091509150818015610fcc5750600081511480610fcb575080806020019051810190610fca9190611d35565b5b5b61100b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100290611fcd565b60405180910390fd5b5050505050565b60008060019054906101000a900460ff16156110895760018260ff16148015611041575061103f3061119f565b155b611080576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110779061205f565b60405180910390fd5b600090506110fd565b8160ff1660008054906101000a900460ff1660ff16106110de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d59061205f565b60405180910390fd5b816000806101000a81548160ff021916908360ff160217905550600190505b919050565b61110c8282610689565b61119b576111318173ffffffffffffffffffffffffffffffffffffffff1660146111c2565b61113f8360001c60206111c2565b604051602001611150929190612153565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611192919061218d565b60405180910390fd5b5050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6060600060028360026111d591906121de565b6111df9190612238565b67ffffffffffffffff8111156111f8576111f7611640565b5b6040519080825280601f01601f19166020018201604052801561122a5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106112625761126161228e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106112c6576112c561228e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000600184600261130691906121de565b6113109190612238565b90505b60018111156113b0577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106113525761135161228e565b5b1a60f81b8282815181106113695761136861228e565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806113a9906122bd565b9050611313565b50600084146113f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113eb90612333565b60405180910390fd5b8091505092915050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61144781611412565b811461145257600080fd5b50565b6000813590506114648161143e565b92915050565b6000602082840312156114805761147f611408565b5b600061148e84828501611455565b91505092915050565b60008115159050919050565b6114ac81611497565b82525050565b60006020820190506114c760008301846114a3565b92915050565b6000819050919050565b6114e0816114cd565b81146114eb57600080fd5b50565b6000813590506114fd816114d7565b92915050565b60006020828403121561151957611518611408565b5b6000611527848285016114ee565b91505092915050565b611539816114cd565b82525050565b60006020820190506115546000830184611530565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115858261155a565b9050919050565b6115958161157a565b81146115a057600080fd5b50565b6000813590506115b28161158c565b92915050565b6000602082840312156115ce576115cd611408565b5b60006115dc848285016115a3565b91505092915050565b600080604083850312156115fc576115fb611408565b5b600061160a858286016114ee565b925050602061161b858286016115a3565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6116788261162f565b810181811067ffffffffffffffff8211171561169757611696611640565b5b80604052505050565b60006116aa6113fe565b90506116b6828261166f565b919050565b600067ffffffffffffffff8211156116d6576116d5611640565b5b6116df8261162f565b9050602081019050919050565b82818337600083830152505050565b600061170e611709846116bb565b6116a0565b90508281526020810184848401111561172a5761172961162a565b5b6117358482856116ec565b509392505050565b600082601f83011261175257611751611625565b5b81356117628482602086016116fb565b91505092915050565b6000819050919050565b61177e8161176b565b811461178957600080fd5b50565b60008135905061179b81611775565b92915050565b600080600080608085870312156117bb576117ba611408565b5b60006117c9878288016115a3565b945050602085013567ffffffffffffffff8111156117ea576117e961140d565b5b6117f68782880161173d565b93505060406118078782880161178c565b92505060606118188782880161178c565b91505092959194509250565b61182d8161176b565b82525050565b60006020820190506118486000830184611824565b92915050565b60006118598261155a565b9050919050565b6118698161184e565b811461187457600080fd5b50565b60008135905061188681611860565b92915050565b6000806000606084860312156118a5576118a4611408565b5b60006118b3868287016115a3565b93505060206118c486828701611877565b92505060406118d58682870161178c565b9150509250925092565b600080600080600060a086880312156118fb576118fa611408565b5b6000611909888289016115a3565b955050602061191a88828901611877565b945050604061192b8882890161178c565b935050606061193c8882890161178c565b925050608061194d8882890161178c565b9150509295509295909350565b6000602082840312156119705761196f611408565b5b600061197e8482850161178c565b91505092915050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b60006119f4602f83611987565b91506119ff82611998565b604082019050919050565b60006020820190508181036000830152611a23816119e7565b9050919050565b7f43616c6c6572206973206e6f742061206d616e61676572000000000000000000600082015250565b6000611a60601783611987565b9150611a6b82611a2a565b602082019050919050565b60006020820190508181036000830152611a8f81611a53565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000611acc601083611987565b9150611ad782611a96565b602082019050919050565b60006020820190508181036000830152611afb81611abf565b9050919050565b611b0b8161157a565b82525050565b6000602082019050611b266000830184611b02565b92915050565b600081519050611b3b81611775565b92915050565b600060208284031215611b5757611b56611408565b5b6000611b6584828501611b2c565b91505092915050565b7f62616c616e636520746f6f206c6f770000000000000000000000000000000000600082015250565b6000611ba4600f83611987565b9150611baf82611b6e565b602082019050919050565b60006020820190508181036000830152611bd381611b97565b9050919050565b600081519050919050565b60005b83811015611c03578082015181840152602081019050611be8565b83811115611c12576000848401525b50505050565b6000611c2382611bda565b611c2d8185611987565b9350611c3d818560208601611be5565b611c468161162f565b840191505092915050565b60006040820190508181036000830152611c6b8185611c18565b9050611c7a6020830184611824565b9392505050565b6000819050919050565b6000611ca6611ca1611c9c8461155a565b611c81565b61155a565b9050919050565b6000611cb882611c8b565b9050919050565b6000611cca82611cad565b9050919050565b611cda81611cbf565b82525050565b6000604082019050611cf56000830185611cd1565b611d026020830184611824565b9392505050565b611d1281611497565b8114611d1d57600080fd5b50565b600081519050611d2f81611d09565b92915050565b600060208284031215611d4b57611d4a611408565b5b6000611d5984828501611d20565b91505092915050565b6000819050919050565b600060ff82169050919050565b6000611d94611d8f611d8a84611d62565b611c81565b611d6c565b9050919050565b611da481611d79565b82525050565b6000602082019050611dbf6000830184611d9b565b92915050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000611dfb601483611987565b9150611e0682611dc5565b602082019050919050565b60006020820190508181036000830152611e2a81611dee565b9050919050565b6000606082019050611e466000830186611b02565b611e536020830185611b02565b611e606040830184611824565b949350505050565b600081519050919050565b600081905092915050565b6000611e8982611e68565b611e938185611e73565b9350611ea3818560208601611be5565b80840191505092915050565b6000611ebb8284611e7e565b915081905092915050565b7f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f464160008201527f494c454400000000000000000000000000000000000000000000000000000000602082015250565b6000611f22602483611987565b9150611f2d82611ec6565b604082019050919050565b60006020820190508181036000830152611f5181611f15565b9050919050565b6000604082019050611f6d6000830185611b02565b611f7a6020830184611824565b9392505050565b7f5472616e7366657248656c7065723a205452414e534645525f4641494c454400600082015250565b6000611fb7601f83611987565b9150611fc282611f81565b602082019050919050565b60006020820190508181036000830152611fe681611faa565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000612049602e83611987565b915061205482611fed565b604082019050919050565b600060208201905081810360008301526120788161203c565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b60006120c060178361207f565b91506120cb8261208a565b601782019050919050565b60006120e182611bda565b6120eb818561207f565b93506120fb818560208601611be5565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b600061213d60118361207f565b915061214882612107565b601182019050919050565b600061215e826120b3565b915061216a82856120d6565b915061217582612130565b915061218182846120d6565b91508190509392505050565b600060208201905081810360008301526121a78184611c18565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006121e98261176b565b91506121f48361176b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561222d5761222c6121af565b5b828202905092915050565b60006122438261176b565b915061224e8361176b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612283576122826121af565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006122c88261176b565b915060008214156122dc576122db6121af565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b600061231d602083611987565b9150612328826122e7565b602082019050919050565b6000602082019050818103600083015261234c81612310565b905091905056fea2646970667358221220a2c60b40980fab92710500ed42e3726350a069002523c68a551b34803aca2a4b64736f6c634300080b0033",
}

// Bridgev1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgev1MetaData.ABI instead.
var Bridgev1ABI = Bridgev1MetaData.ABI

// Bridgev1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bridgev1MetaData.Bin instead.
var Bridgev1Bin = Bridgev1MetaData.Bin

// DeployBridgev1 deploys a new Ethereum contract, binding an instance of Bridgev1 to it.
func DeployBridgev1(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId *big.Int) (common.Address, *types.Transaction, *Bridgev1, error) {
	parsed, err := Bridgev1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bridgev1Bin), backend, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgev1{Bridgev1Caller: Bridgev1Caller{contract: contract}, Bridgev1Transactor: Bridgev1Transactor{contract: contract}, Bridgev1Filterer: Bridgev1Filterer{contract: contract}}, nil
}

// Bridgev1 is an auto generated Go binding around an Ethereum contract.
type Bridgev1 struct {
	Bridgev1Caller     // Read-only binding to the contract
	Bridgev1Transactor // Write-only binding to the contract
	Bridgev1Filterer   // Log filterer for contract events
}

// Bridgev1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgev1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgev1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgev1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgev1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgev1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgev1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgev1Session struct {
	Contract     *Bridgev1         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Bridgev1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgev1CallerSession struct {
	Contract *Bridgev1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Bridgev1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgev1TransactorSession struct {
	Contract     *Bridgev1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Bridgev1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgev1Raw struct {
	Contract *Bridgev1 // Generic contract binding to access the raw methods on
}

// Bridgev1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgev1CallerRaw struct {
	Contract *Bridgev1Caller // Generic read-only contract binding to access the raw methods on
}

// Bridgev1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgev1TransactorRaw struct {
	Contract *Bridgev1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgev1 creates a new instance of Bridgev1, bound to a specific deployed contract.
func NewBridgev1(address common.Address, backend bind.ContractBackend) (*Bridgev1, error) {
	contract, err := bindBridgev1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgev1{Bridgev1Caller: Bridgev1Caller{contract: contract}, Bridgev1Transactor: Bridgev1Transactor{contract: contract}, Bridgev1Filterer: Bridgev1Filterer{contract: contract}}, nil
}

// NewBridgev1Caller creates a new read-only instance of Bridgev1, bound to a specific deployed contract.
func NewBridgev1Caller(address common.Address, caller bind.ContractCaller) (*Bridgev1Caller, error) {
	contract, err := bindBridgev1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgev1Caller{contract: contract}, nil
}

// NewBridgev1Transactor creates a new write-only instance of Bridgev1, bound to a specific deployed contract.
func NewBridgev1Transactor(address common.Address, transactor bind.ContractTransactor) (*Bridgev1Transactor, error) {
	contract, err := bindBridgev1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgev1Transactor{contract: contract}, nil
}

// NewBridgev1Filterer creates a new log filterer instance of Bridgev1, bound to a specific deployed contract.
func NewBridgev1Filterer(address common.Address, filterer bind.ContractFilterer) (*Bridgev1Filterer, error) {
	contract, err := bindBridgev1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgev1Filterer{contract: contract}, nil
}

// bindBridgev1 binds a generic wrapper to an already deployed contract.
func bindBridgev1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Bridgev1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgev1 *Bridgev1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgev1.Contract.Bridgev1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgev1 *Bridgev1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgev1.Contract.Bridgev1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgev1 *Bridgev1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgev1.Contract.Bridgev1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgev1 *Bridgev1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgev1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgev1 *Bridgev1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgev1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgev1 *Bridgev1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgev1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridgev1.Contract.DEFAULTADMINROLE(&_Bridgev1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridgev1.Contract.DEFAULTADMINROLE(&_Bridgev1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1Caller) MANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1Session) MANAGERROLE() ([32]byte, error) {
	return _Bridgev1.Contract.MANAGERROLE(&_Bridgev1.CallOpts)
}

// MANAGERROLE is a free data retrieval call binding the contract method 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (_Bridgev1 *Bridgev1CallerSession) MANAGERROLE() ([32]byte, error) {
	return _Bridgev1.Contract.MANAGERROLE(&_Bridgev1.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridgev1 *Bridgev1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridgev1 *Bridgev1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridgev1.Contract.GetRoleAdmin(&_Bridgev1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridgev1 *Bridgev1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridgev1.Contract.GetRoleAdmin(&_Bridgev1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridgev1 *Bridgev1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridgev1 *Bridgev1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridgev1.Contract.HasRole(&_Bridgev1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridgev1 *Bridgev1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridgev1.Contract.HasRole(&_Bridgev1.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgev1 *Bridgev1Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgev1 *Bridgev1Session) Paused() (bool, error) {
	return _Bridgev1.Contract.Paused(&_Bridgev1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridgev1 *Bridgev1CallerSession) Paused() (bool, error) {
	return _Bridgev1.Contract.Paused(&_Bridgev1.CallOpts)
}

// SelfChainId is a free data retrieval call binding the contract method 0xcc9e3e89.
//
// Solidity: function selfChainId() view returns(uint256)
func (_Bridgev1 *Bridgev1Caller) SelfChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "selfChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelfChainId is a free data retrieval call binding the contract method 0xcc9e3e89.
//
// Solidity: function selfChainId() view returns(uint256)
func (_Bridgev1 *Bridgev1Session) SelfChainId() (*big.Int, error) {
	return _Bridgev1.Contract.SelfChainId(&_Bridgev1.CallOpts)
}

// SelfChainId is a free data retrieval call binding the contract method 0xcc9e3e89.
//
// Solidity: function selfChainId() view returns(uint256)
func (_Bridgev1 *Bridgev1CallerSession) SelfChainId() (*big.Int, error) {
	return _Bridgev1.Contract.SelfChainId(&_Bridgev1.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridgev1 *Bridgev1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridgev1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridgev1 *Bridgev1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridgev1.Contract.SupportsInterface(&_Bridgev1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridgev1 *Bridgev1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridgev1.Contract.SupportsInterface(&_Bridgev1.CallOpts, interfaceId)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bridgev1 *Bridgev1Transactor) AddManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "addManager", manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bridgev1 *Bridgev1Session) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.AddManager(&_Bridgev1.TransactOpts, manager)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address manager) returns()
func (_Bridgev1 *Bridgev1TransactorSession) AddManager(manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.AddManager(&_Bridgev1.TransactOpts, manager)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.GrantRole(&_Bridgev1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.GrantRole(&_Bridgev1.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Bridgev1 *Bridgev1Transactor) Initialize(opts *bind.TransactOpts, _chainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "initialize", _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Bridgev1 *Bridgev1Session) Initialize(_chainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.Initialize(&_Bridgev1.TransactOpts, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Bridgev1 *Bridgev1TransactorSession) Initialize(_chainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.Initialize(&_Bridgev1.TransactOpts, _chainId)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bridgev1 *Bridgev1Transactor) RemoveManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "removeManager", manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bridgev1 *Bridgev1Session) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RemoveManager(&_Bridgev1.TransactOpts, manager)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address manager) returns()
func (_Bridgev1 *Bridgev1TransactorSession) RemoveManager(manager common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RemoveManager(&_Bridgev1.TransactOpts, manager)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RenounceRole(&_Bridgev1.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RenounceRole(&_Bridgev1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RevokeRole(&_Bridgev1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridgev1 *Bridgev1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridgev1.Contract.RevokeRole(&_Bridgev1.TransactOpts, role, account)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Bridgev1 *Bridgev1Transactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Bridgev1 *Bridgev1Session) SetPause() (*types.Transaction, error) {
	return _Bridgev1.Contract.SetPause(&_Bridgev1.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Bridgev1 *Bridgev1TransactorSession) SetPause() (*types.Transaction, error) {
	return _Bridgev1.Contract.SetPause(&_Bridgev1.TransactOpts)
}

// SetUnpause is a paid mutator transaction binding the contract method 0x848cb5c6.
//
// Solidity: function setUnpause() returns()
func (_Bridgev1 *Bridgev1Transactor) SetUnpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "setUnpause")
}

// SetUnpause is a paid mutator transaction binding the contract method 0x848cb5c6.
//
// Solidity: function setUnpause() returns()
func (_Bridgev1 *Bridgev1Session) SetUnpause() (*types.Transaction, error) {
	return _Bridgev1.Contract.SetUnpause(&_Bridgev1.TransactOpts)
}

// SetUnpause is a paid mutator transaction binding the contract method 0x848cb5c6.
//
// Solidity: function setUnpause() returns()
func (_Bridgev1 *Bridgev1TransactorSession) SetUnpause() (*types.Transaction, error) {
	return _Bridgev1.Contract.SetUnpause(&_Bridgev1.TransactOpts)
}

// TransferIn is a paid mutator transaction binding the contract method 0xef146cf6.
//
// Solidity: function transferIn(address _token, address _to, uint256 _amount, uint256 _fromChain, uint256 _toChain) returns()
func (_Bridgev1 *Bridgev1Transactor) TransferIn(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _fromChain *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "transferIn", _token, _to, _amount, _fromChain, _toChain)
}

// TransferIn is a paid mutator transaction binding the contract method 0xef146cf6.
//
// Solidity: function transferIn(address _token, address _to, uint256 _amount, uint256 _fromChain, uint256 _toChain) returns()
func (_Bridgev1 *Bridgev1Session) TransferIn(_token common.Address, _to common.Address, _amount *big.Int, _fromChain *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.TransferIn(&_Bridgev1.TransactOpts, _token, _to, _amount, _fromChain, _toChain)
}

// TransferIn is a paid mutator transaction binding the contract method 0xef146cf6.
//
// Solidity: function transferIn(address _token, address _to, uint256 _amount, uint256 _fromChain, uint256 _toChain) returns()
func (_Bridgev1 *Bridgev1TransactorSession) TransferIn(_token common.Address, _to common.Address, _amount *big.Int, _fromChain *big.Int, _toChain *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.TransferIn(&_Bridgev1.TransactOpts, _token, _to, _amount, _fromChain, _toChain)
}

// TransferOut is a paid mutator transaction binding the contract method 0x8c7cb280.
//
// Solidity: function transferOut(address _token, string _receiver, uint256 _amount, uint256 _toChainId) returns()
func (_Bridgev1 *Bridgev1Transactor) TransferOut(opts *bind.TransactOpts, _token common.Address, _receiver string, _amount *big.Int, _toChainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "transferOut", _token, _receiver, _amount, _toChainId)
}

// TransferOut is a paid mutator transaction binding the contract method 0x8c7cb280.
//
// Solidity: function transferOut(address _token, string _receiver, uint256 _amount, uint256 _toChainId) returns()
func (_Bridgev1 *Bridgev1Session) TransferOut(_token common.Address, _receiver string, _amount *big.Int, _toChainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.TransferOut(&_Bridgev1.TransactOpts, _token, _receiver, _amount, _toChainId)
}

// TransferOut is a paid mutator transaction binding the contract method 0x8c7cb280.
//
// Solidity: function transferOut(address _token, string _receiver, uint256 _amount, uint256 _toChainId) returns()
func (_Bridgev1 *Bridgev1TransactorSession) TransferOut(_token common.Address, _receiver string, _amount *big.Int, _toChainId *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.TransferOut(&_Bridgev1.TransactOpts, _token, _receiver, _amount, _toChainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _receiver, uint256 _amount) returns()
func (_Bridgev1 *Bridgev1Transactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgev1.contract.Transact(opts, "withdraw", _token, _receiver, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _receiver, uint256 _amount) returns()
func (_Bridgev1 *Bridgev1Session) Withdraw(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.Withdraw(&_Bridgev1.TransactOpts, _token, _receiver, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _receiver, uint256 _amount) returns()
func (_Bridgev1 *Bridgev1TransactorSession) Withdraw(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridgev1.Contract.Withdraw(&_Bridgev1.TransactOpts, _token, _receiver, _amount)
}

// Bridgev1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgev1 contract.
type Bridgev1InitializedIterator struct {
	Event *Bridgev1Initialized // Event containing the contract specifics and raw log

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
func (it *Bridgev1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1Initialized)
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
		it.Event = new(Bridgev1Initialized)
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
func (it *Bridgev1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1Initialized represents a Initialized event raised by the Bridgev1 contract.
type Bridgev1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgev1 *Bridgev1Filterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgev1InitializedIterator, error) {

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgev1InitializedIterator{contract: _Bridgev1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgev1 *Bridgev1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgev1Initialized) (event.Subscription, error) {

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1Initialized)
				if err := _Bridgev1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgev1 *Bridgev1Filterer) ParseInitialized(log types.Log) (*Bridgev1Initialized, error) {
	event := new(Bridgev1Initialized)
	if err := _Bridgev1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1MAPTransferInIterator is returned from FilterMAPTransferIn and is used to iterate over the raw logs and unpacked data for MAPTransferIn events raised by the Bridgev1 contract.
type Bridgev1MAPTransferInIterator struct {
	Event *Bridgev1MAPTransferIn // Event containing the contract specifics and raw log

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
func (it *Bridgev1MAPTransferInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1MAPTransferIn)
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
		it.Event = new(Bridgev1MAPTransferIn)
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
func (it *Bridgev1MAPTransferInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1MAPTransferInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1MAPTransferIn represents a MAPTransferIn event raised by the Bridgev1 contract.
type Bridgev1MAPTransferIn struct {
	FromChain *big.Int
	ToChain   *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMAPTransferIn is a free log retrieval operation binding the contract event 0x480c1cd6cfcaa4162f54fd97792200e72dbb4bd1a21259f882414f0d9188941d.
//
// Solidity: event MAPTransferIn(uint256 indexed fromChain, uint256 indexed toChain, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) FilterMAPTransferIn(opts *bind.FilterOpts, fromChain []*big.Int, toChain []*big.Int) (*Bridgev1MAPTransferInIterator, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "MAPTransferIn", fromChainRule, toChainRule)
	if err != nil {
		return nil, err
	}
	return &Bridgev1MAPTransferInIterator{contract: _Bridgev1.contract, event: "MAPTransferIn", logs: logs, sub: sub}, nil
}

// WatchMAPTransferIn is a free log subscription operation binding the contract event 0x480c1cd6cfcaa4162f54fd97792200e72dbb4bd1a21259f882414f0d9188941d.
//
// Solidity: event MAPTransferIn(uint256 indexed fromChain, uint256 indexed toChain, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) WatchMAPTransferIn(opts *bind.WatchOpts, sink chan<- *Bridgev1MAPTransferIn, fromChain []*big.Int, toChain []*big.Int) (event.Subscription, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "MAPTransferIn", fromChainRule, toChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1MAPTransferIn)
				if err := _Bridgev1.contract.UnpackLog(event, "MAPTransferIn", log); err != nil {
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

// ParseMAPTransferIn is a log parse operation binding the contract event 0x480c1cd6cfcaa4162f54fd97792200e72dbb4bd1a21259f882414f0d9188941d.
//
// Solidity: event MAPTransferIn(uint256 indexed fromChain, uint256 indexed toChain, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) ParseMAPTransferIn(log types.Log) (*Bridgev1MAPTransferIn, error) {
	event := new(Bridgev1MAPTransferIn)
	if err := _Bridgev1.contract.UnpackLog(event, "MAPTransferIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1MAPTransferOutIterator is returned from FilterMAPTransferOut and is used to iterate over the raw logs and unpacked data for MAPTransferOut events raised by the Bridgev1 contract.
type Bridgev1MAPTransferOutIterator struct {
	Event *Bridgev1MAPTransferOut // Event containing the contract specifics and raw log

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
func (it *Bridgev1MAPTransferOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1MAPTransferOut)
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
		it.Event = new(Bridgev1MAPTransferOut)
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
func (it *Bridgev1MAPTransferOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1MAPTransferOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1MAPTransferOut represents a MAPTransferOut event raised by the Bridgev1 contract.
type Bridgev1MAPTransferOut struct {
	FromChain *big.Int
	ToChain   *big.Int
	To        string
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMAPTransferOut is a free log retrieval operation binding the contract event 0xc2cf97cb20e3fec72bbd1515ad10ad56582e8085858049990a34a1abe8fbee3b.
//
// Solidity: event MAPTransferOut(uint256 indexed fromChain, uint256 indexed toChain, string to, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) FilterMAPTransferOut(opts *bind.FilterOpts, fromChain []*big.Int, toChain []*big.Int) (*Bridgev1MAPTransferOutIterator, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "MAPTransferOut", fromChainRule, toChainRule)
	if err != nil {
		return nil, err
	}
	return &Bridgev1MAPTransferOutIterator{contract: _Bridgev1.contract, event: "MAPTransferOut", logs: logs, sub: sub}, nil
}

// WatchMAPTransferOut is a free log subscription operation binding the contract event 0xc2cf97cb20e3fec72bbd1515ad10ad56582e8085858049990a34a1abe8fbee3b.
//
// Solidity: event MAPTransferOut(uint256 indexed fromChain, uint256 indexed toChain, string to, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) WatchMAPTransferOut(opts *bind.WatchOpts, sink chan<- *Bridgev1MAPTransferOut, fromChain []*big.Int, toChain []*big.Int) (event.Subscription, error) {

	var fromChainRule []interface{}
	for _, fromChainItem := range fromChain {
		fromChainRule = append(fromChainRule, fromChainItem)
	}
	var toChainRule []interface{}
	for _, toChainItem := range toChain {
		toChainRule = append(toChainRule, toChainItem)
	}

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "MAPTransferOut", fromChainRule, toChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1MAPTransferOut)
				if err := _Bridgev1.contract.UnpackLog(event, "MAPTransferOut", log); err != nil {
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

// ParseMAPTransferOut is a log parse operation binding the contract event 0xc2cf97cb20e3fec72bbd1515ad10ad56582e8085858049990a34a1abe8fbee3b.
//
// Solidity: event MAPTransferOut(uint256 indexed fromChain, uint256 indexed toChain, string to, uint256 amount)
func (_Bridgev1 *Bridgev1Filterer) ParseMAPTransferOut(log types.Log) (*Bridgev1MAPTransferOut, error) {
	event := new(Bridgev1MAPTransferOut)
	if err := _Bridgev1.contract.UnpackLog(event, "MAPTransferOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridgev1 contract.
type Bridgev1PausedIterator struct {
	Event *Bridgev1Paused // Event containing the contract specifics and raw log

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
func (it *Bridgev1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1Paused)
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
		it.Event = new(Bridgev1Paused)
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
func (it *Bridgev1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1Paused represents a Paused event raised by the Bridgev1 contract.
type Bridgev1Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridgev1 *Bridgev1Filterer) FilterPaused(opts *bind.FilterOpts) (*Bridgev1PausedIterator, error) {

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Bridgev1PausedIterator{contract: _Bridgev1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridgev1 *Bridgev1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Bridgev1Paused) (event.Subscription, error) {

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1Paused)
				if err := _Bridgev1.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridgev1 *Bridgev1Filterer) ParsePaused(log types.Log) (*Bridgev1Paused, error) {
	event := new(Bridgev1Paused)
	if err := _Bridgev1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridgev1 contract.
type Bridgev1RoleAdminChangedIterator struct {
	Event *Bridgev1RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Bridgev1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1RoleAdminChanged)
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
		it.Event = new(Bridgev1RoleAdminChanged)
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
func (it *Bridgev1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1RoleAdminChanged represents a RoleAdminChanged event raised by the Bridgev1 contract.
type Bridgev1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridgev1 *Bridgev1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Bridgev1RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Bridgev1RoleAdminChangedIterator{contract: _Bridgev1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridgev1 *Bridgev1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Bridgev1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1RoleAdminChanged)
				if err := _Bridgev1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridgev1 *Bridgev1Filterer) ParseRoleAdminChanged(log types.Log) (*Bridgev1RoleAdminChanged, error) {
	event := new(Bridgev1RoleAdminChanged)
	if err := _Bridgev1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridgev1 contract.
type Bridgev1RoleGrantedIterator struct {
	Event *Bridgev1RoleGranted // Event containing the contract specifics and raw log

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
func (it *Bridgev1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1RoleGranted)
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
		it.Event = new(Bridgev1RoleGranted)
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
func (it *Bridgev1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1RoleGranted represents a RoleGranted event raised by the Bridgev1 contract.
type Bridgev1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bridgev1RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bridgev1RoleGrantedIterator{contract: _Bridgev1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Bridgev1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1RoleGranted)
				if err := _Bridgev1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) ParseRoleGranted(log types.Log) (*Bridgev1RoleGranted, error) {
	event := new(Bridgev1RoleGranted)
	if err := _Bridgev1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridgev1 contract.
type Bridgev1RoleRevokedIterator struct {
	Event *Bridgev1RoleRevoked // Event containing the contract specifics and raw log

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
func (it *Bridgev1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1RoleRevoked)
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
		it.Event = new(Bridgev1RoleRevoked)
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
func (it *Bridgev1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1RoleRevoked represents a RoleRevoked event raised by the Bridgev1 contract.
type Bridgev1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Bridgev1RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Bridgev1RoleRevokedIterator{contract: _Bridgev1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Bridgev1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1RoleRevoked)
				if err := _Bridgev1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridgev1 *Bridgev1Filterer) ParseRoleRevoked(log types.Log) (*Bridgev1RoleRevoked, error) {
	event := new(Bridgev1RoleRevoked)
	if err := _Bridgev1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgev1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridgev1 contract.
type Bridgev1UnpausedIterator struct {
	Event *Bridgev1Unpaused // Event containing the contract specifics and raw log

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
func (it *Bridgev1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgev1Unpaused)
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
		it.Event = new(Bridgev1Unpaused)
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
func (it *Bridgev1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgev1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgev1Unpaused represents a Unpaused event raised by the Bridgev1 contract.
type Bridgev1Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridgev1 *Bridgev1Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Bridgev1UnpausedIterator, error) {

	logs, sub, err := _Bridgev1.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Bridgev1UnpausedIterator{contract: _Bridgev1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridgev1 *Bridgev1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Bridgev1Unpaused) (event.Subscription, error) {

	logs, sub, err := _Bridgev1.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgev1Unpaused)
				if err := _Bridgev1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridgev1 *Bridgev1Filterer) ParseUnpaused(log types.Log) (*Bridgev1Unpaused, error) {
	event := new(Bridgev1Unpaused)
	if err := _Bridgev1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
