// Package auth contains the api authentication mechanism.
package auth

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// EIP191Auth implements ethereum signed message authentication middleware for gin rest api
// For auth, relayer should pass in eth signed message following eip-191 with the message
// as the current unix timestamp in seconds
// i.e. signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix())))
// so that full auth header string: auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
// see: https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_sign
func EIP191Auth(account string, deadline int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		// check given account is a valid address
		if common.HexToAddress(account).Hex() != account {
			err := fmt.Errorf("account is invalid address")
			c.JSON(400, gin.H{"msg": err})
			return
		}

		// parse <timestamp>:<signature>
		s := strings.Split(auth, ":")
		if len(s) != 2 {
			err := fmt.Errorf("invalid authorization header format")
			c.JSON(400, gin.H{"msg": err})
			return
		}

		// check timestamp is not older than given deadline
		timestamp, err := strconv.ParseInt(s[0], 10, 64)
		if err != nil {
			err := fmt.Errorf("invalid timestamp in authorization")
			c.JSON(400, gin.H{"msg": err})
			return
		} else if timestamp < deadline {
			err := fmt.Errorf("authorization too old")
			c.JSON(401, gin.H{"msg": err}) // Unauthorized
			return
		}

		// check signature matches eth signed data of timestamp signed by given account
		signature, err := hexutil.Decode(s[1])
		if err != nil {
			err = fmt.Errorf("signature not hex encoded in authorization")
			c.JSON(400, gin.H{"msg": err})
			return
		}

		data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(s[0])) + s[0]
		digest := crypto.Keccak256([]byte(data)) // TODO: check []byte(data) ok

		recovered, err := crypto.SigToPub(digest, signature)
		if err != nil {
			err = fmt.Errorf("failed to recover signer from authorization")
			c.JSON(400, gin.H{"msg": err})
			return
		}

		signer := crypto.PubkeyToAddress(*recovered).Hex()
		if signer != account {
			err = fmt.Errorf("account != signer from authorization")
			c.JSON(401, gin.H{"msg": err}) // Unauthorized
			return
		}
	}
}
