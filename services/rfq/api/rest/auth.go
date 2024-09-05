package rest

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
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
func EIP191Auth(c *gin.Context, deadline int64) (accountRecovered common.Address, err error) {
	auth := c.Request.Header.Get("Authorization")

	// parse <timestamp>:<signature>
	s := strings.Split(auth, ":")
	if len(s) != 2 {
		err = fmt.Errorf("invalid authorization header format")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return common.Address{}, err
	}

	// check timestamp is not older than given deadline
	var timestamp int64
	timestamp, err = strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid timestamp in authorization")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return common.Address{}, err
	} else if timestamp < deadline {
		err = fmt.Errorf("authorization too old")
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err}) // Unauthorized
		return common.Address{}, err
	}

	// check signature matches eth signed data of timestamp signed by given account
	var signature []byte
	signature, err = hexutil.Decode(s[1])
	if err != nil {
		err = fmt.Errorf("signature not hex encoded in authorization")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return common.Address{}, err
	}

	data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(s[0])) + s[0]
	digest := crypto.Keccak256([]byte(data)) // TODO: check []byte(data) ok

	// identify byte position of the recovery ID ("v")
	vIndex := len(signature) - 1

	// Ethereum signatures commonly use v values of 27 or 28 for EIP-191.
	// Some libraries may return v as 0 or 1, so we need to handle both cases. crypto.SigToPub expects 0/1
	switch signature[vIndex] {
	//nolint: mnd
	case 27, 28:
		signature[vIndex] -= 27 // Normalize to 0 or 1 for crypto.SigToPub
	//nolint: mnd
	case 0, 1:
		// do nothing, already normalized
	default:
		err = fmt.Errorf("unrecognized recovery ID value - expected 0, 1, 27, or 28")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return common.Address{}, err
	}

	var recovered *ecdsa.PublicKey
	recovered, err = crypto.SigToPub(digest, signature)
	if err != nil {
		err = fmt.Errorf("failed to recover signer from authorization")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return common.Address{}, err
	}

	signer := crypto.PubkeyToAddress(*recovered)

	return signer, nil
}
