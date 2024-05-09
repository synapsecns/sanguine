package main

//
// import (
// 	"context"
// 	"fmt"
// 	"os"
//
// 	"github.com/synapsecns/sanguine/committee/db/connect"
// 	"github.com/synapsecns/sanguine/committee/p2p"
// 	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
// 	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
//
// 	"strings"
//
// 	"github.com/synapsecns/sanguine/core/dbcommon"
// 	"github.com/synapsecns/sanguine/core/metrics"
// )
//
// // stupid script to find out what the host addresses are
// func main() {
// 	managers := makeManagers()
// 	_ = combineHostAddresses(managers...)
//
// 	err := os.RemoveAll("./tmp")
// 	if err != nil {
// 		panic(err)
// 	}
//
// }
//
// func readPrivateKeys() ([]string, error) {
// 	files, err := os.ReadDir("./config")
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var keys []string
// 	for _, file := range files {
// 		if file.IsDir() {
// 			continue
// 		}
// 		if strings.HasPrefix(file.Name(), "committee-signer") {
// 			key, err := os.ReadFile("./config/" + file.Name())
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			keys = append(keys, string(key))
// 		}
// 	}
//
// 	return keys, nil
// }
//
// func makeManagers() []p2p.LibP2PManager {
// 	keys, err := readPrivateKeys()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var managers []p2p.LibP2PManager
// 	port := 12345
// 	for _, key := range keys {
// 		wall, err := wallet.FromHex(key[:len(key)-1])
// 		if err != nil {
// 			panic(err)
// 		}
// 		signer := localsigner.NewSigner(wall.PrivateKey())
// 		db, _ := connect.Connect(context.Background(), dbcommon.Sqlite, "tmp/getaddrs", metrics.Get())
// 		manager, _ := p2p.NewLibP2PManager(
// 			context.Background(), metrics.Get(), signer, db, port, true)
// 		port += 1
// 		managers = append(managers, manager)
// 	}
//
// 	fmt.Println(fmt.Sprint(len(managers)) + " managers created")
// 	return managers
// }
//
// func combineHostAddresses(hostLikes ...p2p.LibP2PManager) []string {
// 	var addresses []string
// 	for _, hl := range hostLikes {
// 		fmt.Println(hl.Host().ID())
// 		fmt.Println("----------------------------------")
// 		for _, addr := range hl.Host().Addrs() {
// 			address := fmt.Sprintf("%s/p2p/%s", addr, hl.Host().ID())
// 			addresses = append(addresses, address)
// 			fmt.Println("- " + address)
// 		}
// 		fmt.Println("==================================")
// 	}
// 	return addresses
// }
//
