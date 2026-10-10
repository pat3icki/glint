// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package glintdough

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

// GlintDoughTokenData is an auto generated low-level Go binding around an user-defined struct.
type GlintDoughTokenData struct {
	Decimals      uint8
	SmartContract common.Address
	Symbol        string
	IsActive      bool
}

// GlintdoughMetaData contains all meta data concerning the Glintdough contract.
var GlintdoughMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creationfeesAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feesAddress\",\"type\":\"address\"}],\"name\":\"FeesAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterRoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"checksum\",\"type\":\"bytes32\"}],\"name\":\"PhysicalCardRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"TokenStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"denomination\",\"type\":\"uint256\"}],\"name\":\"VoucherAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"VoucherMintedSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"VoucherRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"VoucherToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"VouchersMintedBatch\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"activateCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"denomination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationFees\",\"type\":\"uint256\"}],\"name\":\"addVoucher\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"collateralTokens\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"deactivateCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"}],\"name\":\"existsByAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structGlintDough.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"}],\"name\":\"getTokenByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"smartContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structGlintDough.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"grantMinterRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"redeemTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeesAddress\",\"type\":\"address\"}],\"name\":\"setFeesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"toggleVoucher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vouchers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collateral_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"denomination\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationFees\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"is_active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600855348015610014575f5ffd5b50604051613a6c380380613a6c8339810160408190526100339161017f565b8161003d81610093565b5060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005561006c5f336100a3565b50600780546001600160a01b0319166001600160a01b03929092169190911790555061037e565b600261009f82826102c4565b5050565b5f8281526003602090815260408083206001600160a01b038516845290915281205460ff16610147575f8381526003602090815260408083206001600160a01b03861684529091529020805460ff191660011790556100ff3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161014a565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b038116811461017a575f5ffd5b919050565b5f5f60408385031215610190575f5ffd5b82516001600160401b038111156101a5575f5ffd5b8301601f810185136101b5575f5ffd5b80516001600160401b038111156101ce576101ce610150565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101fc576101fc610150565b604052818152828201602001871015610213575f5ffd5b8160208401602083015e5f6020838301015280945050505061023760208401610164565b90509250929050565b600181811c9082168061025457607f821691505b60208210810361027257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102bf57805f5260205f20601f840160051c8101602085101561029d5750805b601f840160051c820191505b818110156102bc575f81556001016102a9565b50505b505050565b81516001600160401b038111156102dd576102dd610150565b6102f1816102eb8454610240565b84610278565b6020601f821160018114610323575f831561030c5750848201515b5f19600385901b1c1916600184901b1784556102bc565b5f84815260208120601f198516915b828110156103525787850151825560209485019460019092019101610332565b508482101561036f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6136e18061038b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610212575f3560e01c8063736e4f571161011f578063bdbacb95116100a9578063d547741f11610079578063d547741f146104ac578063e0b31602146104bf578063e985e9c5146104d2578063f08f97961461050d578063f242432a14610520575f5ffd5b8063bdbacb951461045f578063c1733f6814610472578063cf9e099c14610485578063d539139314610498575f5ffd5b806391d14854116100ef57806391d14854146103ff57806391ded8fa1461041257806394a94f3114610432578063a217fddf14610445578063a22cb4651461044c575f5ffd5b8063736e4f57146103c857806375794a3c146103db5780637cbc2373146103e45780638456cb59146103f7575f5ffd5b806336568abe116101a057806346bcc9531161017057806346bcc953146103645780634dd91ff9146103775780634e1273f41461038a5780635c975abb146103aa5780635ca5933a146103b5575f5ffd5b806336568abe146103125780633dd1eb61146103255780633e97b43f146103385780633f4ba83a1461035c575f5ffd5b80630e89341c116101e65780630e89341c14610297578063248a9ca3146102b7578063261a323e146102d95780632eb2c2d6146102ec5780632f2ff15d146102ff575f5ffd5b8062fdd58e1461021657806301ffc9a71461023c57806308983f6d1461025f5780630d398b5a14610274575b5f5ffd5b610229610224366004612b57565b610533565b6040519081526020015b60405180910390f35b61024f61024a366004612b94565b61055a565b6040519015158152602001610233565b61027261026d366004612c64565b610564565b005b610287610282366004612c64565b610670565b6040516102339493929190612ccb565b6102aa6102a5366004612d0a565b61073b565b6040516102339190612d21565b6102296102c5366004612d0a565b5f9081526003602052604090206001015490565b61024f6102e7366004612c64565b6107cd565b6102726102fa366004612dbf565b61080f565b61027261030d366004612e6b565b61082e565b610272610320366004612e6b565b610858565b610272610333366004612e95565b610890565b61034b610346366004612d0a565b6108f2565b604051610233959493929190612eae565b6102726109b3565b610272610372366004612eef565b6109c8565b610229610385366004612f22565b6109f2565b61039d610398366004612f6b565b610cbf565b6040516102339190613066565b60045460ff1661024f565b6102726103c3366004613078565b610d89565b6102726103d6366004612e95565b611082565b61022960085481565b6102726103f23660046130d4565b611120565b61027261114e565b61024f61040d366004612e6b565b611160565b610425610420366004612e95565b61118a565b60405161023391906130f4565b61024f610440366004612e95565b61136b565b6102295f81565b61027261045a366004613155565b611396565b61027261046d36600461317d565b6113a1565b610425610480366004612c64565b611627565b610272610493366004613207565b611761565b6102295f51602061366c5f395f51905f5281565b6102726104ba366004612e6b565b611b2e565b6102726104cd366004612c64565b611b52565b61024f6104e0366004613271565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205460ff1690565b61027261051b366004613299565b611c30565b61027261052e3660046132ba565b611ce8565b5f818152602081815260408083206001600160a01b03861684529091529020545b92915050565b5f61055482611cfe565b5f61056e81611d22565b5f60058360405161057f919061330d565b908152604051908190036020019020805490915061010090046001600160a01b03166105c65760405162461bcd60e51b81526004016105bd90613323565b60405180910390fd5b600281015460ff16156106135760405162461bcd60e51b8152602060048201526015602482015274546f6b656e3a20616c72656164792061637469766560581b60448201526064016105bd565b60028101805460ff1916600117905560405161063090849061330d565b60405190819003812060018252907f41cda4530bfc26ff01b587bd9fc0bcc8c67aa176a168da1542f9dd2f333be8a29060200160405180910390a2505050565b80516020818301810180516005825292820191909301209152805460018201805460ff8316936101009093046001600160a01b03169291906106b19061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546106dd9061334d565b80156107285780601f106106ff57610100808354040283529160200191610728565b820191905f5260205f20905b81548152906001019060200180831161070b57829003601f168201915b5050506002909301549192505060ff1684565b60606002805461074a9061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546107769061334d565b80156107c15780601f10610798576101008083540402835291602001916107c1565b820191905f5260205f20905b8154815290600101906020018083116107a457829003601f168201915b50505050509050919050565b5f5f6001600160a01b03166005836040516107e8919061330d565b908152604051908190036020019020546001600160a01b0361010090910416141592915050565b61081a335b86611d2c565b6108278585858585611da4565b5050505050565b5f8281526003602052604090206001015461084881611d22565b6108528383611e04565b50505050565b6001600160a01b03811633146108815760405163334bd91960e11b815260040160405180910390fd5b61088b8282611e95565b505050565b5f61089a81611d22565b6108b15f51602061366c5f395f51905f5283611e04565b506040516001600160a01b03831681527fbd42b6cd963a984d8cbff85acad4dd2c87a6d85ac756d16e6d2ec5ce6c8f07309060200160405180910390a15050565b60096020525f908152604090208054819061090c9061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546109389061334d565b80156109835780601f1061095a57610100808354040283529160200191610983565b820191905f5260205f20905b81548152906001019060200180831161096657829003601f168201915b5050505060018301546002840154600385015460049095015493946001600160a01b039092169390925060ff1685565b5f6109bd81611d22565b6109c5611f00565b50565b6109d0611f52565b6109dc33848484611f6d565b61088b60015f51602061368c5f395f51905f5255565b5f806109fd81611d22565b5f8411610a565760405162461bcd60e51b815260206004820152602160248201527f44656e6f6d696e6174696f6e206d7573742062652067726561746572207a65726044820152606f60f81b60648201526084016105bd565b5f600586604051610a67919061330d565b9081526040805160209281900383018120608082018352805460ff8116835261010090046001600160a01b031693820193909352600183018054919392840191610ab09061334d565b80601f0160208091040260200160405190810160405280929190818152602001828054610adc9061334d565b8015610b275780601f10610afe57610100808354040283529160200191610b27565b820191905f5260205f20905b815481529060010190602001808311610b0a57829003601f168201915b50505091835250506002919091015460ff1615156020918201528101519091506001600160a01b0316610b6c5760405162461bcd60e51b81526004016105bd90613323565b8060600151610bbd5760405162461bcd60e51b815260206004820152601f60248201527f436f6c6c61746572616c20746f6b656e2069732064656163746976617465640060448201526064016105bd565b600880545f9182610bcd83613399565b909155506040805160a0810182528981526020858101516001600160a01b0316818301528183018a905260608201899052600160808301525f848152600990915291909120815192935090918190610c2590826133f5565b506020828101516001830180546001600160a01b0319166001600160a01b039283161790556040808501516002850155606085015160038501556080909401516004909301805460ff191693151593909317909255848101519251898152929091169183917f7b8024ce6a9ab55c17f27580f8dc379c6d6beee58d682c382fa50ae828925f74910160405180910390a39695505050505050565b60608151835114610cf05781518351604051635b05999160e01b8152600481019290925260248201526044016105bd565b5f83516001600160401b03811115610d0a57610d0a612baf565b604051908082528060200260200182016040528015610d33578160200160208202803683370190505b5090505f5b8451811015610d8157602080820286010151610d5c90602080840287010151610533565b828281518110610d6e57610d6e6134af565b6020908102919091010152600101610d38565b509392505050565b5f51602061366c5f395f51905f52610da081611d22565b610da8612209565b610db0611f52565b5f85118015610dc0575060085485105b610dff5760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a5908151bdad95b88125160821b60448201526064016105bd565b5f8411610e1e5760405162461bcd60e51b81526004016105bd906134c3565b6001600160a01b038316610e655760405162461bcd60e51b815260206004820152600e60248201526d16995c9bc81c9958da5c1a595b9d60921b60448201526064016105bd565b5f858152600960205260409020600481015460ff16610e965760405162461bcd60e51b81526004016105bd906134f1565b5f858260020154610ea7919061351d565b90505f868360030154610eba919061351d565b60018401549091506001600160a01b0316610ed58284613534565b604051636eb1769f60e11b81523360048201523060248201526001600160a01b0383169063dd62ed3e90604401602060405180830381865afa158015610f1d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f419190613547565b1015610f885760405162461bcd60e51b8152602060048201526016602482015275496e73756666696369656e7420616c6c6f77616e636560501b60448201526064016105bd565b8115610ff9576007546001600160a01b0316610fdd5760405162461bcd60e51b81526020600482015260146024820152731199595cc81859191c995cdcc81b9bdd081cd95d60621b60448201526064016105bd565b600754610ff9906001600160a01b03838116913391168561222f565b61100e6001600160a01b03821633308661222f565b61101a878a8a89612265565b866001600160a01b031689336001600160a01b03167fdb940767cbdf4dfef9c754bd48608769ec95df94cd670e0399e9523b0d52239b8b60405161106091815260200190565b60405180910390a45050505061082760015f51602061368c5f395f51905f5255565b5f61108c81611d22565b6001600160a01b0382166110d65760405162461bcd60e51b81526020600482015260116024820152705a65726f2066656573206164647265737360781b60448201526064016105bd565b600780546001600160a01b0319166001600160a01b0384169081179091556040517f3c7fda2a649f437dd79167a620fa573f40c14c0c0204f16b578cde06433bbb14905f90a25050565b611128611f52565b61113433833384611f6d565b61114a60015f51602061368c5f395f51905f5255565b5050565b5f61115881611d22565b6109c56122c9565b5f9182526003602090815260408084206001600160a01b0393909316845291905290205460ff1690565b604080516080810182525f808252602082018190526060928201839052918101919091526001600160a01b0382165f90815260066020526040812080546111d09061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546111fc9061334d565b80156112475780601f1061121e57610100808354040283529160200191611247565b820191905f5260205f20905b81548152906001019060200180831161122a57829003601f168201915b505050505090505f60058260405161125f919061330d565b9081526040805160209281900383018120608082018352805460ff8116835261010090046001600160a01b0316938201939093526001830180549193928401916112a89061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546112d49061334d565b801561131f5780601f106112f65761010080835404028352916020019161131f565b820191905f5260205f20905b81548152906001019060200180831161130257829003601f168201915b50505091835250506002919091015460ff1615156020918201528101519091506001600160a01b03166113645760405162461bcd60e51b81526004016105bd90613323565b9392505050565b6001600160a01b0381165f908152600660205260408120805461138d9061334d565b15159392505050565b61114a338383612306565b5f6113ab81611d22565b6001600160a01b0384166113fa5760405162461bcd60e51b8152602060048201526016602482015275546f6b656e3a20696e76616c6964206164647265737360501b60448201526064016105bd565b5f6001600160a01b0316600586604051611414919061330d565b908152604051908190036020019020546001600160a01b0361010090910416146114805760405162461bcd60e51b815260206004820152601960248201527f546f6b656e3a206b657920616c7265616479206578697374730000000000000060448201526064016105bd565b6001600160a01b0384165f90815260066020526040902080546114a29061334d565b1590506114fb5760405162461bcd60e51b815260206004820152602160248201527f546f6b656e3a206164647265737320616c7265616479207265676973746572656044820152601960fa1b60648201526084016105bd565b6040805160808101825260ff841681526001600160a01b038616602082015280820185905260016060820152905160059061153790889061330d565b9081526040805191829003602090810190922083518154938501516001600160a01b0316610100026001600160a81b031990941660ff90911617929092178255820151600182019061158990826133f5565b50606091909101516002909101805460ff19169115159190911790556001600160a01b0384165f9081526006602052604090206115c686826133f5565b50836001600160a01b0316856040516115df919061330d565b60405180910390207f03fcf4176a3f2bdb8a0e11fc763d1b887586774f6f6d3e87506e3aa3c20e8e31858560405161161892919061355e565b60405180910390a35050505050565b604080516080810182525f808252602082018190526060928201839052918101919091525f60058360405161165c919061330d565b9081526040805160209281900383018120608082018352805460ff8116835261010090046001600160a01b0316938201939093526001830180549193928401916116a59061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546116d19061334d565b801561171c5780601f106116f35761010080835404028352916020019161171c565b820191905f5260205f20905b8154815290600101906020018083116116ff57829003601f168201915b50505091835250506002919091015460ff1615156020918201528101519091506001600160a01b03166105545760405162461bcd60e51b81526004016105bd90613323565b5f51602061366c5f395f51905f5261177881611d22565b611780612209565b611788611f52565b83518551146117cb5760405162461bcd60e51b815260206004820152600f60248201526e098cadccee8d040dad2e6dac2e8c6d608b1b60448201526064016105bd565b5f8551116118095760405162461bcd60e51b815260206004820152600b60248201526a08adae0e8f240c4c2e8c6d60ab1b60448201526064016105bd565b6001600160a01b0383166118505760405162461bcd60e51b815260206004820152600e60248201526d16995c9bc81c9958da5c1a595b9d60921b60448201526064016105bd565b5f5b8551811015611abe575f86828151811061186e5761186e6134af565b602002602001015190505f86838151811061188b5761188b6134af565b602002602001015190505f821180156118a5575060085482105b6118e45760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a5908151bdad95b88125160821b60448201526064016105bd565b5f81116119035760405162461bcd60e51b81526004016105bd906134c3565b5f828152600960205260409020600481015460ff166119345760405162461bcd60e51b81526004016105bd906134f1565b5f828260020154611945919061351d565b90505f838360030154611958919061351d565b60018401549091506001600160a01b03166119738284613534565b604051636eb1769f60e11b81523360048201523060248201526001600160a01b0383169063dd62ed3e90604401602060405180830381865afa1580156119bb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119df9190613547565b1015611a265760405162461bcd60e51b8152602060048201526016602482015275496e73756666696369656e7420616c6c6f77616e636560501b60448201526064016105bd565b8115611a97576007546001600160a01b0316611a7b5760405162461bcd60e51b81526020600482015260146024820152731199595cc81859191c995cdcc81b9bdd081cd95d60621b60448201526064016105bd565b600754611a97906001600160a01b03838116913391168561222f565b611aac6001600160a01b03821633308661222f565b50506001909401935061185292505050565b50611acb838686856123c3565b826001600160a01b0316336001600160a01b03167f433b9aa33bcc3cdcfd6a434c1936e7034c1d70a0bd333d45aad92d5f423c07e08787604051611b10929190613582565b60405180910390a361082760015f51602061368c5f395f51905f5255565b5f82815260036020526040902060010154611b4881611d22565b6108528383611e95565b5f611b5c81611d22565b5f6001600160a01b0316600583604051611b76919061330d565b908152604051908190036020019020546001600160a01b036101009091041603611bb25760405162461bcd60e51b81526004016105bd90613323565b5f600583604051611bc3919061330d565b908152604051908190036020018120600201805492151560ff1990931692909217909155611bf290839061330d565b6040519081900381205f8252907f41cda4530bfc26ff01b587bd9fc0bcc8c67aa176a168da1542f9dd2f333be8a29060200160405180910390a25050565b5f611c3a81611d22565b5f83815260096020526040902060018101546001600160a01b0316611c9a5760405162461bcd60e51b8152602060048201526016602482015275159bdd58da195c88191bd95cc81b9bdd08195e1a5cdd60521b60448201526064016105bd565b60048101805460ff191684151590811790915560405190815284907f44850097d5058840642ec07d1f10e064b3a8714a43b01f0a9a01a15fc577deb69060200160405180910390a250505050565b611cf133610814565b61082785858585856123fb565b5f6001600160e01b03198216637965db0b60e01b1480610554575061055482612488565b6109c581336124d7565b816001600160a01b0316816001600160a01b031614158015611d7357506001600160a01b038082165f9081526001602090815260408083209386168352929052205460ff16155b1561114a5760405163711bec9160e11b81526001600160a01b038084166004830152821660248201526044016105bd565b6001600160a01b038416611dcd57604051632bfa23e760e11b81525f60048201526024016105bd565b6001600160a01b038516611df557604051626a0d4560e21b81525f60048201526024016105bd565b61082785858585856001612510565b5f611e0f8383611160565b611e8e575f8381526003602090815260408083206001600160a01b03861684529091529020805460ff19166001179055611e463390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610554565b505f610554565b5f611ea08383611160565b15611e8e575f8381526003602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610554565b611f08612568565b6004805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b611f5a61258b565b60025f51602061368c5f395f51905f5255565b5f8111611f8c5760405162461bcd60e51b81526004016105bd906134c3565b6001600160a01b038216611fe25760405162461bcd60e51b815260206004820152601d60248201527f43616e6e6f742072656465656d20746f207a65726f206164647265737300000060448201526064016105bd565b80611fed8585610533565b10156120325760405162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b60448201526064016105bd565b5f83815260096020526040808220815160a081019092528054829082906120589061334d565b80601f01602080910402602001604051908101604052809291908181526020018280546120849061334d565b80156120cf5780601f106120a6576101008083540402835291602001916120cf565b820191905f5260205f20905b8154815290600101906020018083116120b257829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282015460408201526003820154606082015260049091015460ff1615156080918201528101519091506121315760405162461bcd60e51b81526004016105bd906134f1565b60208101516001600160a01b03166121825760405162461bcd60e51b8152602060048201526014602482015273496e76616c696420766f7563686572207479706560601b60448201526064016105bd565b5f828260400151612193919061351d565b90506121a08686856125ba565b60208201516121b9906001600160a01b03168583612621565b604080516001600160a01b0386811682526020820186905288169187917fea6258ffcc941e51a8b38b05fdbef327d28c7015ac1edb675e13086a0439f222910160405180910390a3505050505050565b60045460ff161561222d5760405163d93c066560e01b815260040160405180910390fd5b565b61223d848484846001612656565b61085257604051635274afe760e01b81526001600160a01b03851660048201526024016105bd565b6001600160a01b03841661228e57604051632bfa23e760e11b81525f60048201526024016105bd565b604080516001808252602082018690528183019081526060820185905260808201909252906122c15f8784848784612510565b505050505050565b6122d1612209565b6004805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611f353390565b6001600160a01b03831661232f57604051631f18c42760e11b81525f60048201526024016105bd565b6001600160a01b0382166123575760405162ced3e160e81b81525f60048201526024016105bd565b6001600160a01b038381165f81815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0384166123ec57604051632bfa23e760e11b81525f60048201526024016105bd565b6108525f858585856001612510565b6001600160a01b03841661242457604051632bfa23e760e11b81525f60048201526024016105bd565b6001600160a01b03851661244c57604051626a0d4560e21b81525f60048201526024016105bd565b6040805160018082526020820186905281830190815260608201859052608082019092529061247f87878484875f612510565b50505050505050565b5f6001600160e01b03198216636cdb3d1360e11b14806124b857506001600160e01b031982166303a24d0760e21b145b8061055457506301ffc9a760e01b6001600160e01b0319831614610554565b6124e18282611160565b61114a5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016105bd565b61251c868686866126c3565b6001600160a01b038516156122c157338115612545576125408188888888886128d2565b61247f565b6020858101519085015161255d838a8a85858a6129f3565b505050505050505050565b60045460ff1661222d57604051638dfc202b60e01b815260040160405180910390fd5b5f51602061368c5f395f51905f525460020361222d57604051633ee5aeb560e01b815260040160405180910390fd5b6001600160a01b0383166125e257604051626a0d4560e21b81525f60048201526024016105bd565b604080516001808252602082018590528183019081526060820184905260a082019092525f608082018181529192916108279187918590859083612510565b61262e8383836001612ada565b61088b57604051635274afe760e01b81526001600160a01b03841660048201526024016105bd565b6040516323b872dd60e01b5f8181526001600160a01b038781166004528616602452604485905291602083606481808c5af1925060015f511483166126b25783831516156126a6573d5f823e3d81fd5b5f883b113d1516831692505b604052505f60605295945050505050565b80518251146126f25781518151604051635b05999160e01b8152600481019290925260248201526044016105bd565b335f5b83518110156127f4576020818102858101820151908501909101516001600160a01b038816156127a6575f828152602081815260408083206001600160a01b038c16845290915290205481811015612780576040516303dee4c560e01b81526001600160a01b038a1660048201526024810182905260448101839052606481018490526084016105bd565b5f838152602081815260408083206001600160a01b038d16845290915290209082900390555b6001600160a01b038716156127ea575f828152602081815260408083206001600160a01b038b168452909152812080548392906127e4908490613534565b90915550505b50506001016126f5565b5082516001036128745760208301515f906020840151909150856001600160a01b0316876001600160a01b0316846001600160a01b03167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051612865929190918252602082015260400190565b60405180910390a45050610827565b836001600160a01b0316856001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb86866040516128c3929190613582565b60405180910390a45050505050565b6001600160a01b0384163b156122c15760405163bc197c8160e01b81526001600160a01b0385169063bc197c819061291690899089908890889088906004016135af565b6020604051808303815f875af1925050508015612950575060408051601f3d908101601f1916820190925261294d9181019061360c565b60015b6129b7573d80801561297d576040519150601f19603f3d011682016040523d82523d5f602084013e612982565b606091505b5080515f036129af57604051632bfa23e760e11b81526001600160a01b03861660048201526024016105bd565b805160208201fd5b6001600160e01b0319811663bc197c8160e01b1461247f57604051632bfa23e760e11b81526001600160a01b03861660048201526024016105bd565b6001600160a01b0384163b156122c15760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190612a379089908990889088908890600401613627565b6020604051808303815f875af1925050508015612a71575060408051601f3d908101601f19168201909252612a6e9181019061360c565b60015b612a9e573d80801561297d576040519150601f19603f3d011682016040523d82523d5f602084013e612982565b6001600160e01b0319811663f23a6e6160e01b1461247f57604051632bfa23e760e11b81526001600160a01b03861660048201526024016105bd565b60405163a9059cbb60e01b5f8181526001600160a01b038616600452602485905291602083604481808b5af1925060015f51148316612b30578383151615612b24573d5f823e3d81fd5b5f873b113d1516831692505b60405250949350505050565b80356001600160a01b0381168114612b52575f5ffd5b919050565b5f5f60408385031215612b68575f5ffd5b612b7183612b3c565b946020939093013593505050565b6001600160e01b0319811681146109c5575f5ffd5b5f60208284031215612ba4575f5ffd5b813561136481612b7f565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715612beb57612beb612baf565b604052919050565b5f82601f830112612c02575f5ffd5b8135602083015f5f6001600160401b03841115612c2157612c21612baf565b50601f8301601f1916602001612c3681612bc3565b915050828152858383011115612c4a575f5ffd5b828260208301375f92810160200192909252509392505050565b5f60208284031215612c74575f5ffd5b81356001600160401b03811115612c89575f5ffd5b612c9584828501612bf3565b949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60ff851681526001600160a01b03841660208201526080604082018190525f90612cf790830185612c9d565b9050821515606083015295945050505050565b5f60208284031215612d1a575f5ffd5b5035919050565b602081525f6113646020830184612c9d565b5f6001600160401b03821115612d4b57612d4b612baf565b5060051b60200190565b5f82601f830112612d64575f5ffd5b8135612d77612d7282612d33565b612bc3565b8082825260208201915060208360051b860101925085831115612d98575f5ffd5b602085015b83811015612db5578035835260209283019201612d9d565b5095945050505050565b5f5f5f5f5f60a08688031215612dd3575f5ffd5b612ddc86612b3c565b9450612dea60208701612b3c565b935060408601356001600160401b03811115612e04575f5ffd5b612e1088828901612d55565b93505060608601356001600160401b03811115612e2b575f5ffd5b612e3788828901612d55565b92505060808601356001600160401b03811115612e52575f5ffd5b612e5e88828901612bf3565b9150509295509295909350565b5f5f60408385031215612e7c575f5ffd5b82359150612e8c60208401612b3c565b90509250929050565b5f60208284031215612ea5575f5ffd5b61136482612b3c565b60a081525f612ec060a0830188612c9d565b6001600160a01b0396909616602083015250604081019390935260608301919091521515608090910152919050565b5f5f5f60608486031215612f01575f5ffd5b83359250612f1160208501612b3c565b929592945050506040919091013590565b5f5f5f60608486031215612f34575f5ffd5b83356001600160401b03811115612f49575f5ffd5b612f5586828701612bf3565b9660208601359650604090950135949350505050565b5f5f60408385031215612f7c575f5ffd5b82356001600160401b03811115612f91575f5ffd5b8301601f81018513612fa1575f5ffd5b8035612faf612d7282612d33565b8082825260208201915060208360051b850101925087831115612fd0575f5ffd5b6020840193505b82841015612ff957612fe884612b3c565b825260209384019390910190612fd7565b945050505060208301356001600160401b03811115613016575f5ffd5b61302285828601612d55565b9150509250929050565b5f8151808452602084019350602083015f5b8281101561305c57815186526020958601959091019060010161303e565b5093949350505050565b602081525f611364602083018461302c565b5f5f5f5f6080858703121561308b575f5ffd5b84359350602085013592506130a260408601612b3c565b915060608501356001600160401b038111156130bc575f5ffd5b6130c887828801612bf3565b91505092959194509250565b5f5f604083850312156130e5575f5ffd5b50508035926020909101359150565b6020815260ff825116602082015260018060a01b0360208301511660408201525f60408301516080606084015261312e60a0840182612c9d565b90506060840151151560808401528091505092915050565b80358015158114612b52575f5ffd5b5f5f60408385031215613166575f5ffd5b61316f83612b3c565b9150612e8c60208401613146565b5f5f5f5f60808587031215613190575f5ffd5b84356001600160401b038111156131a5575f5ffd5b6131b187828801612bf3565b9450506131c060208601612b3c565b925060408501356001600160401b038111156131da575f5ffd5b6131e687828801612bf3565b925050606085013560ff811681146131fc575f5ffd5b939692955090935050565b5f5f5f5f6080858703121561321a575f5ffd5b84356001600160401b0381111561322f575f5ffd5b61323b87828801612d55565b94505060208501356001600160401b03811115613256575f5ffd5b61326287828801612d55565b9350506130a260408601612b3c565b5f5f60408385031215613282575f5ffd5b61328b83612b3c565b9150612e8c60208401612b3c565b5f5f604083850312156132aa575f5ffd5b82359150612e8c60208401613146565b5f5f5f5f5f60a086880312156132ce575f5ffd5b6132d786612b3c565b94506132e560208701612b3c565b9350604086013592506060860135915060808601356001600160401b03811115612e52575f5ffd5b5f82518060208501845e5f920191825250919050565b60208082526010908201526f151bdad95b8e881b9bdd08199bdd5b9960821b604082015260600190565b600181811c9082168061336157607f821691505b60208210810361337f57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f600182016133aa576133aa613385565b5060010190565b601f82111561088b57805f5260205f20601f840160051c810160208510156133d65750805b601f840160051c820191505b81811015610827575f81556001016133e2565b81516001600160401b0381111561340e5761340e612baf565b6134228161341c845461334d565b846133b1565b6020601f821160018114613454575f831561343d5750848201515b5f19600385901b1c1916600184901b178455610827565b5f84815260208120601f198516915b828110156134835787850151825560209485019460019092019101613463565b50848210156134a057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b60208082526014908201527305175616e74697479206d757374206265203e20360641b604082015260600190565b602080825260129082015271566f7563686572206e6f742061637469766560701b604082015260600190565b808202811582820484141761055457610554613385565b8082018082111561055457610554613385565b5f60208284031215613557575f5ffd5b5051919050565b604081525f6135706040830185612c9d565b905060ff831660208301529392505050565b604081525f613594604083018561302c565b82810360208401526135a6818561302c565b95945050505050565b6001600160a01b0386811682528516602082015260a0604082018190525f906135da9083018661302c565b82810360608401526135ec818661302c565b905082810360808401526136008185612c9d565b98975050505050505050565b5f6020828403121561361c575f5ffd5b815161136481612b7f565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190525f9061366090830184612c9d565b97965050505050505056fe9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a69b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212201a52fdcbbb4ca5cd3eabd3880db9448130ce8091613da40f38222f57df4ba7ff64736f6c634300081c0033",
}

// GlintdoughABI is the input ABI used to generate the binding from.
// Deprecated: Use GlintdoughMetaData.ABI instead.
var GlintdoughABI = GlintdoughMetaData.ABI

// GlintdoughBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlintdoughMetaData.Bin instead.
var GlintdoughBin = GlintdoughMetaData.Bin

// DeployGlintdough deploys a new Ethereum contract, binding an instance of Glintdough to it.
func DeployGlintdough(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string, creationfeesAddress common.Address) (common.Address, *types.Transaction, *Glintdough, error) {
	parsed, err := GlintdoughMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlintdoughBin), backend, uri_, creationfeesAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Glintdough{GlintdoughCaller: GlintdoughCaller{contract: contract}, GlintdoughTransactor: GlintdoughTransactor{contract: contract}, GlintdoughFilterer: GlintdoughFilterer{contract: contract}}, nil
}

// Glintdough is an auto generated Go binding around an Ethereum contract.
type Glintdough struct {
	GlintdoughCaller     // Read-only binding to the contract
	GlintdoughTransactor // Write-only binding to the contract
	GlintdoughFilterer   // Log filterer for contract events
}

// GlintdoughCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlintdoughCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlintdoughTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlintdoughTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlintdoughFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlintdoughFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlintdoughSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlintdoughSession struct {
	Contract     *Glintdough       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlintdoughCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlintdoughCallerSession struct {
	Contract *GlintdoughCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GlintdoughTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlintdoughTransactorSession struct {
	Contract     *GlintdoughTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GlintdoughRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlintdoughRaw struct {
	Contract *Glintdough // Generic contract binding to access the raw methods on
}

// GlintdoughCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlintdoughCallerRaw struct {
	Contract *GlintdoughCaller // Generic read-only contract binding to access the raw methods on
}

// GlintdoughTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlintdoughTransactorRaw struct {
	Contract *GlintdoughTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlintdough creates a new instance of Glintdough, bound to a specific deployed contract.
func NewGlintdough(address common.Address, backend bind.ContractBackend) (*Glintdough, error) {
	contract, err := bindGlintdough(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Glintdough{GlintdoughCaller: GlintdoughCaller{contract: contract}, GlintdoughTransactor: GlintdoughTransactor{contract: contract}, GlintdoughFilterer: GlintdoughFilterer{contract: contract}}, nil
}

// NewGlintdoughCaller creates a new read-only instance of Glintdough, bound to a specific deployed contract.
func NewGlintdoughCaller(address common.Address, caller bind.ContractCaller) (*GlintdoughCaller, error) {
	contract, err := bindGlintdough(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlintdoughCaller{contract: contract}, nil
}

// NewGlintdoughTransactor creates a new write-only instance of Glintdough, bound to a specific deployed contract.
func NewGlintdoughTransactor(address common.Address, transactor bind.ContractTransactor) (*GlintdoughTransactor, error) {
	contract, err := bindGlintdough(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlintdoughTransactor{contract: contract}, nil
}

// NewGlintdoughFilterer creates a new log filterer instance of Glintdough, bound to a specific deployed contract.
func NewGlintdoughFilterer(address common.Address, filterer bind.ContractFilterer) (*GlintdoughFilterer, error) {
	contract, err := bindGlintdough(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlintdoughFilterer{contract: contract}, nil
}

// bindGlintdough binds a generic wrapper to an already deployed contract.
func bindGlintdough(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlintdoughMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Glintdough *GlintdoughRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Glintdough.Contract.GlintdoughCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Glintdough *GlintdoughRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glintdough.Contract.GlintdoughTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Glintdough *GlintdoughRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Glintdough.Contract.GlintdoughTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Glintdough *GlintdoughCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Glintdough.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Glintdough *GlintdoughTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glintdough.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Glintdough *GlintdoughTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Glintdough.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Glintdough.Contract.DEFAULTADMINROLE(&_Glintdough.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Glintdough.Contract.DEFAULTADMINROLE(&_Glintdough.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughSession) MINTERROLE() ([32]byte, error) {
	return _Glintdough.Contract.MINTERROLE(&_Glintdough.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Glintdough *GlintdoughCallerSession) MINTERROLE() ([32]byte, error) {
	return _Glintdough.Contract.MINTERROLE(&_Glintdough.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Glintdough *GlintdoughCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Glintdough *GlintdoughSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Glintdough.Contract.BalanceOf(&_Glintdough.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Glintdough *GlintdoughCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Glintdough.Contract.BalanceOf(&_Glintdough.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Glintdough *GlintdoughCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Glintdough *GlintdoughSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Glintdough.Contract.BalanceOfBatch(&_Glintdough.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Glintdough *GlintdoughCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Glintdough.Contract.BalanceOfBatch(&_Glintdough.CallOpts, accounts, ids)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x0d398b5a.
//
// Solidity: function collateralTokens(string ) view returns(uint8 decimals, address smartContract, string symbol, bool isActive)
func (_Glintdough *GlintdoughCaller) CollateralTokens(opts *bind.CallOpts, arg0 string) (struct {
	Decimals      uint8
	SmartContract common.Address
	Symbol        string
	IsActive      bool
}, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "collateralTokens", arg0)

	outstruct := new(struct {
		Decimals      uint8
		SmartContract common.Address
		Symbol        string
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimals = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SmartContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// CollateralTokens is a free data retrieval call binding the contract method 0x0d398b5a.
//
// Solidity: function collateralTokens(string ) view returns(uint8 decimals, address smartContract, string symbol, bool isActive)
func (_Glintdough *GlintdoughSession) CollateralTokens(arg0 string) (struct {
	Decimals      uint8
	SmartContract common.Address
	Symbol        string
	IsActive      bool
}, error) {
	return _Glintdough.Contract.CollateralTokens(&_Glintdough.CallOpts, arg0)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x0d398b5a.
//
// Solidity: function collateralTokens(string ) view returns(uint8 decimals, address smartContract, string symbol, bool isActive)
func (_Glintdough *GlintdoughCallerSession) CollateralTokens(arg0 string) (struct {
	Decimals      uint8
	SmartContract common.Address
	Symbol        string
	IsActive      bool
}, error) {
	return _Glintdough.Contract.CollateralTokens(&_Glintdough.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string key) view returns(bool)
func (_Glintdough *GlintdoughCaller) Exists(opts *bind.CallOpts, key string) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "exists", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string key) view returns(bool)
func (_Glintdough *GlintdoughSession) Exists(key string) (bool, error) {
	return _Glintdough.Contract.Exists(&_Glintdough.CallOpts, key)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string key) view returns(bool)
func (_Glintdough *GlintdoughCallerSession) Exists(key string) (bool, error) {
	return _Glintdough.Contract.Exists(&_Glintdough.CallOpts, key)
}

// ExistsByAddress is a free data retrieval call binding the contract method 0x94a94f31.
//
// Solidity: function existsByAddress(address smartContract) view returns(bool)
func (_Glintdough *GlintdoughCaller) ExistsByAddress(opts *bind.CallOpts, smartContract common.Address) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "existsByAddress", smartContract)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExistsByAddress is a free data retrieval call binding the contract method 0x94a94f31.
//
// Solidity: function existsByAddress(address smartContract) view returns(bool)
func (_Glintdough *GlintdoughSession) ExistsByAddress(smartContract common.Address) (bool, error) {
	return _Glintdough.Contract.ExistsByAddress(&_Glintdough.CallOpts, smartContract)
}

// ExistsByAddress is a free data retrieval call binding the contract method 0x94a94f31.
//
// Solidity: function existsByAddress(address smartContract) view returns(bool)
func (_Glintdough *GlintdoughCallerSession) ExistsByAddress(smartContract common.Address) (bool, error) {
	return _Glintdough.Contract.ExistsByAddress(&_Glintdough.CallOpts, smartContract)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Glintdough *GlintdoughCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Glintdough *GlintdoughSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Glintdough.Contract.GetRoleAdmin(&_Glintdough.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Glintdough *GlintdoughCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Glintdough.Contract.GetRoleAdmin(&_Glintdough.CallOpts, role)
}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(string key) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughCaller) GetToken(opts *bind.CallOpts, key string) (GlintDoughTokenData, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "getToken", key)

	if err != nil {
		return *new(GlintDoughTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(GlintDoughTokenData)).(*GlintDoughTokenData)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(string key) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughSession) GetToken(key string) (GlintDoughTokenData, error) {
	return _Glintdough.Contract.GetToken(&_Glintdough.CallOpts, key)
}

// GetToken is a free data retrieval call binding the contract method 0xc1733f68.
//
// Solidity: function getToken(string key) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughCallerSession) GetToken(key string) (GlintDoughTokenData, error) {
	return _Glintdough.Contract.GetToken(&_Glintdough.CallOpts, key)
}

// GetTokenByAddress is a free data retrieval call binding the contract method 0x91ded8fa.
//
// Solidity: function getTokenByAddress(address smartContract) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughCaller) GetTokenByAddress(opts *bind.CallOpts, smartContract common.Address) (GlintDoughTokenData, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "getTokenByAddress", smartContract)

	if err != nil {
		return *new(GlintDoughTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(GlintDoughTokenData)).(*GlintDoughTokenData)

	return out0, err

}

// GetTokenByAddress is a free data retrieval call binding the contract method 0x91ded8fa.
//
// Solidity: function getTokenByAddress(address smartContract) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughSession) GetTokenByAddress(smartContract common.Address) (GlintDoughTokenData, error) {
	return _Glintdough.Contract.GetTokenByAddress(&_Glintdough.CallOpts, smartContract)
}

// GetTokenByAddress is a free data retrieval call binding the contract method 0x91ded8fa.
//
// Solidity: function getTokenByAddress(address smartContract) view returns((uint8,address,string,bool))
func (_Glintdough *GlintdoughCallerSession) GetTokenByAddress(smartContract common.Address) (GlintDoughTokenData, error) {
	return _Glintdough.Contract.GetTokenByAddress(&_Glintdough.CallOpts, smartContract)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Glintdough *GlintdoughCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Glintdough *GlintdoughSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Glintdough.Contract.HasRole(&_Glintdough.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Glintdough *GlintdoughCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Glintdough.Contract.HasRole(&_Glintdough.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Glintdough *GlintdoughCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Glintdough *GlintdoughSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Glintdough.Contract.IsApprovedForAll(&_Glintdough.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Glintdough *GlintdoughCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Glintdough.Contract.IsApprovedForAll(&_Glintdough.CallOpts, account, operator)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Glintdough *GlintdoughCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Glintdough *GlintdoughSession) NextTokenId() (*big.Int, error) {
	return _Glintdough.Contract.NextTokenId(&_Glintdough.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Glintdough *GlintdoughCallerSession) NextTokenId() (*big.Int, error) {
	return _Glintdough.Contract.NextTokenId(&_Glintdough.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Glintdough *GlintdoughCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Glintdough *GlintdoughSession) Paused() (bool, error) {
	return _Glintdough.Contract.Paused(&_Glintdough.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Glintdough *GlintdoughCallerSession) Paused() (bool, error) {
	return _Glintdough.Contract.Paused(&_Glintdough.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Glintdough *GlintdoughCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Glintdough *GlintdoughSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Glintdough.Contract.SupportsInterface(&_Glintdough.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Glintdough *GlintdoughCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Glintdough.Contract.SupportsInterface(&_Glintdough.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Glintdough *GlintdoughCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Glintdough *GlintdoughSession) Uri(arg0 *big.Int) (string, error) {
	return _Glintdough.Contract.Uri(&_Glintdough.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_Glintdough *GlintdoughCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _Glintdough.Contract.Uri(&_Glintdough.CallOpts, arg0)
}

// Vouchers is a free data retrieval call binding the contract method 0x3e97b43f.
//
// Solidity: function vouchers(uint256 ) view returns(string token, address collateral_token, uint256 denomination, uint256 creationFees, bool is_active)
func (_Glintdough *GlintdoughCaller) Vouchers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token           string
	CollateralToken common.Address
	Denomination    *big.Int
	CreationFees    *big.Int
	IsActive        bool
}, error) {
	var out []interface{}
	err := _Glintdough.contract.Call(opts, &out, "vouchers", arg0)

	outstruct := new(struct {
		Token           string
		CollateralToken common.Address
		Denomination    *big.Int
		CreationFees    *big.Int
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CollateralToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Denomination = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreationFees = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Vouchers is a free data retrieval call binding the contract method 0x3e97b43f.
//
// Solidity: function vouchers(uint256 ) view returns(string token, address collateral_token, uint256 denomination, uint256 creationFees, bool is_active)
func (_Glintdough *GlintdoughSession) Vouchers(arg0 *big.Int) (struct {
	Token           string
	CollateralToken common.Address
	Denomination    *big.Int
	CreationFees    *big.Int
	IsActive        bool
}, error) {
	return _Glintdough.Contract.Vouchers(&_Glintdough.CallOpts, arg0)
}

// Vouchers is a free data retrieval call binding the contract method 0x3e97b43f.
//
// Solidity: function vouchers(uint256 ) view returns(string token, address collateral_token, uint256 denomination, uint256 creationFees, bool is_active)
func (_Glintdough *GlintdoughCallerSession) Vouchers(arg0 *big.Int) (struct {
	Token           string
	CollateralToken common.Address
	Denomination    *big.Int
	CreationFees    *big.Int
	IsActive        bool
}, error) {
	return _Glintdough.Contract.Vouchers(&_Glintdough.CallOpts, arg0)
}

// ActivateCollateralToken is a paid mutator transaction binding the contract method 0x08983f6d.
//
// Solidity: function activateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughTransactor) ActivateCollateralToken(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "activateCollateralToken", key)
}

// ActivateCollateralToken is a paid mutator transaction binding the contract method 0x08983f6d.
//
// Solidity: function activateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughSession) ActivateCollateralToken(key string) (*types.Transaction, error) {
	return _Glintdough.Contract.ActivateCollateralToken(&_Glintdough.TransactOpts, key)
}

// ActivateCollateralToken is a paid mutator transaction binding the contract method 0x08983f6d.
//
// Solidity: function activateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughTransactorSession) ActivateCollateralToken(key string) (*types.Transaction, error) {
	return _Glintdough.Contract.ActivateCollateralToken(&_Glintdough.TransactOpts, key)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0xbdbacb95.
//
// Solidity: function addCollateralToken(string key, address smartContract, string symbol, uint8 decimals) returns()
func (_Glintdough *GlintdoughTransactor) AddCollateralToken(opts *bind.TransactOpts, key string, smartContract common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "addCollateralToken", key, smartContract, symbol, decimals)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0xbdbacb95.
//
// Solidity: function addCollateralToken(string key, address smartContract, string symbol, uint8 decimals) returns()
func (_Glintdough *GlintdoughSession) AddCollateralToken(key string, smartContract common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Glintdough.Contract.AddCollateralToken(&_Glintdough.TransactOpts, key, smartContract, symbol, decimals)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0xbdbacb95.
//
// Solidity: function addCollateralToken(string key, address smartContract, string symbol, uint8 decimals) returns()
func (_Glintdough *GlintdoughTransactorSession) AddCollateralToken(key string, smartContract common.Address, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Glintdough.Contract.AddCollateralToken(&_Glintdough.TransactOpts, key, smartContract, symbol, decimals)
}

// AddVoucher is a paid mutator transaction binding the contract method 0x4dd91ff9.
//
// Solidity: function addVoucher(string token, uint256 denomination, uint256 creationFees) returns(uint256)
func (_Glintdough *GlintdoughTransactor) AddVoucher(opts *bind.TransactOpts, token string, denomination *big.Int, creationFees *big.Int) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "addVoucher", token, denomination, creationFees)
}

// AddVoucher is a paid mutator transaction binding the contract method 0x4dd91ff9.
//
// Solidity: function addVoucher(string token, uint256 denomination, uint256 creationFees) returns(uint256)
func (_Glintdough *GlintdoughSession) AddVoucher(token string, denomination *big.Int, creationFees *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.AddVoucher(&_Glintdough.TransactOpts, token, denomination, creationFees)
}

// AddVoucher is a paid mutator transaction binding the contract method 0x4dd91ff9.
//
// Solidity: function addVoucher(string token, uint256 denomination, uint256 creationFees) returns(uint256)
func (_Glintdough *GlintdoughTransactorSession) AddVoucher(token string, denomination *big.Int, creationFees *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.AddVoucher(&_Glintdough.TransactOpts, token, denomination, creationFees)
}

// DeactivateCollateralToken is a paid mutator transaction binding the contract method 0xe0b31602.
//
// Solidity: function deactivateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughTransactor) DeactivateCollateralToken(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "deactivateCollateralToken", key)
}

// DeactivateCollateralToken is a paid mutator transaction binding the contract method 0xe0b31602.
//
// Solidity: function deactivateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughSession) DeactivateCollateralToken(key string) (*types.Transaction, error) {
	return _Glintdough.Contract.DeactivateCollateralToken(&_Glintdough.TransactOpts, key)
}

// DeactivateCollateralToken is a paid mutator transaction binding the contract method 0xe0b31602.
//
// Solidity: function deactivateCollateralToken(string key) returns()
func (_Glintdough *GlintdoughTransactorSession) DeactivateCollateralToken(key string) (*types.Transaction, error) {
	return _Glintdough.Contract.DeactivateCollateralToken(&_Glintdough.TransactOpts, key)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_Glintdough *GlintdoughTransactor) GrantMinterRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "grantMinterRole", minter)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_Glintdough *GlintdoughSession) GrantMinterRole(minter common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.GrantMinterRole(&_Glintdough.TransactOpts, minter)
}

// GrantMinterRole is a paid mutator transaction binding the contract method 0x3dd1eb61.
//
// Solidity: function grantMinterRole(address minter) returns()
func (_Glintdough *GlintdoughTransactorSession) GrantMinterRole(minter common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.GrantMinterRole(&_Glintdough.TransactOpts, minter)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.GrantRole(&_Glintdough.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.GrantRole(&_Glintdough.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x5ca5933a.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, address to, bytes data) returns()
func (_Glintdough *GlintdoughTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, quantity *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "mint", tokenId, quantity, to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x5ca5933a.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, address to, bytes data) returns()
func (_Glintdough *GlintdoughSession) Mint(tokenId *big.Int, quantity *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.Mint(&_Glintdough.TransactOpts, tokenId, quantity, to, data)
}

// Mint is a paid mutator transaction binding the contract method 0x5ca5933a.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, address to, bytes data) returns()
func (_Glintdough *GlintdoughTransactorSession) Mint(tokenId *big.Int, quantity *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.Mint(&_Glintdough.TransactOpts, tokenId, quantity, to, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0xcf9e099c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, address to, bytes data) returns()
func (_Glintdough *GlintdoughTransactor) MintBatch(opts *bind.TransactOpts, tokenIds []*big.Int, quantities []*big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "mintBatch", tokenIds, quantities, to, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0xcf9e099c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, address to, bytes data) returns()
func (_Glintdough *GlintdoughSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.MintBatch(&_Glintdough.TransactOpts, tokenIds, quantities, to, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0xcf9e099c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, address to, bytes data) returns()
func (_Glintdough *GlintdoughTransactorSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.MintBatch(&_Glintdough.TransactOpts, tokenIds, quantities, to, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Glintdough *GlintdoughTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Glintdough *GlintdoughSession) Pause() (*types.Transaction, error) {
	return _Glintdough.Contract.Pause(&_Glintdough.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Glintdough *GlintdoughTransactorSession) Pause() (*types.Transaction, error) {
	return _Glintdough.Contract.Pause(&_Glintdough.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 tokenId, uint256 quantity) returns()
func (_Glintdough *GlintdoughTransactor) Redeem(opts *bind.TransactOpts, tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "redeem", tokenId, quantity)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 tokenId, uint256 quantity) returns()
func (_Glintdough *GlintdoughSession) Redeem(tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.Redeem(&_Glintdough.TransactOpts, tokenId, quantity)
}

// Redeem is a paid mutator transaction binding the contract method 0x7cbc2373.
//
// Solidity: function redeem(uint256 tokenId, uint256 quantity) returns()
func (_Glintdough *GlintdoughTransactorSession) Redeem(tokenId *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.Redeem(&_Glintdough.TransactOpts, tokenId, quantity)
}

// RedeemTo is a paid mutator transaction binding the contract method 0x46bcc953.
//
// Solidity: function redeemTo(uint256 tokenId, address to, uint256 quantity) returns()
func (_Glintdough *GlintdoughTransactor) RedeemTo(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "redeemTo", tokenId, to, quantity)
}

// RedeemTo is a paid mutator transaction binding the contract method 0x46bcc953.
//
// Solidity: function redeemTo(uint256 tokenId, address to, uint256 quantity) returns()
func (_Glintdough *GlintdoughSession) RedeemTo(tokenId *big.Int, to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.RedeemTo(&_Glintdough.TransactOpts, tokenId, to, quantity)
}

// RedeemTo is a paid mutator transaction binding the contract method 0x46bcc953.
//
// Solidity: function redeemTo(uint256 tokenId, address to, uint256 quantity) returns()
func (_Glintdough *GlintdoughTransactorSession) RedeemTo(tokenId *big.Int, to common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Glintdough.Contract.RedeemTo(&_Glintdough.TransactOpts, tokenId, to, quantity)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Glintdough *GlintdoughTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Glintdough *GlintdoughSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.RenounceRole(&_Glintdough.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Glintdough *GlintdoughTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.RenounceRole(&_Glintdough.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.RevokeRole(&_Glintdough.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Glintdough *GlintdoughTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.RevokeRole(&_Glintdough.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Glintdough *GlintdoughTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Glintdough *GlintdoughSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.SafeBatchTransferFrom(&_Glintdough.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_Glintdough *GlintdoughTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.SafeBatchTransferFrom(&_Glintdough.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Glintdough *GlintdoughTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Glintdough *GlintdoughSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.SafeTransferFrom(&_Glintdough.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_Glintdough *GlintdoughTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Glintdough.Contract.SafeTransferFrom(&_Glintdough.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Glintdough *GlintdoughTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Glintdough *GlintdoughSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Glintdough.Contract.SetApprovalForAll(&_Glintdough.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Glintdough *GlintdoughTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Glintdough.Contract.SetApprovalForAll(&_Glintdough.TransactOpts, operator, approved)
}

// SetFeesAddress is a paid mutator transaction binding the contract method 0x736e4f57.
//
// Solidity: function setFeesAddress(address newFeesAddress) returns()
func (_Glintdough *GlintdoughTransactor) SetFeesAddress(opts *bind.TransactOpts, newFeesAddress common.Address) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "setFeesAddress", newFeesAddress)
}

// SetFeesAddress is a paid mutator transaction binding the contract method 0x736e4f57.
//
// Solidity: function setFeesAddress(address newFeesAddress) returns()
func (_Glintdough *GlintdoughSession) SetFeesAddress(newFeesAddress common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.SetFeesAddress(&_Glintdough.TransactOpts, newFeesAddress)
}

// SetFeesAddress is a paid mutator transaction binding the contract method 0x736e4f57.
//
// Solidity: function setFeesAddress(address newFeesAddress) returns()
func (_Glintdough *GlintdoughTransactorSession) SetFeesAddress(newFeesAddress common.Address) (*types.Transaction, error) {
	return _Glintdough.Contract.SetFeesAddress(&_Glintdough.TransactOpts, newFeesAddress)
}

// ToggleVoucher is a paid mutator transaction binding the contract method 0xf08f9796.
//
// Solidity: function toggleVoucher(uint256 tokenId, bool status) returns()
func (_Glintdough *GlintdoughTransactor) ToggleVoucher(opts *bind.TransactOpts, tokenId *big.Int, status bool) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "toggleVoucher", tokenId, status)
}

// ToggleVoucher is a paid mutator transaction binding the contract method 0xf08f9796.
//
// Solidity: function toggleVoucher(uint256 tokenId, bool status) returns()
func (_Glintdough *GlintdoughSession) ToggleVoucher(tokenId *big.Int, status bool) (*types.Transaction, error) {
	return _Glintdough.Contract.ToggleVoucher(&_Glintdough.TransactOpts, tokenId, status)
}

// ToggleVoucher is a paid mutator transaction binding the contract method 0xf08f9796.
//
// Solidity: function toggleVoucher(uint256 tokenId, bool status) returns()
func (_Glintdough *GlintdoughTransactorSession) ToggleVoucher(tokenId *big.Int, status bool) (*types.Transaction, error) {
	return _Glintdough.Contract.ToggleVoucher(&_Glintdough.TransactOpts, tokenId, status)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Glintdough *GlintdoughTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Glintdough.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Glintdough *GlintdoughSession) Unpause() (*types.Transaction, error) {
	return _Glintdough.Contract.Unpause(&_Glintdough.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Glintdough *GlintdoughTransactorSession) Unpause() (*types.Transaction, error) {
	return _Glintdough.Contract.Unpause(&_Glintdough.TransactOpts)
}

// GlintdoughApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Glintdough contract.
type GlintdoughApprovalForAllIterator struct {
	Event *GlintdoughApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GlintdoughApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughApprovalForAll)
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
		it.Event = new(GlintdoughApprovalForAll)
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
func (it *GlintdoughApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughApprovalForAll represents a ApprovalForAll event raised by the Glintdough contract.
type GlintdoughApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Glintdough *GlintdoughFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*GlintdoughApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughApprovalForAllIterator{contract: _Glintdough.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Glintdough *GlintdoughFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GlintdoughApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughApprovalForAll)
				if err := _Glintdough.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Glintdough *GlintdoughFilterer) ParseApprovalForAll(log types.Log) (*GlintdoughApprovalForAll, error) {
	event := new(GlintdoughApprovalForAll)
	if err := _Glintdough.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughFeesAddressUpdatedIterator is returned from FilterFeesAddressUpdated and is used to iterate over the raw logs and unpacked data for FeesAddressUpdated events raised by the Glintdough contract.
type GlintdoughFeesAddressUpdatedIterator struct {
	Event *GlintdoughFeesAddressUpdated // Event containing the contract specifics and raw log

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
func (it *GlintdoughFeesAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughFeesAddressUpdated)
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
		it.Event = new(GlintdoughFeesAddressUpdated)
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
func (it *GlintdoughFeesAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughFeesAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughFeesAddressUpdated represents a FeesAddressUpdated event raised by the Glintdough contract.
type GlintdoughFeesAddressUpdated struct {
	FeesAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesAddressUpdated is a free log retrieval operation binding the contract event 0x3c7fda2a649f437dd79167a620fa573f40c14c0c0204f16b578cde06433bbb14.
//
// Solidity: event FeesAddressUpdated(address indexed feesAddress)
func (_Glintdough *GlintdoughFilterer) FilterFeesAddressUpdated(opts *bind.FilterOpts, feesAddress []common.Address) (*GlintdoughFeesAddressUpdatedIterator, error) {

	var feesAddressRule []interface{}
	for _, feesAddressItem := range feesAddress {
		feesAddressRule = append(feesAddressRule, feesAddressItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "FeesAddressUpdated", feesAddressRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughFeesAddressUpdatedIterator{contract: _Glintdough.contract, event: "FeesAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesAddressUpdated is a free log subscription operation binding the contract event 0x3c7fda2a649f437dd79167a620fa573f40c14c0c0204f16b578cde06433bbb14.
//
// Solidity: event FeesAddressUpdated(address indexed feesAddress)
func (_Glintdough *GlintdoughFilterer) WatchFeesAddressUpdated(opts *bind.WatchOpts, sink chan<- *GlintdoughFeesAddressUpdated, feesAddress []common.Address) (event.Subscription, error) {

	var feesAddressRule []interface{}
	for _, feesAddressItem := range feesAddress {
		feesAddressRule = append(feesAddressRule, feesAddressItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "FeesAddressUpdated", feesAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughFeesAddressUpdated)
				if err := _Glintdough.contract.UnpackLog(event, "FeesAddressUpdated", log); err != nil {
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

// ParseFeesAddressUpdated is a log parse operation binding the contract event 0x3c7fda2a649f437dd79167a620fa573f40c14c0c0204f16b578cde06433bbb14.
//
// Solidity: event FeesAddressUpdated(address indexed feesAddress)
func (_Glintdough *GlintdoughFilterer) ParseFeesAddressUpdated(log types.Log) (*GlintdoughFeesAddressUpdated, error) {
	event := new(GlintdoughFeesAddressUpdated)
	if err := _Glintdough.contract.UnpackLog(event, "FeesAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughMinterRoleGrantedIterator is returned from FilterMinterRoleGranted and is used to iterate over the raw logs and unpacked data for MinterRoleGranted events raised by the Glintdough contract.
type GlintdoughMinterRoleGrantedIterator struct {
	Event *GlintdoughMinterRoleGranted // Event containing the contract specifics and raw log

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
func (it *GlintdoughMinterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughMinterRoleGranted)
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
		it.Event = new(GlintdoughMinterRoleGranted)
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
func (it *GlintdoughMinterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughMinterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughMinterRoleGranted represents a MinterRoleGranted event raised by the Glintdough contract.
type GlintdoughMinterRoleGranted struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRoleGranted is a free log retrieval operation binding the contract event 0xbd42b6cd963a984d8cbff85acad4dd2c87a6d85ac756d16e6d2ec5ce6c8f0730.
//
// Solidity: event MinterRoleGranted(address minter)
func (_Glintdough *GlintdoughFilterer) FilterMinterRoleGranted(opts *bind.FilterOpts) (*GlintdoughMinterRoleGrantedIterator, error) {

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "MinterRoleGranted")
	if err != nil {
		return nil, err
	}
	return &GlintdoughMinterRoleGrantedIterator{contract: _Glintdough.contract, event: "MinterRoleGranted", logs: logs, sub: sub}, nil
}

// WatchMinterRoleGranted is a free log subscription operation binding the contract event 0xbd42b6cd963a984d8cbff85acad4dd2c87a6d85ac756d16e6d2ec5ce6c8f0730.
//
// Solidity: event MinterRoleGranted(address minter)
func (_Glintdough *GlintdoughFilterer) WatchMinterRoleGranted(opts *bind.WatchOpts, sink chan<- *GlintdoughMinterRoleGranted) (event.Subscription, error) {

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "MinterRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughMinterRoleGranted)
				if err := _Glintdough.contract.UnpackLog(event, "MinterRoleGranted", log); err != nil {
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

// ParseMinterRoleGranted is a log parse operation binding the contract event 0xbd42b6cd963a984d8cbff85acad4dd2c87a6d85ac756d16e6d2ec5ce6c8f0730.
//
// Solidity: event MinterRoleGranted(address minter)
func (_Glintdough *GlintdoughFilterer) ParseMinterRoleGranted(log types.Log) (*GlintdoughMinterRoleGranted, error) {
	event := new(GlintdoughMinterRoleGranted)
	if err := _Glintdough.contract.UnpackLog(event, "MinterRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Glintdough contract.
type GlintdoughPausedIterator struct {
	Event *GlintdoughPaused // Event containing the contract specifics and raw log

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
func (it *GlintdoughPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughPaused)
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
		it.Event = new(GlintdoughPaused)
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
func (it *GlintdoughPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughPaused represents a Paused event raised by the Glintdough contract.
type GlintdoughPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Glintdough *GlintdoughFilterer) FilterPaused(opts *bind.FilterOpts) (*GlintdoughPausedIterator, error) {

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GlintdoughPausedIterator{contract: _Glintdough.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Glintdough *GlintdoughFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GlintdoughPaused) (event.Subscription, error) {

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughPaused)
				if err := _Glintdough.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Glintdough *GlintdoughFilterer) ParsePaused(log types.Log) (*GlintdoughPaused, error) {
	event := new(GlintdoughPaused)
	if err := _Glintdough.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughPhysicalCardRegisteredIterator is returned from FilterPhysicalCardRegistered and is used to iterate over the raw logs and unpacked data for PhysicalCardRegistered events raised by the Glintdough contract.
type GlintdoughPhysicalCardRegisteredIterator struct {
	Event *GlintdoughPhysicalCardRegistered // Event containing the contract specifics and raw log

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
func (it *GlintdoughPhysicalCardRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughPhysicalCardRegistered)
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
		it.Event = new(GlintdoughPhysicalCardRegistered)
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
func (it *GlintdoughPhysicalCardRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughPhysicalCardRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughPhysicalCardRegistered represents a PhysicalCardRegistered event raised by the Glintdough contract.
type GlintdoughPhysicalCardRegistered struct {
	SerialHash [32]byte
	Checksum   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPhysicalCardRegistered is a free log retrieval operation binding the contract event 0x0611d64e75dc545b8c42fa01df57f011f9e4127e54faf9bf61d9992a63fb091b.
//
// Solidity: event PhysicalCardRegistered(bytes32 indexed serialHash, bytes32 checksum)
func (_Glintdough *GlintdoughFilterer) FilterPhysicalCardRegistered(opts *bind.FilterOpts, serialHash [][32]byte) (*GlintdoughPhysicalCardRegisteredIterator, error) {

	var serialHashRule []interface{}
	for _, serialHashItem := range serialHash {
		serialHashRule = append(serialHashRule, serialHashItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "PhysicalCardRegistered", serialHashRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughPhysicalCardRegisteredIterator{contract: _Glintdough.contract, event: "PhysicalCardRegistered", logs: logs, sub: sub}, nil
}

// WatchPhysicalCardRegistered is a free log subscription operation binding the contract event 0x0611d64e75dc545b8c42fa01df57f011f9e4127e54faf9bf61d9992a63fb091b.
//
// Solidity: event PhysicalCardRegistered(bytes32 indexed serialHash, bytes32 checksum)
func (_Glintdough *GlintdoughFilterer) WatchPhysicalCardRegistered(opts *bind.WatchOpts, sink chan<- *GlintdoughPhysicalCardRegistered, serialHash [][32]byte) (event.Subscription, error) {

	var serialHashRule []interface{}
	for _, serialHashItem := range serialHash {
		serialHashRule = append(serialHashRule, serialHashItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "PhysicalCardRegistered", serialHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughPhysicalCardRegistered)
				if err := _Glintdough.contract.UnpackLog(event, "PhysicalCardRegistered", log); err != nil {
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

// ParsePhysicalCardRegistered is a log parse operation binding the contract event 0x0611d64e75dc545b8c42fa01df57f011f9e4127e54faf9bf61d9992a63fb091b.
//
// Solidity: event PhysicalCardRegistered(bytes32 indexed serialHash, bytes32 checksum)
func (_Glintdough *GlintdoughFilterer) ParsePhysicalCardRegistered(log types.Log) (*GlintdoughPhysicalCardRegistered, error) {
	event := new(GlintdoughPhysicalCardRegistered)
	if err := _Glintdough.contract.UnpackLog(event, "PhysicalCardRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Glintdough contract.
type GlintdoughRoleAdminChangedIterator struct {
	Event *GlintdoughRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GlintdoughRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughRoleAdminChanged)
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
		it.Event = new(GlintdoughRoleAdminChanged)
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
func (it *GlintdoughRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughRoleAdminChanged represents a RoleAdminChanged event raised by the Glintdough contract.
type GlintdoughRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Glintdough *GlintdoughFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GlintdoughRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughRoleAdminChangedIterator{contract: _Glintdough.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Glintdough *GlintdoughFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GlintdoughRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughRoleAdminChanged)
				if err := _Glintdough.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Glintdough *GlintdoughFilterer) ParseRoleAdminChanged(log types.Log) (*GlintdoughRoleAdminChanged, error) {
	event := new(GlintdoughRoleAdminChanged)
	if err := _Glintdough.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Glintdough contract.
type GlintdoughRoleGrantedIterator struct {
	Event *GlintdoughRoleGranted // Event containing the contract specifics and raw log

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
func (it *GlintdoughRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughRoleGranted)
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
		it.Event = new(GlintdoughRoleGranted)
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
func (it *GlintdoughRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughRoleGranted represents a RoleGranted event raised by the Glintdough contract.
type GlintdoughRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Glintdough *GlintdoughFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GlintdoughRoleGrantedIterator, error) {

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

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughRoleGrantedIterator{contract: _Glintdough.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Glintdough *GlintdoughFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GlintdoughRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughRoleGranted)
				if err := _Glintdough.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Glintdough *GlintdoughFilterer) ParseRoleGranted(log types.Log) (*GlintdoughRoleGranted, error) {
	event := new(GlintdoughRoleGranted)
	if err := _Glintdough.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Glintdough contract.
type GlintdoughRoleRevokedIterator struct {
	Event *GlintdoughRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GlintdoughRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughRoleRevoked)
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
		it.Event = new(GlintdoughRoleRevoked)
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
func (it *GlintdoughRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughRoleRevoked represents a RoleRevoked event raised by the Glintdough contract.
type GlintdoughRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Glintdough *GlintdoughFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GlintdoughRoleRevokedIterator, error) {

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

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughRoleRevokedIterator{contract: _Glintdough.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Glintdough *GlintdoughFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GlintdoughRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughRoleRevoked)
				if err := _Glintdough.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Glintdough *GlintdoughFilterer) ParseRoleRevoked(log types.Log) (*GlintdoughRoleRevoked, error) {
	event := new(GlintdoughRoleRevoked)
	if err := _Glintdough.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Glintdough contract.
type GlintdoughTokenAddedIterator struct {
	Event *GlintdoughTokenAdded // Event containing the contract specifics and raw log

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
func (it *GlintdoughTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughTokenAdded)
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
		it.Event = new(GlintdoughTokenAdded)
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
func (it *GlintdoughTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughTokenAdded represents a TokenAdded event raised by the Glintdough contract.
type GlintdoughTokenAdded struct {
	Key           common.Hash
	SmartContract common.Address
	Symbol        string
	Decimals      uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x03fcf4176a3f2bdb8a0e11fc763d1b887586774f6f6d3e87506e3aa3c20e8e31.
//
// Solidity: event TokenAdded(string indexed key, address indexed smartContract, string symbol, uint8 decimals)
func (_Glintdough *GlintdoughFilterer) FilterTokenAdded(opts *bind.FilterOpts, key []string, smartContract []common.Address) (*GlintdoughTokenAddedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var smartContractRule []interface{}
	for _, smartContractItem := range smartContract {
		smartContractRule = append(smartContractRule, smartContractItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "TokenAdded", keyRule, smartContractRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughTokenAddedIterator{contract: _Glintdough.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x03fcf4176a3f2bdb8a0e11fc763d1b887586774f6f6d3e87506e3aa3c20e8e31.
//
// Solidity: event TokenAdded(string indexed key, address indexed smartContract, string symbol, uint8 decimals)
func (_Glintdough *GlintdoughFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *GlintdoughTokenAdded, key []string, smartContract []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var smartContractRule []interface{}
	for _, smartContractItem := range smartContract {
		smartContractRule = append(smartContractRule, smartContractItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "TokenAdded", keyRule, smartContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughTokenAdded)
				if err := _Glintdough.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x03fcf4176a3f2bdb8a0e11fc763d1b887586774f6f6d3e87506e3aa3c20e8e31.
//
// Solidity: event TokenAdded(string indexed key, address indexed smartContract, string symbol, uint8 decimals)
func (_Glintdough *GlintdoughFilterer) ParseTokenAdded(log types.Log) (*GlintdoughTokenAdded, error) {
	event := new(GlintdoughTokenAdded)
	if err := _Glintdough.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughTokenStatusChangedIterator is returned from FilterTokenStatusChanged and is used to iterate over the raw logs and unpacked data for TokenStatusChanged events raised by the Glintdough contract.
type GlintdoughTokenStatusChangedIterator struct {
	Event *GlintdoughTokenStatusChanged // Event containing the contract specifics and raw log

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
func (it *GlintdoughTokenStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughTokenStatusChanged)
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
		it.Event = new(GlintdoughTokenStatusChanged)
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
func (it *GlintdoughTokenStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughTokenStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughTokenStatusChanged represents a TokenStatusChanged event raised by the Glintdough contract.
type GlintdoughTokenStatusChanged struct {
	Key      common.Hash
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenStatusChanged is a free log retrieval operation binding the contract event 0x41cda4530bfc26ff01b587bd9fc0bcc8c67aa176a168da1542f9dd2f333be8a2.
//
// Solidity: event TokenStatusChanged(string indexed key, bool isActive)
func (_Glintdough *GlintdoughFilterer) FilterTokenStatusChanged(opts *bind.FilterOpts, key []string) (*GlintdoughTokenStatusChangedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "TokenStatusChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughTokenStatusChangedIterator{contract: _Glintdough.contract, event: "TokenStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTokenStatusChanged is a free log subscription operation binding the contract event 0x41cda4530bfc26ff01b587bd9fc0bcc8c67aa176a168da1542f9dd2f333be8a2.
//
// Solidity: event TokenStatusChanged(string indexed key, bool isActive)
func (_Glintdough *GlintdoughFilterer) WatchTokenStatusChanged(opts *bind.WatchOpts, sink chan<- *GlintdoughTokenStatusChanged, key []string) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "TokenStatusChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughTokenStatusChanged)
				if err := _Glintdough.contract.UnpackLog(event, "TokenStatusChanged", log); err != nil {
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

// ParseTokenStatusChanged is a log parse operation binding the contract event 0x41cda4530bfc26ff01b587bd9fc0bcc8c67aa176a168da1542f9dd2f333be8a2.
//
// Solidity: event TokenStatusChanged(string indexed key, bool isActive)
func (_Glintdough *GlintdoughFilterer) ParseTokenStatusChanged(log types.Log) (*GlintdoughTokenStatusChanged, error) {
	event := new(GlintdoughTokenStatusChanged)
	if err := _Glintdough.contract.UnpackLog(event, "TokenStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Glintdough contract.
type GlintdoughTransferBatchIterator struct {
	Event *GlintdoughTransferBatch // Event containing the contract specifics and raw log

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
func (it *GlintdoughTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughTransferBatch)
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
		it.Event = new(GlintdoughTransferBatch)
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
func (it *GlintdoughTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughTransferBatch represents a TransferBatch event raised by the Glintdough contract.
type GlintdoughTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Glintdough *GlintdoughFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*GlintdoughTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughTransferBatchIterator{contract: _Glintdough.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Glintdough *GlintdoughFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *GlintdoughTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughTransferBatch)
				if err := _Glintdough.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Glintdough *GlintdoughFilterer) ParseTransferBatch(log types.Log) (*GlintdoughTransferBatch, error) {
	event := new(GlintdoughTransferBatch)
	if err := _Glintdough.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Glintdough contract.
type GlintdoughTransferSingleIterator struct {
	Event *GlintdoughTransferSingle // Event containing the contract specifics and raw log

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
func (it *GlintdoughTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughTransferSingle)
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
		it.Event = new(GlintdoughTransferSingle)
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
func (it *GlintdoughTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughTransferSingle represents a TransferSingle event raised by the Glintdough contract.
type GlintdoughTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Glintdough *GlintdoughFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*GlintdoughTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughTransferSingleIterator{contract: _Glintdough.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Glintdough *GlintdoughFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *GlintdoughTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughTransferSingle)
				if err := _Glintdough.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Glintdough *GlintdoughFilterer) ParseTransferSingle(log types.Log) (*GlintdoughTransferSingle, error) {
	event := new(GlintdoughTransferSingle)
	if err := _Glintdough.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Glintdough contract.
type GlintdoughURIIterator struct {
	Event *GlintdoughURI // Event containing the contract specifics and raw log

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
func (it *GlintdoughURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughURI)
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
		it.Event = new(GlintdoughURI)
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
func (it *GlintdoughURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughURI represents a URI event raised by the Glintdough contract.
type GlintdoughURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Glintdough *GlintdoughFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*GlintdoughURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughURIIterator{contract: _Glintdough.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Glintdough *GlintdoughFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *GlintdoughURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughURI)
				if err := _Glintdough.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Glintdough *GlintdoughFilterer) ParseURI(log types.Log) (*GlintdoughURI, error) {
	event := new(GlintdoughURI)
	if err := _Glintdough.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Glintdough contract.
type GlintdoughUnpausedIterator struct {
	Event *GlintdoughUnpaused // Event containing the contract specifics and raw log

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
func (it *GlintdoughUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughUnpaused)
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
		it.Event = new(GlintdoughUnpaused)
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
func (it *GlintdoughUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughUnpaused represents a Unpaused event raised by the Glintdough contract.
type GlintdoughUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Glintdough *GlintdoughFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GlintdoughUnpausedIterator, error) {

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GlintdoughUnpausedIterator{contract: _Glintdough.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Glintdough *GlintdoughFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GlintdoughUnpaused) (event.Subscription, error) {

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughUnpaused)
				if err := _Glintdough.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Glintdough *GlintdoughFilterer) ParseUnpaused(log types.Log) (*GlintdoughUnpaused, error) {
	event := new(GlintdoughUnpaused)
	if err := _Glintdough.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughVoucherAddedIterator is returned from FilterVoucherAdded and is used to iterate over the raw logs and unpacked data for VoucherAdded events raised by the Glintdough contract.
type GlintdoughVoucherAddedIterator struct {
	Event *GlintdoughVoucherAdded // Event containing the contract specifics and raw log

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
func (it *GlintdoughVoucherAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughVoucherAdded)
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
		it.Event = new(GlintdoughVoucherAdded)
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
func (it *GlintdoughVoucherAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughVoucherAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughVoucherAdded represents a VoucherAdded event raised by the Glintdough contract.
type GlintdoughVoucherAdded struct {
	TokenId      *big.Int
	Token        common.Address
	Denomination *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoucherAdded is a free log retrieval operation binding the contract event 0x7b8024ce6a9ab55c17f27580f8dc379c6d6beee58d682c382fa50ae828925f74.
//
// Solidity: event VoucherAdded(uint256 indexed tokenId, address indexed token, uint256 denomination)
func (_Glintdough *GlintdoughFilterer) FilterVoucherAdded(opts *bind.FilterOpts, tokenId []*big.Int, token []common.Address) (*GlintdoughVoucherAddedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "VoucherAdded", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughVoucherAddedIterator{contract: _Glintdough.contract, event: "VoucherAdded", logs: logs, sub: sub}, nil
}

// WatchVoucherAdded is a free log subscription operation binding the contract event 0x7b8024ce6a9ab55c17f27580f8dc379c6d6beee58d682c382fa50ae828925f74.
//
// Solidity: event VoucherAdded(uint256 indexed tokenId, address indexed token, uint256 denomination)
func (_Glintdough *GlintdoughFilterer) WatchVoucherAdded(opts *bind.WatchOpts, sink chan<- *GlintdoughVoucherAdded, tokenId []*big.Int, token []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "VoucherAdded", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughVoucherAdded)
				if err := _Glintdough.contract.UnpackLog(event, "VoucherAdded", log); err != nil {
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

// ParseVoucherAdded is a log parse operation binding the contract event 0x7b8024ce6a9ab55c17f27580f8dc379c6d6beee58d682c382fa50ae828925f74.
//
// Solidity: event VoucherAdded(uint256 indexed tokenId, address indexed token, uint256 denomination)
func (_Glintdough *GlintdoughFilterer) ParseVoucherAdded(log types.Log) (*GlintdoughVoucherAdded, error) {
	event := new(GlintdoughVoucherAdded)
	if err := _Glintdough.contract.UnpackLog(event, "VoucherAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughVoucherMintedSingleIterator is returned from FilterVoucherMintedSingle and is used to iterate over the raw logs and unpacked data for VoucherMintedSingle events raised by the Glintdough contract.
type GlintdoughVoucherMintedSingleIterator struct {
	Event *GlintdoughVoucherMintedSingle // Event containing the contract specifics and raw log

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
func (it *GlintdoughVoucherMintedSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughVoucherMintedSingle)
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
		it.Event = new(GlintdoughVoucherMintedSingle)
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
func (it *GlintdoughVoucherMintedSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughVoucherMintedSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughVoucherMintedSingle represents a VoucherMintedSingle event raised by the Glintdough contract.
type GlintdoughVoucherMintedSingle struct {
	Minter   common.Address
	TokenId  *big.Int
	Quantity *big.Int
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoucherMintedSingle is a free log retrieval operation binding the contract event 0xdb940767cbdf4dfef9c754bd48608769ec95df94cd670e0399e9523b0d52239b.
//
// Solidity: event VoucherMintedSingle(address indexed minter, uint256 indexed tokenId, uint256 quantity, address indexed to)
func (_Glintdough *GlintdoughFilterer) FilterVoucherMintedSingle(opts *bind.FilterOpts, minter []common.Address, tokenId []*big.Int, to []common.Address) (*GlintdoughVoucherMintedSingleIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "VoucherMintedSingle", minterRule, tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughVoucherMintedSingleIterator{contract: _Glintdough.contract, event: "VoucherMintedSingle", logs: logs, sub: sub}, nil
}

// WatchVoucherMintedSingle is a free log subscription operation binding the contract event 0xdb940767cbdf4dfef9c754bd48608769ec95df94cd670e0399e9523b0d52239b.
//
// Solidity: event VoucherMintedSingle(address indexed minter, uint256 indexed tokenId, uint256 quantity, address indexed to)
func (_Glintdough *GlintdoughFilterer) WatchVoucherMintedSingle(opts *bind.WatchOpts, sink chan<- *GlintdoughVoucherMintedSingle, minter []common.Address, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "VoucherMintedSingle", minterRule, tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughVoucherMintedSingle)
				if err := _Glintdough.contract.UnpackLog(event, "VoucherMintedSingle", log); err != nil {
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

// ParseVoucherMintedSingle is a log parse operation binding the contract event 0xdb940767cbdf4dfef9c754bd48608769ec95df94cd670e0399e9523b0d52239b.
//
// Solidity: event VoucherMintedSingle(address indexed minter, uint256 indexed tokenId, uint256 quantity, address indexed to)
func (_Glintdough *GlintdoughFilterer) ParseVoucherMintedSingle(log types.Log) (*GlintdoughVoucherMintedSingle, error) {
	event := new(GlintdoughVoucherMintedSingle)
	if err := _Glintdough.contract.UnpackLog(event, "VoucherMintedSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughVoucherRedeemedIterator is returned from FilterVoucherRedeemed and is used to iterate over the raw logs and unpacked data for VoucherRedeemed events raised by the Glintdough contract.
type GlintdoughVoucherRedeemedIterator struct {
	Event *GlintdoughVoucherRedeemed // Event containing the contract specifics and raw log

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
func (it *GlintdoughVoucherRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughVoucherRedeemed)
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
		it.Event = new(GlintdoughVoucherRedeemed)
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
func (it *GlintdoughVoucherRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughVoucherRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughVoucherRedeemed represents a VoucherRedeemed event raised by the Glintdough contract.
type GlintdoughVoucherRedeemed struct {
	TokenId  *big.Int
	Redeemer common.Address
	To       common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoucherRedeemed is a free log retrieval operation binding the contract event 0xea6258ffcc941e51a8b38b05fdbef327d28c7015ac1edb675e13086a0439f222.
//
// Solidity: event VoucherRedeemed(uint256 indexed tokenId, address indexed redeemer, address to, uint256 quantity)
func (_Glintdough *GlintdoughFilterer) FilterVoucherRedeemed(opts *bind.FilterOpts, tokenId []*big.Int, redeemer []common.Address) (*GlintdoughVoucherRedeemedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "VoucherRedeemed", tokenIdRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughVoucherRedeemedIterator{contract: _Glintdough.contract, event: "VoucherRedeemed", logs: logs, sub: sub}, nil
}

// WatchVoucherRedeemed is a free log subscription operation binding the contract event 0xea6258ffcc941e51a8b38b05fdbef327d28c7015ac1edb675e13086a0439f222.
//
// Solidity: event VoucherRedeemed(uint256 indexed tokenId, address indexed redeemer, address to, uint256 quantity)
func (_Glintdough *GlintdoughFilterer) WatchVoucherRedeemed(opts *bind.WatchOpts, sink chan<- *GlintdoughVoucherRedeemed, tokenId []*big.Int, redeemer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "VoucherRedeemed", tokenIdRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughVoucherRedeemed)
				if err := _Glintdough.contract.UnpackLog(event, "VoucherRedeemed", log); err != nil {
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

// ParseVoucherRedeemed is a log parse operation binding the contract event 0xea6258ffcc941e51a8b38b05fdbef327d28c7015ac1edb675e13086a0439f222.
//
// Solidity: event VoucherRedeemed(uint256 indexed tokenId, address indexed redeemer, address to, uint256 quantity)
func (_Glintdough *GlintdoughFilterer) ParseVoucherRedeemed(log types.Log) (*GlintdoughVoucherRedeemed, error) {
	event := new(GlintdoughVoucherRedeemed)
	if err := _Glintdough.contract.UnpackLog(event, "VoucherRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughVoucherToggledIterator is returned from FilterVoucherToggled and is used to iterate over the raw logs and unpacked data for VoucherToggled events raised by the Glintdough contract.
type GlintdoughVoucherToggledIterator struct {
	Event *GlintdoughVoucherToggled // Event containing the contract specifics and raw log

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
func (it *GlintdoughVoucherToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughVoucherToggled)
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
		it.Event = new(GlintdoughVoucherToggled)
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
func (it *GlintdoughVoucherToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughVoucherToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughVoucherToggled represents a VoucherToggled event raised by the Glintdough contract.
type GlintdoughVoucherToggled struct {
	TokenId *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoucherToggled is a free log retrieval operation binding the contract event 0x44850097d5058840642ec07d1f10e064b3a8714a43b01f0a9a01a15fc577deb6.
//
// Solidity: event VoucherToggled(uint256 indexed tokenId, bool status)
func (_Glintdough *GlintdoughFilterer) FilterVoucherToggled(opts *bind.FilterOpts, tokenId []*big.Int) (*GlintdoughVoucherToggledIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "VoucherToggled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughVoucherToggledIterator{contract: _Glintdough.contract, event: "VoucherToggled", logs: logs, sub: sub}, nil
}

// WatchVoucherToggled is a free log subscription operation binding the contract event 0x44850097d5058840642ec07d1f10e064b3a8714a43b01f0a9a01a15fc577deb6.
//
// Solidity: event VoucherToggled(uint256 indexed tokenId, bool status)
func (_Glintdough *GlintdoughFilterer) WatchVoucherToggled(opts *bind.WatchOpts, sink chan<- *GlintdoughVoucherToggled, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "VoucherToggled", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughVoucherToggled)
				if err := _Glintdough.contract.UnpackLog(event, "VoucherToggled", log); err != nil {
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

// ParseVoucherToggled is a log parse operation binding the contract event 0x44850097d5058840642ec07d1f10e064b3a8714a43b01f0a9a01a15fc577deb6.
//
// Solidity: event VoucherToggled(uint256 indexed tokenId, bool status)
func (_Glintdough *GlintdoughFilterer) ParseVoucherToggled(log types.Log) (*GlintdoughVoucherToggled, error) {
	event := new(GlintdoughVoucherToggled)
	if err := _Glintdough.contract.UnpackLog(event, "VoucherToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlintdoughVouchersMintedBatchIterator is returned from FilterVouchersMintedBatch and is used to iterate over the raw logs and unpacked data for VouchersMintedBatch events raised by the Glintdough contract.
type GlintdoughVouchersMintedBatchIterator struct {
	Event *GlintdoughVouchersMintedBatch // Event containing the contract specifics and raw log

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
func (it *GlintdoughVouchersMintedBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlintdoughVouchersMintedBatch)
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
		it.Event = new(GlintdoughVouchersMintedBatch)
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
func (it *GlintdoughVouchersMintedBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlintdoughVouchersMintedBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlintdoughVouchersMintedBatch represents a VouchersMintedBatch event raised by the Glintdough contract.
type GlintdoughVouchersMintedBatch struct {
	Minter     common.Address
	TokenIds   []*big.Int
	Quantities []*big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVouchersMintedBatch is a free log retrieval operation binding the contract event 0x433b9aa33bcc3cdcfd6a434c1936e7034c1d70a0bd333d45aad92d5f423c07e0.
//
// Solidity: event VouchersMintedBatch(address indexed minter, uint256[] tokenIds, uint256[] quantities, address indexed to)
func (_Glintdough *GlintdoughFilterer) FilterVouchersMintedBatch(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*GlintdoughVouchersMintedBatchIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.FilterLogs(opts, "VouchersMintedBatch", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GlintdoughVouchersMintedBatchIterator{contract: _Glintdough.contract, event: "VouchersMintedBatch", logs: logs, sub: sub}, nil
}

// WatchVouchersMintedBatch is a free log subscription operation binding the contract event 0x433b9aa33bcc3cdcfd6a434c1936e7034c1d70a0bd333d45aad92d5f423c07e0.
//
// Solidity: event VouchersMintedBatch(address indexed minter, uint256[] tokenIds, uint256[] quantities, address indexed to)
func (_Glintdough *GlintdoughFilterer) WatchVouchersMintedBatch(opts *bind.WatchOpts, sink chan<- *GlintdoughVouchersMintedBatch, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Glintdough.contract.WatchLogs(opts, "VouchersMintedBatch", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlintdoughVouchersMintedBatch)
				if err := _Glintdough.contract.UnpackLog(event, "VouchersMintedBatch", log); err != nil {
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

// ParseVouchersMintedBatch is a log parse operation binding the contract event 0x433b9aa33bcc3cdcfd6a434c1936e7034c1d70a0bd333d45aad92d5f423c07e0.
//
// Solidity: event VouchersMintedBatch(address indexed minter, uint256[] tokenIds, uint256[] quantities, address indexed to)
func (_Glintdough *GlintdoughFilterer) ParseVouchersMintedBatch(log types.Log) (*GlintdoughVouchersMintedBatch, error) {
	event := new(GlintdoughVouchersMintedBatch)
	if err := _Glintdough.contract.UnpackLog(event, "VouchersMintedBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
