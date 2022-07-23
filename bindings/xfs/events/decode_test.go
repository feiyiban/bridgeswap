package events

import (
	"fmt"
	"testing"

	"bridgeswap/sdk/xfsgo/common"
)

var BRIDGETOKENABI = `{"events":{"0x06976fbcd8c1677b395e07d3c109af81384fea5da3dbce252de69d194c97d4bb":{"name":"BridgeTransferInEvent","argc":6,"args":[{"name":"BankAddress","type":"CTypeAddress"},{"name":"DepositorAddress","type":"CTypeAddress"},{"name":"ContractAddress","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"},{"name":"FromChainId","type":"CTypeUint256"},{"name":"ToChainId","type":"CTypeUint256"}]},"0x0db84af94380c5f8b90cb557eabb4534c01d07dd3775ba4ab0d91cb94ec898a6":{"name":"BridgeTransferOutEvent","argc":7,"args":[{"name":"BankAddress","type":"CTypeAddress"},{"name":"DepositorAddress","type":"CTypeAddress"},{"name":"ContractAddress","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"},{"name":"FromChainId","type":"CTypeUint256"},{"name":"ToChainId","type":"CTypeUint256"},{"name":"ToAddress","type":"CTypeString"}]}},"methods":{"0x0000000000000000000000000000000000000000000000000000000000000000":{"name":"Create","argc":4,"args":[{"name":"","type":"CTypeString"},{"name":"","type":"CTypeString"},{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x40d2e803bbe3904840cece9e5997f40fa8cc16db20357606ddb096193dcdcdce":{"name":"TransferIn","argc":3,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xb6b0adfd06b688ce0bebc64dfdeefdddbc38a71a7ca0de0869499668245e41d4":{"name":"TransferOut","argc":4,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeString"},{"name":"","type":"CTypeUint256"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xd4e97969f51b509b07924c476ae8983687162f7e7e5068603ca2780ebf18451b":{"name":"GetChainId","argc":0,"args":[],"return_type":"CTypeUint256"}}}`
var XFSTOKENABI = `{"events":{"0x011f3f6cad22a2efb7ae1c8e484a01b51b384f4dee84a4c4e776d1abbc7ad9e4":{"name":"StdTokenTransferEvent","argc":3,"args":[{"name":"From","type":"CTypeAddress"},{"name":"To","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]},"0x473c5d5f7beec0001489b92d9fa4b05bca8c1b7bce26ee9de20e410b27db2b3b":{"name":"StdTokenApprovalEvent","argc":3,"args":[{"name":"Owner","type":"CTypeAddress"},{"name":"Spender","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]}},"methods":{"0x0000000000000000000000000000000000000000000000000000000000000000":{"name":"Create","argc":4,"args":[{"name":"","type":"CTypeString"},{"name":"","type":"CTypeString"},{"name":"","type":"CTypeUint8"},{"name":"","type":"CTypeUint256"}],"return_type":"error"},"0x03f4098a5e9d39a5104a34a4a19025c1cefd1551ebaedb871af3bcc12250f295":{"name":"GetTotalSupply","argc":0,"args":[],"return_type":"CTypeUint256"},"0x1162f326f21ac342307b16730bc30e1cfb6fd35acfd527a2d6adf39d44b56522":{"name":"GetName","argc":0,"args":[],"return_type":"CTypeString"},"0x2561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a":{"name":"TransferFrom","argc":3,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x2b99b4d70435e95aac2a5b0fe9f1286ac033b46dec731828b7de558a17d869f5":{"name":"Allowance","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0x6007acbe30b2cd98703e83350ea665c06009fcd51f26dd73b309294235f45f21":{"name":"Approve","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x61945fbcd9ffbebe7dcf1ec99e8bd195e6b235295dbe5f84df2f8a2b72174e1c":{"name":"BalanceOf","argc":1,"args":[{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0x926c5b4314047434601585221956407b3818b5f1cda70febda6e25d04f204e4c":{"name":"Burn","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xb00e879ffa3a243b7b964ad38c7616c1ee2d027dc05a6c11569a737f9a700a53":{"name":"GetDecimals","argc":0,"args":[],"return_type":"CTypeUint8"},"0xced97cc4a377b5b4386d9c67bc4f4e14febb561903a27409ce7a2886368b75bb":{"name":"Mint","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xd24b7074b8d5ee3e7e0a471901324f6870e175419253f5e497b42272f6919234":{"name":"GetSymbol","argc":0,"args":[],"return_type":"CTypeString"},"0xdde8bef78cbb720683fa1fe76bfb900592099ed4346ed995bcbc514e9aa67256":{"name":"Transfer","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"}}}`

func Test_Decode(t *testing.T) {
	es, err := JSON(BRIDGETOKENABI)
	if err != nil {
		t.Error(err)
	}
	testBridgeEventHash := "0x06976fbcd8c1677b395e07d3c109af81384fea5da3dbce252de69d194c97d4bb"
	testBridgeEventValue := "7b2262616e6b41646472657373223a223031313866306135356564626532646166333961326463326265376161353530323462666131316235383361613130633065222c226465706f7369746f7241646472657373223a223031306534356239366339316138336233613064303335373239666536633765646230666635633165303035303536346433222c22636f6e747261637441646472657373223a223031363065383361356533666439653737306162313136643762663762643034643032313763643261646166306162636235222c2276616c7565223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303061222c2266726f6d436861696e4964223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303033222c22746f436861696e4964223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303033227d"
	enval, err := es.Decode(testBridgeEventHash, testBridgeEventValue)
	if err != nil {
		t.Error(err)
	}
	bs, err := common.MarshalIndent(enval)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bs))
}
