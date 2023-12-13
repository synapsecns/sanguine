package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Client interface {
	CreateQuote(q *APIQuote) error
}

type clientImpl struct {
	rClient *resty.Client
}

// NewClient creates a new client for the RFQ quoting API
func NewClient(rfqURL string, reqSigner signer.Signer) (Client, error) {
	client := resty.New().
		SetBaseURL(rfqURL).
		OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
			// i.e. signature (hex encoded) = keccak(bytes.concat("\x19Ethereum Signed Message:\n", len(strconv.Itoa(time.Now().Unix()), strconv.Itoa(time.Now().Unix())))
			// so that full auth header string: auth = strconv.Itoa(time.Now().Unix()) + ":" + signature
			now := strconv.Itoa(int(time.Now().Unix()))
			data := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(now)) + now

			sig, err := reqSigner.SignMessage(request.Context(), []byte(data), true)
			if err != nil {
				return fmt.Errorf("failed to sign request: %w", err)
			}

			res := fmt.Sprintf("%s:%s", now, signer.EncodeHex(sig))

			request.SetHeader("Authorization", res)

			return nil
		})

	return clientImpl{
		rClient: client,
	}, nil
}

// CreateQuote creates a new quote in the RFQ quoting API
func (c clientImpl) CreateQuote(q *APIQuote) error {
	res, err := c.rClient.R().
		SetBody(q).
		Post(rest.QUOTE_ROUTE)
	// TODO: Figure out if there's anyhting to do with the response, right now it's result: 1
	_ = res

	return err
}

// APIQuote is the struct for the quote API.
type APIQuote struct {
	Relayer string `json:"relayer" binding:"required"`

	OriginChainID uint   `json:"origin_chain_id" binding:"required"`
	OriginToken   string `json:"origin_token" binding:"required"`
	OriginAmount  string `json:"origin_amount" binding:"required"`
	// TODO: origin amount norm should be a string
	OriginAmountNorm float64 `json:"origin_amount_norm" binding:"required"`
	OriginDecimals   uint8   `json:"origin_decimals" binding:"required"`

	DestChainID    uint    `json:"dest_chain_id" binding:"required"`
	DestToken      string  `json:"dest_token" binding:"required"`
	DestAmount     string  `json:"dest_amount" binding:"required"`
	DestAmountNorm float64 `json:"dest_amount_norm" binding:"required"`
	DestDecimals   uint8   `json:"dest_decimals" binding:"required"`

	Price     float64        `json:"price"` // price = destAmount <quote> / originAmount <base>
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
