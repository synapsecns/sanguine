// Package botmd provides the bot server. Here botmd=cmd not markdown.
// nolint: forcetypeassert, mnd, cyclop
package botmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hako/durafmt"
	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
	"github.com/synapsecns/sanguine/contrib/opbot/signoz"
	"github.com/synapsecns/sanguine/ethergo/chaindata"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"log"
	"strings"
	"sync"
	"time"
)

// nolint: traceCommand, gocognit
// TODO: add trace middleware.
func (b *Bot) traceCommand() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command:     "trace <tags>",
		Description: "find a transaction in signoz",
		Examples: []string{
			"trace transaction_id:0x1234 serviceName:rfq",
		},
		Handler: func(ctx *slacker.CommandContext) {
			tags := ctx.Request().Param("tags")
			splitTags := strings.Split(tags, " ")
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

			slackBlocks := []slack.Block{slack.NewHeaderBlock(slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf("Traces for %s", tags), false, false))}

			for _, results := range traceList {
				trace := results.Data["traceID"].(string)
				spanID := results.Data["spanID"].(string)
				serviceName := results.Data["serviceName"].(string)

				// anon.to is used here to prevent unfurl https://github.com/slack-io/slacker/issues/11
				url := fmt.Sprintf("https://anon.to/?%s/trace/%s?spanId=%s", b.cfg.SignozBaseURL, trace, spanID)
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

			_, err = ctx.Response().ReplyBlocks(slackBlocks)
			if err != nil {
				log.Println(err)
			}
		},
	}
}

func (b *Bot) rfqLookupCommand() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Command:     "rfq <tx>",
		Description: "find a rfq transaction by either tx hash or txid on all configured relayers",
		Examples: []string{
			"rfq 0x30f96b45ba689c809f7e936c140609eb31c99b182bef54fccf49778716a7e1ca",
		},
		Handler: func(ctx *slacker.CommandContext) {
			type Status struct {
				relayer string
				*relapi.GetQuoteRequestStatusResponse
			}

			var statuses []Status
			var sliceMux sync.Mutex

			if len(b.cfg.RelayerURLS) == 0 {
				_, err := ctx.Response().Reply("no relayer urls configured")
				if err != nil {
					log.Println(err)
				}
				return
			}

			tx := ctx.Request().Param("tx")

			var wg sync.WaitGroup
			// 2 routines per relayer, one for tx hashh one for tx id
			wg.Add(len(b.cfg.RelayerURLS) * 2)
			for _, relayer := range b.cfg.RelayerURLS {
				client := relapi.NewRelayerClient(b.handler, relayer)
				go func() {
					defer wg.Done()
					res, err := client.GetQuoteRequestStatusByTxHash(ctx.Context(), tx)
					if err != nil {
						log.Printf("error fetching quote request status by tx hash: %v\n", err)
						return
					}
					sliceMux.Lock()
					defer sliceMux.Unlock()
					statuses = append(statuses, Status{relayer: relayer, GetQuoteRequestStatusResponse: res})
				}()

				go func() {
					defer wg.Done()
					res, err := client.GetQuoteRequestStatusByTxID(ctx.Context(), tx)
					if err != nil {
						log.Printf("error fetching quote request status by tx id: %v\n", err)
						return
					}
					sliceMux.Lock()
					defer sliceMux.Unlock()
					statuses = append(statuses, Status{relayer: relayer, GetQuoteRequestStatusResponse: res})
				}()
			}
			wg.Wait()

			if len(statuses) == 0 {
				_, err := ctx.Response().Reply("no quote request found")
				if err != nil {
					log.Println(err)
				}
				return
			}

			var slackBlocks []slack.Block
			for _, status := range statuses {
				objects := []*slack.TextBlockObject{
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*Relayer*: %s", status.relayer),
					},
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*Status*: %s", status.Status),
					},
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*TxID*: %s", toExplorerSlackLink(status.TxID)),
					},
					{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*OriginTxHash*: %s", toTXSlackLink(status.OriginTxHash, status.OriginChainID)),
					},
				}

				if status.DestTxHash == (common.Hash{}).String() {
					objects = append(objects, &slack.TextBlockObject{
						Type: slack.MarkdownType,
						Text: "*DestTxHash*: not available",
					})
				} else {
					objects = append(objects, &slack.TextBlockObject{
						Type: slack.MarkdownType,
						Text: fmt.Sprintf("*DestTxHash*: %s", toTXSlackLink(status.DestTxHash, status.DestChainID)),
					})
				}

				slackBlocks = append(slackBlocks, slack.NewSectionBlock(nil, objects, nil))
			}

			_, err := ctx.Response().ReplyBlocks(slackBlocks)
			if err != nil {
				log.Println(err)
			}
		}}
}

func toExplorerSlackLink(ogHash string) string {
	rfqHash := strings.ToUpper(ogHash)
	// cut off 0x
	if strings.HasPrefix(rfqHash, "0x") {
		rfqHash = strings.ToLower(rfqHash[2:])
	}

	return fmt.Sprintf("<https://anon.to/?https://explorer.synapseprotocol.com/tx/%s|%s>", rfqHash, ogHash)
}

// produce a salck link if the explorer exists.
func toTXSlackLink(txHash string, chainID uint32) string {
	url := chaindata.ToTXLink(int64(chainID), txHash)
	if url == "" {
		return txHash
	}

	// TODO: remove when we can contorl unfurl
	return fmt.Sprintf("<https://anon.to/?%s|%s>", url, txHash)
}
