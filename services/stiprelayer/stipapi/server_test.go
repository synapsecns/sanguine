package stipapi_test

// func TestLoadAndConvertFeesAndRebates(t *testing.T) {
// 	// Load the config
// 	config, err := stipconfig.LoadConfig("../test-config.yaml")
// 	if err != nil {
// 		t.Fatalf("Failed to load config: %s", err)
// 	}

// 	// Convert the FeesAndRebates data
// 	jsonOutput := stipapi.ConvertFeesAndRebatesToJSON(config.FeesAndRebates)

// 	// Expected output (based on your requirements and sample data)
// 	expectedOutput := map[int]interface{}{
// 		42161: map[string]interface{}{
// 			"anyFromChain": map[string]interface{}{
// 				"SynapseBridge": map[string]interface{}{
// 					"nETH": map[string]int{"fee": 4, "rebate": 6},
// 					"WETH": map[string]int{"fee": 4, "rebate": 6},
// 					"nUSD": map[string]int{"fee": 4, "rebate": 6},
// 					"GMX":  map[string]int{"fee": 5, "rebate": 6},
// 				},
// 				"SynapseCCTP": map[string]interface{}{
// 					"USDC": map[string]int{"fee": 4, "rebate": 5},
// 				},
// 				"SynapseRFQ": map[string]interface{}{
// 					"USDC": map[string]int{"fee": 4, "rebate": 5},
// 				},
// 			},
// 		},
// 		1: map[string]interface{}{
// 			"42161": map[string]interface{}{
// 				"SynapseBridge": map[string]interface{}{
// 					"nETH": map[string]int{"fee": 10, "rebate": 12},
// 					"WETH": map[string]int{"fee": 10, "rebate": 12},
// 					"nUSD": map[string]int{"fee": 12, "rebate": 14},
// 				},
// 				"SynapseCCTP": map[string]interface{}{
// 					"USDC": map[string]int{"fee": 4, "rebate": 5},
// 				},
// 				"SynapseRFQ": map[string]interface{}{
// 					"USDC": map[string]int{"fee": 4, "rebate": 5},
// 				},
// 			},
// 		},
// 		43114: map[string]interface{}{
// 			"42161": map[string]interface{}{
// 				"SynapseBridge": map[string]interface{}{
// 					"GMX": map[string]int{"fee": 5, "rebate": 6},
// 				},
// 				"SynapseCCTP": map[string]interface{}{},
// 				"SynapseRFQ":  map[string]interface{}{},
// 			},
// 		},
// 	}

// }
