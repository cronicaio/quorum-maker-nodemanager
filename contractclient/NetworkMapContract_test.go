package contractclient

import (
	"testing"
	"synechron.com/NodeManagerGo/contracthandler"
	"synechron.com/NodeManagerGo/client"
	"fmt"
	"synechron.com/NodeManagerGo/util"
	"time"
)

func TestRegisterNode(t *testing.T) {

	cp := getContractParam()

	ec := client.EthClient{"http://localhost:22000"}

	nmc := NetworkMapContractClient{ec, cp}

	txRec := nmc.RegisterNode(
		"c5f4b39a1c40c5affc99ec6f7be64e7c20d78c96ac55f20ee1156ce87175732a9c2b518aa6897f1590ea78b911be0c6a524d8496a420107651251048332bb04e",
		"SynechronBank",
		"Bank")

	if txRec == "" {
		t.Error("Error Registering Node")
	}
}

func TestUpdateNode(t *testing.T) {

	cp := getContractParam()

	ec := client.EthClient{"http://localhost:22000"}

	nmc := NetworkMapContractClient{ec, cp}

	txRec := nmc.UpdateNode(
		"c5f4b39a1c40c5affc99ec6f7be64e7c20d78c96ac55f20ee1156ce87175732a9c2b518aa6897f1590ea78b911be0c6a524d8496a420107651251048332bb04e",
		"BB&T",
		"Bank")

	if txRec == "" {
		t.Error("Error Updating Node")
	}
}

func TestGetNodeDetails(t *testing.T) {

	defer util.TotalTime(time.Now().Nanosecond())
	cp := getContractParam()

	ec := client.EthClient{"http://localhost:22000"}

	nmc := NetworkMapContractClient{ec, cp}

	nd := nmc.GetNodeDetails(1)

	fmt.Println(nd.Name)
}

func getContractParam() contracthandler.ContractParam {
	return contracthandler.ContractParam{
		"0x2c049a350bc1284a662de7296d79c8c486867bdc",
		"0xf8bc1fc5e20c554315cbeb7d5a00b5635b512dd8",
		"",
		nil,
	}
}