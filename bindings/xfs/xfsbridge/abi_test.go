package xfsbridge

// func Test_JSON(t *testing.T) {
// 	objClass, _ := JSON(NFTOKENABI)
// 	bs, _ := common.MarshalIndent(objClass.Events)
// 	fmt.Println(string(bs))
// }
// func Test_PackEventsName(t *testing.T) {
// 	abi, err := JSON(NFTOKENABI)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	// 0xd023022561555cf5bdc523a9cdcbb7810211f424a3477c8e4ae5773e6a37475247d78a01796e28058b703693d4c786f2b5d408706316364132acc3820000000000000001fe1623aecc1ee2c37f78c952e4954b8516400c7548963dd6000000000000000000000000000000000000000000000000000000000000000000000000000001
// 	// testStr := `{"from":"00000000000000000000000000000000000000000000000000","to":"01796e28058b703693d4c786f2b5d408706316364132acc382","value":"00000000000000000000000000000000000000000000003635c9adc5dea00000"}`
// 	testStr, _ := hex.DecodeString("7b2266726f6d223a223031373936653238303538623730333639336434633738366632623564343038373036333136333634313332616363333832222c22746f223a223031666531363233616563633165653263333766373863393532653439353462383531363430306337353438393633646436222c22746f6b656e4964223a2230303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303031227d")
// 	events, err := Str2Events(string(testStr))
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	eventsobj, err := abi.PackEventsName(events)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	eventsResp := Events2Map(eventsobj)
// 	bs, err := common.MarshalIndent(eventsResp)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	fmt.Println(string(bs))
// }