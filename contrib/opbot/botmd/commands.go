// Package botmd provides the bot server. Here botmd=cmd not markdown.
// nolint: forcetypeassert, mnd, cyclop
package botmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/chaindata"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	rfqClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

// nolint: gocognit
func (b *Bot) rfqLookupCommand() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command:     "rfq <tx>",
		Description: "find a rfq transaction by either tx hash or txid from the rfq-indexer api",
		Examples: []string{
			"rfq 0x30f96b45ba689c809f7e936c140609eb31c99b182bef54fccf49778716a7e1ca",
		},
		Handler: func(ctx *slacker.CommandContext) {
			tx := stripLinks(ctx.Request().Param("tx"))

			res, status, err := b.rfqClient.GetRFQ(ctx.Context(), tx)
			if err != nil {
				b.logger.Errorf(ctx.Context(), "error fetching quote request: %v", err)
				_, err := ctx.Response().Reply(fmt.Sprintf("error fetching quote request %s", err.Error()))
				if err != nil {
					log.Println(err)
				}
				return
			}

			var slackBlocks []slack.Block

			objects := []*slack.TextBlockObject{
				{
					Type: slack.MarkdownType,
					Text: fmt.Sprintf("*Relayer*: %s", res.BridgeRelay.Relayer),
				},
				{
					Type: slack.MarkdownType,
					Text: fmt.Sprintf("*Status*: %s", status),
				},
				{
					Type: slack.MarkdownType,
					Text: fmt.Sprintf("*TxID*: %s", toExplorerSlackLink(res.Bridge.TransactionID)),
				},
				{
					Type: slack.MarkdownType,
					//nolint: gosec
					Text: fmt.Sprintf("*OriginTxHash*: %s", toTXSlackLink(res.BridgeRequest.TransactionHash, uint32(res.Bridge.OriginChainID))),
				},
				{
					Type: slack.MarkdownType,
					Text: fmt.Sprintf("*Estimated Tx Age*: %s", humanize.Time(time.Unix(res.BridgeRelay.BlockTimestamp, 0))),
				},
			}

			if status == "Requested" {
				objects = append(objects, &slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "*DestTxHash*: not available",
				})
			} else {
				//nolint: gosec
				objects = append(objects, &slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: fmt.Sprintf("*DestTxHash*: %s", toTXSlackLink(res.BridgeRelay.TransactionHash, uint32(res.Bridge.DestChainID))),
				})
			}

			slackBlocks = append(slackBlocks, slack.NewSectionBlock(nil, objects, nil))

			_, err = ctx.Response().ReplyBlocks(slackBlocks, slacker.WithUnfurlLinks(false))
			if err != nil {
				log.Println(err)
			}
		},
	}
}

// nolint: gocognit, cyclop, gosec.
func (b *Bot) rfqRefund() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command:     "refund <tx>",
		Description: "refund a quote request",
		Examples:    []string{"refund 0x1234"},
		Handler: func(ctx *slacker.CommandContext) {
			tx := stripLinks(ctx.Request().Param("tx"))

			if len(tx) == 0 {
				_, err := ctx.Response().Reply("please provide a tx hash")
				if err != nil {
					log.Println(err)
				}
				return
			}

			rawRequest, _, err := b.rfqClient.GetRFQ(ctx.Context(), tx)
			if err != nil {
				b.logger.Errorf(ctx.Context(), "error fetching quote request: %v", err)
				_, err := ctx.Response().Reply("error fetching quote request")
				if err != nil {
					log.Println(err)
				}
				return
			}

			//nolint: gosec
			fastBridgeContractOrigin, err := b.makeFastBridge(ctx.Context(), uint32(rawRequest.Bridge.OriginChainID))
			if err != nil {
				_, err := ctx.Response().Reply(err.Error())
				if err != nil {
					log.Println(err)
				}
				return
			}

			isScreened, err := b.screener.ScreenAddress(ctx.Context(), rawRequest.Bridge.Sender)
			if err != nil {
				_, err := ctx.Response().Reply("error screening address")
				if err != nil {
					log.Println(err)
				}
				return
			}
			if isScreened {
				_, err := ctx.Response().Reply("address cannot be refunded")
				if err != nil {
					log.Println(err)
				}
				return
			}

			//nolint:gosec
			fastBridgeContractDest, err := b.makeFastBridge(ctx.Context(), uint32(rawRequest.Bridge.DestChainID))
			if err != nil {
				_, err := ctx.Response().Reply(err.Error())
				if err != nil {
					log.Println(err)
				}
				return
			}
			txBz, err := core.BytesToArray(common.Hex2Bytes(rawRequest.Bridge.TransactionID[2:]))
			if err != nil {
				_, err := ctx.Response().Reply("error converting tx id")
				if err != nil {
					log.Println(err)
				}
				return
			}
			isRelayed, err := fastBridgeContractDest.BridgeRelays(nil, txBz)
			if err != nil {
				_, err := ctx.Response().Reply("error fetching bridge relays")
				if err != nil {
					log.Println(err)
				}
				return
			}
			if isRelayed {
				_, err := ctx.Response().Reply("transaction has already been relayed")
				if err != nil {
					log.Println(err)
				}
				return
			}

			nonce, err := b.submitter.SubmitTransaction(
				ctx.Context(),
				big.NewInt(int64(rawRequest.Bridge.OriginChainID)),
				func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
					tx, err = fastBridgeContractOrigin.Refund(transactor, common.Hex2Bytes(rawRequest.Bridge.Request[2:]))
					if err != nil {
						return nil, fmt.Errorf("error submitting refund: %w", err)
					}
					return tx, nil
				})
			if err != nil {
				log.Printf("error submitting refund: %v\n", err)
				_, err := ctx.Response().Reply("error submitting refund")
				if err != nil {
					log.Println(err)
				}
				return
			}

			var status submitter.SubmissionStatus
			err = retry.WithBackoff(
				ctx.Context(),
				func(ctx context.Context) error {
					status, err = b.submitter.GetSubmissionStatus(ctx, big.NewInt(int64(rawRequest.Bridge.OriginChainID)), nonce)
					if err != nil || !status.HasTx() {
						b.logger.Errorf(ctx, "error fetching quote request: %v", err)
						return fmt.Errorf("error fetching quote request: %w", err)
					} else if !status.HasTx() {
						return fmt.Errorf("no transaction hash found yet")
					}
					return nil
				},
				retry.WithMaxTotalTime(1*time.Minute),
			)

			if err != nil {
				b.logger.Errorf(ctx.Context(), "error fetching quote request: %v", err)
				_, err := ctx.Response().Reply(fmt.Sprintf("refund submitted with nonce %d", nonce))
				if err != nil {
					b.logger.Errorf(ctx.Context(), "error fetching quote request: %v", err)
				}
				return
			}

			//nolint: gosec
			_, err = ctx.Response().Reply(fmt.Sprintf("refund submitted: %s", toTXSlackLink(status.TxHash().String(), uint32(rawRequest.Bridge.OriginChainID))))
			if err != nil {
				log.Println(err)
			}

		},
	}
}

func (b *Bot) makeFastBridge(ctx context.Context, chainID uint32) (*fastbridge.FastBridge, error) {
	client, err := rfqClient.NewUnauthenticatedClient(b.handler, b.cfg.RFQApiURL)
	if err != nil {
		return nil, fmt.Errorf("error creating rfq client: %w", err)
	}

	contracts, err := client.GetRFQContracts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching rfq contracts: %w", err)
	}

	chainClient, err := b.rpcClient.GetChainClient(ctx, int(chainID))
	if err != nil {
		return nil, fmt.Errorf("error getting chain client for chain ID %d: %w", chainID, err)
	}

	contractAddress, ok := contracts.Contracts[chainID]
	if !ok {
		return nil, fmt.Errorf("no contract address for chain ID")
	}

	fastBridgeHandle, err := fastbridge.NewFastBridge(common.HexToAddress(contractAddress), chainClient)
	if err != nil {
		return nil, fmt.Errorf("error creating fast bridge for chain ID %d: %w", chainID, err)
	}

	return fastBridgeHandle, nil

}

func toExplorerSlackLink(ogHash string) string {
	rfqHash := strings.ToUpper(ogHash)
	// cut off 0x
	if strings.HasPrefix(rfqHash, "0X") {
		rfqHash = strings.ToLower(rfqHash[2:])
	}

	return fmt.Sprintf("<https://explorer.synapseprotocol.com/tx/%s|%s>", rfqHash, ogHash)
}

// produce a salck link if the explorer exists.
func toTXSlackLink(txHash string, chainID uint32) string {
	url := chaindata.ToTXLink(int64(chainID), txHash)
	if url == "" {
		return txHash
	}

	// TODO: remove when we can contorl unfurl
	return fmt.Sprintf("<%s|%s>", url, txHash)
}

func stripLinks(input string) string {
	linkRegex := regexp.MustCompile(`<https?://[^|>]+\|([^>]+)>`)
	return linkRegex.ReplaceAllString(input, "$1")
}
