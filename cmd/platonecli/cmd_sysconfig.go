package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/PlatONEnetwork/PlatONE-Go/cmd/utils"

	"github.com/PlatONEnetwork/PlatONE-Go/core/vm"
	"gopkg.in/urfave/cli.v1"
)

const (
	txUseGas    = "use-gas" // IsTxUseGas
	txNotUseGas = "not-use"

	conAudit    = "audit"
	conNotAudit = "not-audit"

	checkPerm    = "with-perm"
	notCheckPerm = "without-perm"

	prodEmp    = "allow-empty"
	notProdEmp = "notallow-empty"
)

const (
	txGasLim             = "TxGasLimit"
	blockGasLim          = "BlockGasLimit"
	isTxUseGases         = "IsTxUseGas"
	isApprDeployedCon    = "IsApproveDeployedContract"
	isCheckConDeployPerm = "CheckContractDeployPermission"
	isProdEmptyBlock     = "IsProduceEmptyBlock"
	gasContract          = "GasContractName"
)

var (
	SysConfigCmd = cli.Command{
		Name:  "sysconfig",
		Usage: "Manage the system configurations",

		Subcommands: []cli.Command{
			setCfg,
			getCfg,
		},
	}

	setCfg = cli.Command{
		Name:   "set",
		Usage:  "set the system configurations",
		Action: setSysConfig,
		Flags:  sysConfigCmdFlags,
	}

	getCfg = cli.Command{
		Name:   "get",
		Usage:  "get the system configurations",
		Action: getSysConfig,
		Flags:  getSysConfigCmdFlags,
	}
)

func setSysConfig(c *cli.Context) {

	if c.NumFlags() > 1 {
		utils.Fatalf("please set one system configuration at a time")
	}

	txGasLimit := c.String(TxGasLimitFlags.Name)
	blockGasLimit := c.String(BlockGasLimitFlags.Name)
	isTxUseGas := c.String(IsTxUseGasFlags.Name)
	isApproveDeployedContract := c.String(IsApproveDeployedContractFlags.Name)
	isCheckContractDeployPermission := c.String(IsCheckContractDeployPermissionFlags.Name)
	isProduceEmptyBlock := c.String(IsProduceEmptyBlockFlags.Name)
	gasContractName := c.String(GasContractNameFlags.Name)

	setConfig(c, txGasLimit, txGasLim)
	setConfig(c, blockGasLimit, blockGasLim)
	setConfig(c, isTxUseGas, isTxUseGases)
	setConfig(c, isApproveDeployedContract, isApprDeployedCon)
	setConfig(c, isCheckContractDeployPermission, isCheckConDeployPerm)
	setConfig(c, isProduceEmptyBlock, isProdEmptyBlock)
	setConfig(c, gasContractName, gasContract)

}

func setConfig(c *cli.Context, param string, name string) {
	if !checkConfigParam(param, name) {
		return
	}

	newParam, err := sysConfigConvert(param, name)
	if err != nil {
		utils.Fatalf(err.Error())
	}

	funcName := "set" + name
	funcParams := CombineFuncParams(newParam)

	result := contractCall(c, funcParams, funcName, parameterManagementAddress)
	fmt.Printf("%s\n", result)
}

func checkConfigParam(param string, key string) bool {

	switch key {
	case "TxGasLimit":
		// number check
		num, err := strconv.ParseUint(param, 10, 0)
		if err != nil {
			return false
		}

		// param check
		isInRange := vm.TxGasLimitMinValue < num && vm.TxGasLimitMaxValue > num
		if !isInRange {
			fmt.Printf("the transaction gas limit should be within (%d, %d)\n",
				vm.TxGasLimitMinValue, vm.TxGasLimitMaxValue)
			return false
		}
	case "BlockGasLimit":
		num, err := strconv.ParseUint(param, 10, 0)
		if err != nil {
			return false
		}

		isInRange := vm.BlockGasLimitMinValue < num && vm.BlockGasLimitMaxValue > num
		if !isInRange {
			fmt.Printf("the block gas limit should be within (%d, %d)\n",
				vm.BlockGasLimitMinValue, vm.BlockGasLimitMaxValue)
			return false
		}
	default:
		if param == "" {
			return false
		}
	}

	return true
}

func getSysConfig(c *cli.Context) {

	txGasLimit := c.Bool(TxGasLimitFlags.Name)
	blockGasLimit := c.Bool(BlockGasLimitFlags.Name)
	isTxUseGas := c.Bool(IsTxUseGasFlags.Name)
	isApproveDeployedContract := c.Bool(IsApproveDeployedContractFlags.Name)
	isCheckContractDeployPermission := c.Bool(IsCheckContractDeployPermissionFlags.Name)
	isProduceEmptyBlock := c.Bool(IsProduceEmptyBlockFlags.Name)
	gasContractName := c.Bool(GasContractNameFlags.Name)

	getConfig(c, txGasLimit, txGasLim)
	getConfig(c, blockGasLimit, blockGasLim)
	getConfig(c, isTxUseGas, isTxUseGases)
	getConfig(c, isApproveDeployedContract, isApprDeployedCon)
	getConfig(c, isCheckContractDeployPermission, isCheckConDeployPerm)
	getConfig(c, isProduceEmptyBlock, isProdEmptyBlock)
	getConfig(c, gasContractName, gasContract)

}

func getConfig(c *cli.Context, isGet bool, name string) {

	funcName := "get" + name

	if isGet {
		result := contractCall(c, nil, funcName, parameterManagementAddress)
		result = sysconfigToString(result)
		str := sysConfigParsing(result, name)

		fmt.Printf("%s: %v\n", name, str)
	}
}

func sysconfigToString(param interface{}) interface{} {
	value := reflect.TypeOf(param)

	switch value.Kind() {
	case reflect.Uint64:
		return strconv.FormatUint(param.(uint64), 10)

	case reflect.Uint32:
		return strconv.FormatUint(uint64(param.(uint32)), 10)

	default:
		panic("not support, please add the corresponding type")
	}
}

func sysConfigParsing(param interface{}, paramName string) string {
	conv := genConfigConverter(paramName)
	return conv.parse(param)
}

func sysConfigConvert(param, paramName string) (string, error) {

	conv := genConfigConverter(paramName)
	resutl, err := conv.convert(param)
	if err != nil {
		return "", err
	}

	return resutl.(string), nil
}

func genConfigConverter(paramName string) *convert {
	var conv *convert

	switch paramName {
	case isTxUseGases:
		conv = newConvert(txUseGas, txNotUseGas, "1", "0", paramName)
	case isApprDeployedCon:
		conv = newConvert(conAudit, conNotAudit, "1", "0", paramName)
	case isCheckConDeployPerm: // node type
		conv = newConvert(checkPerm, notCheckPerm, "1", "0", paramName)
	case isProdEmptyBlock:
		conv = newConvert(prodEmp, notProdEmp, "1", "0", paramName)
	default:
		utils.Fatalf("invalid system configuration %v", paramName)
	}

	return conv
}