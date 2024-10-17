// Package botmd provides the bot server. Here botmd=cmd not markdown.
// nolint: forcetypeassert, mnd, cyclop
package botmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hako/durafmt"
	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/contrib/opbot/internal"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/chaindata"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	rfqClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
)

func (b *Bot) requiresSignoz(definition *slacker.CommandDefinition) *slacker.CommandDefinition {
	if b.signozEnabled {
		return definition
	}
	return &slacker.CommandDefinition{
		Command:     definition.Command,
		Description: fmt.Sprintf("normally this would \"%s\", but signoz is not configured", definition.Description),
		Examples:    definition.Examples,
		Handler: func(ctx *slacker.CommandContext) {
			_, err := ctx.Response().Reply("cannot run command: signoz is not configured")
			if err != nil {
				log.Println(err)
			}
		},
	}
}

// nolint: traceCommand, gocognit
// TODO: add trace middleware.
func (b *Bot) traceCommand() *slacker.CommandDefinition {
	return b.requiresSignoz(&slacker.CommandDefinition{
		Command:     "trace {tags} {order}",
		Description: "find a transaction in signoz",
		Examples: []string{
			"trace transaction_id:0x1234@serviceName:rfq",
			"trace transaction_id:0x1234@serviceName:rfq a",
			"trace transaction_id:0x1234@serviceName:rfq asc",
		},
		Handler: func(ctx *slacker.CommandContext) {
			tags := stripLinks(ctx.Request().Param("tags"))
			splitTags := strings.Split(tags, "@")
			if len(splitTags) == 0 {
				_, err := ctx.Response().Reply("please provide tags in a key:value format")
				if err != nil {
					log.Println(err)
				}
				return
			}

			searchMap := make(map[string]string)
			for _, combinedTag := range splitTags {
				tag := strings.Split(combinedTag, ":")
				if len(tag) != 2 {
					_, err := ctx.Response().Reply("please provide tags in a key:value format")
					if err != nil {
						log.Println(err)
					}
					return
				}
				searchMap[tag[0]] = tag[1]
			}

			// search for the transaction
			res, err := b.signozClient.SearchTraces(ctx.Context(), signoz.Last3Hr, searchMap)
			if err != nil {
				b.logger.Errorf(ctx.Context(), "error searching for the transaction: %v", err)
				_, err := ctx.Response().Reply("error searching for the transaction")
				if err != nil {
					log.Println(err)
				}
				return
			}

			if res.Status != "success" || res.Data.ContextTimeout || len(res.Data.Result) != 1 {
				_, err := ctx.Response().Reply(fmt.Sprintf("error searching for the transaction %s", res.Data.ContextTimeoutMessage))
				if err != nil {
					log.Println(err)
				}
				return
			}

			traceList := res.Data.Result[0].List
			if len(traceList) == 0 {
				_, err := ctx.Response().Reply("no transaction found")
				if err != nil {
					log.Println(err)
				}
				return
			}

			order := strings.ToLower(ctx.Request().Param("order"))
			isAscending := order == "a" || order == "asc"
			if isAscending {
				sort.Slice(traceList, func(i, j int) bool {
					return traceList[i].Timestamp.Before(traceList[j].Timestamp)
				})
			}

			slackBlocks := []slack.Block{slack.NewHeaderBlock(slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf("Traces for %s", tags), false, false))}

			for _, results := range traceList {
				trace := results.Data["traceID"].(string)
				spanID := results.Data["spanID"].(string)
				serviceName := results.Data["serviceName"].(string)

				url := fmt.Sprintf("%s/trace/%s?spanId=%s", b.cfg.SignozBaseURL, trace, spanID)
				traceName := fmt.Sprintf("<%s|%s>", url, results.Data["name"].(string))

				relativeTime := durafmt.Parse(time.Since(results.Timestamp)).LimitFirstN(1).String()

				slackBlocks = append(slackBlocks, slack.NewSectionBlock(nil, []*slack.TextBlockObject{
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*Name*: %s", traceName),
					},
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*Service*: %s", serviceName),
					},
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*When*: %s", fmt.Sprintf("%s ago", relativeTime)),
					},
				}, nil))
			}

			_, err = ctx.Response().ReplyBlocks(slackBlocks, slacker.WithUnfurlLinks(false))
			if err != nil {
				log.Println(err)
			}
		},
	})
}

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

// nolint: gocognit, cyclop.
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

			fastBridgeContract, err := b.makeFastBridge(ctx.Context(), rawRequest)
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

			nonce, err := b.submitter.SubmitTransaction(ctx.Context(), big.NewInt(int64(rawRequest.Bridge.OriginChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
				tx, err = fastBridgeContract.Refund(transactor, common.Hex2Bytes(rawRequest.QuoteRequestRaw))
				if err != nil {
					return nil, fmt.Errorf("error submitting refund: %w", err)
				}
				return tx, nil
			})
			if err != nil {
				log.Printf("error submitting refund: %v\n", err)
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
					}
					return nil
				},
				retry.WithMaxAttempts(5),
				retry.WithMaxAttemptTime(30*time.Second),
			)
			if err != nil {
				b.logger.Errorf(ctx.Context(), "error fetching quote request: %v", err)
				_, err := ctx.Response().Reply(fmt.Sprintf("refund submitted with nonce %d", nonce))
				if err != nil {
					log.Println(err)
				}
				return
			}

			_, err = ctx.Response().Reply(fmt.Sprintf("refund submitted: %s", toTXSlackLink(status.TxHash().String(), uint32(rawRequest.Bridge.OriginChainID))))
			if err != nil {
				log.Println(err)
			}
		},
	}
}

func (b *Bot) makeFastBridge(ctx context.Context, req *internal.GetRFQByTxIDResponse) (*fastbridge.FastBridge, error) {
	client, err := rfqClient.NewUnauthenticatedClient(b.handler, b.cfg.RFQApiURL)
	if err != nil {
		return nil, fmt.Errorf("error creating rfq client: %w", err)
	}

	contracts, err := client.GetRFQContracts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching rfq contracts: %w", err)
	}

	chainClient, err := b.rpcClient.GetChainClient(ctx, int(req.Bridge.OriginChainID))
	if err != nil {
		return nil, fmt.Errorf("error getting chain client: %w", err)
	}

	contractAddress, ok := contracts.Contracts[uint32(req.Bridge.OriginChainID)]
	if !ok {
		return nil, errors.New("contract address not found")
	}

	fastBridgeHandle, err := fastbridge.NewFastBridge(common.HexToAddress(contractAddress), chainClient)
	if err != nil {
		return nil, fmt.Errorf("error creating fast bridge: %w", err)
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
