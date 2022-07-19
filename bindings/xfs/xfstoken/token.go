package xfstoken

const (
	CREATE         = "0x0000000000000000000000000000000000000000000000000000000000000000"
	NAME           = "0x1162f326f21ac342307b16730bc30e1cfb6fd35acfd527a2d6adf39d44b56522"
	SYMBOL         = "0xd24b7074b8d5ee3e7e0a471901324f6870e175419253f5e497b42272f6919234"
	GETDECIMALS    = "0xb00e879ffa3a243b7b964ad38c7616c1ee2d027dc05a6c11569a737f9a700a53"
	GETTOTALSUPPLY = "0x03f4098a5e9d39a5104a34a4a19025c1cefd1551ebaedb871af3bcc12250f295"
	BALANCEOF      = "0x61945fbcd9ffbebe7dcf1ec99e8bd195e6b235295dbe5f84df2f8a2b72174e1c"

	MINT = "0xced97cc4a377b5b4386d9c67bc4f4e14febb561903a27409ce7a2886368b75bb"
	BURN = "0x926c5b4314047434601585221956407b3818b5f1cda70febda6e25d04f204e4c"

	APPROVE      = "0x6007acbe30b2cd98703e83350ea665c06009fcd51f26dd73b309294235f45f21"
	ALLOWANCE    = "0x2b99b4d70435e95aac2a5b0fe9f1286ac033b46dec731828b7de558a17d869f5"
	TRANSFERFROM = "0x2561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a"
)

var XFSTOKENABI = `{"events":{"0x011f3f6cad22a2efb7ae1c8e484a01b51b384f4dee84a4c4e776d1abbc7ad9e4":{"name":"StdTokenTransferEvent","argc":3,"args":[{"name":"From","type":"CTypeAddress"},{"name":"To","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]},"0x473c5d5f7beec0001489b92d9fa4b05bca8c1b7bce26ee9de20e410b27db2b3b":{"name":"StdTokenApprovalEvent","argc":3,"args":[{"name":"Owner","type":"CTypeAddress"},{"name":"Spender","type":"CTypeAddress"},{"name":"Value","type":"CTypeUint256"}]}},"methods":{"0x0000000000000000000000000000000000000000000000000000000000000000":{"name":"Create","argc":4,"args":[{"name":"","type":"CTypeString"},{"name":"","type":"CTypeString"},{"name":"","type":"CTypeUint8"},{"name":"","type":"CTypeUint256"}],"return_type":"error"},"0x03f4098a5e9d39a5104a34a4a19025c1cefd1551ebaedb871af3bcc12250f295":{"name":"GetTotalSupply","argc":0,"args":[],"return_type":"CTypeUint256"},"0x1162f326f21ac342307b16730bc30e1cfb6fd35acfd527a2d6adf39d44b56522":{"name":"GetName","argc":0,"args":[],"return_type":"CTypeString"},"0x2561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a":{"name":"TransferFrom","argc":3,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x2b99b4d70435e95aac2a5b0fe9f1286ac033b46dec731828b7de558a17d869f5":{"name":"Allowance","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0x6007acbe30b2cd98703e83350ea665c06009fcd51f26dd73b309294235f45f21":{"name":"Approve","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0x61945fbcd9ffbebe7dcf1ec99e8bd195e6b235295dbe5f84df2f8a2b72174e1c":{"name":"BalanceOf","argc":1,"args":[{"name":"","type":"CTypeAddress"}],"return_type":"CTypeUint256"},"0x926c5b4314047434601585221956407b3818b5f1cda70febda6e25d04f204e4c":{"name":"Burn","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xb00e879ffa3a243b7b964ad38c7616c1ee2d027dc05a6c11569a737f9a700a53":{"name":"GetDecimals","argc":0,"args":[],"return_type":"CTypeUint8"},"0xced97cc4a377b5b4386d9c67bc4f4e14febb561903a27409ce7a2886368b75bb":{"name":"Mint","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"},"0xd24b7074b8d5ee3e7e0a471901324f6870e175419253f5e497b42272f6919234":{"name":"GetSymbol","argc":0,"args":[],"return_type":"CTypeString"},"0xdde8bef78cbb720683fa1fe76bfb900592099ed4346ed995bcbc514e9aa67256":{"name":"Transfer","argc":2,"args":[{"name":"","type":"CTypeAddress"},{"name":"","type":"CTypeUint256"}],"return_type":"CTypeBool"}}}`
var XFSTOKENBin = "0xd02301"
