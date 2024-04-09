// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const AlphabetVMStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"test/mocks/AlphabetVM.sol:AlphabetVM\",\"label\":\"oracle\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_contract(IPreimageOracle)1001\"}],\"types\":{\"t_contract(IPreimageOracle)1001\":{\"encoding\":\"inplace\",\"label\":\"contract IPreimageOracle\",\"numberOfBytes\":\"20\"}}}"

var AlphabetVMStorageLayout = new(solc.StorageLayout)

var AlphabetVMDeployedBin = "0x608060405234801561001057600080fd5b50600436106100365760003560e01c80637dc0d1d01461003b578063e14ced3214610085575b600080fd5b60005461005b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610098610093366004610213565b6100a6565b60405190815260200161007c565b600080600060087f0000000000000000000000000000000000000000000000000000000000000000901b600889896040516100e2929190610287565b6040518091039020901b03610108576000915061010187890189610297565b9050610127565b610114878901896102b0565b90925090508161012381610301565b9250505b81610133826001610339565b604080516020810193909352820152606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001798975050505050505050565b60008083601f8401126101dc57600080fd5b50813567ffffffffffffffff8111156101f457600080fd5b60208301915083602082850101111561020c57600080fd5b9250929050565b60008060008060006060868803121561022b57600080fd5b853567ffffffffffffffff8082111561024357600080fd5b61024f89838a016101ca565b9097509550602088013591508082111561026857600080fd5b50610275888289016101ca565b96999598509660400135949350505050565b8183823760009101908152919050565b6000602082840312156102a957600080fd5b5035919050565b600080604083850312156102c357600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610332576103326102d2565b5060010190565b6000821982111561034c5761034c6102d2565b50019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(AlphabetVMStorageLayoutJSON), AlphabetVMStorageLayout); err != nil {
		panic(err)
	}

	layouts["AlphabetVM"] = AlphabetVMStorageLayout
	deployedBytecodes["AlphabetVM"] = AlphabetVMDeployedBin
	immutableReferences["AlphabetVM"] = true
}
