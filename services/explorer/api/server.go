package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/services/explorer/api/cache"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"time"

	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/graphql/client"
	gqlServer "github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// Config contains the config for the api.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16
	// Address is the address of the database
	Address string
	// ScribeURL is the url of the scribe service
	ScribeURL string
}

const cacheRehydrationInterval = 300

var logger = log.Logger("explorer-api")

// Start starts the api server.
func Start(ctx context.Context, cfg Config) error {
	router := ginhelper.New(logger)
	// initialize the database
	consumerDB, err := InitDB(ctx, cfg.Address, true)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	// get the fetcher
	fetcher := fetcher.NewFetcher(client.NewClient(http.DefaultClient, cfg.ScribeURL))

	// response cache
	responseCache, err := cache.NewApiCacheService()
	if err != nil {
		return err
	}
	gqlServer.EnableGraphql(router, consumerDB, *fetcher, responseCache)

	fmt.Printf("started graphiql gqlServer on port: http://localhost:%d/graphiql\n", cfg.HTTPPort)

	ticker := time.NewTicker(cacheRehydrationInterval * time.Second)
	g, ctx := errgroup.WithContext(ctx)
	client := gqlClient.NewClient(http.DefaultClient, fmt.Sprintf("http://localhost:%d/graphql", cfg.HTTPPort))
	rehydrateCache(ctx, client, responseCache)
	// refill cache
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				rehydrateCache(ctx, client, responseCache)
			}
		}
	}()

	g.Go(func() error {
		connection := baseServer.Server{}
		err = connection.ListenAndServe(ctx, fmt.Sprintf(":%d", cfg.HTTPPort), router)
		if err != nil {
			return fmt.Errorf("could not start gqlServer: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// InitDB initializes a database given a database type and path.
func InitDB(ctx context.Context, address string, readOnly bool) (db.ConsumerDB, error) {
	if address == "default" {
		cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
		if cleanup == nil {
			return nil, fmt.Errorf("clickhouse spin up failure, no open port found: %w", err)
		}
		if port == nil || err != nil {
			cleanup()
			return nil, fmt.Errorf("clickhouse spin up failure, no open port found: %w", err)
		}
		address = "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	}
	clickhouseDB, err := sql.OpenGormClickhouse(ctx, address, readOnly)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	return clickhouseDB, nil
}

// TODO make this nicer
func rehydrateCache(parentCtx context.Context, client *gqlClient.Client, service cache.Service) error {
	fmt.Println("rehydrating Cache")
	totalVolumeType := model.StatisticTypeTotalVolumeUsd
	totalFeeType := model.StatisticTypeTotalFeeUsd
	countAddressType := model.StatisticTypeCountAddresses
	countTxType := model.StatisticTypeCountTransactions

	allPlatformType := model.PlatformAll
	bridgeType := model.PlatformBridge
	swapType := model.PlatformSwap
	messagingType := model.PlatformMessageBus

	volumeType := model.DailyStatisticTypeVolume
	feeType := model.DailyStatisticTypeFee
	txType := model.DailyStatisticTypeTransactions
	addrType := model.DailyStatisticTypeAddresses

	monthType := model.DurationPastMonth
	yearType := model.DurationPastYear
	allTimeType := model.DurationAllTime
	g, ctx := errgroup.WithContext(parentCtx)
	g.Go(func() error {
		statsVolAll, err := client.GetAmountStatistic(ctx, totalVolumeType, &allPlatformType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, ALL, ALL_TIME, , , ", handleJsonAmountStat(statsVolAll))
		statsFeeAll, err := client.GetAmountStatistic(ctx, totalFeeType, &allPlatformType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_FEE_USD, ALL, ALL_TIME, , , ", handleJsonAmountStat(statsFeeAll))
		statsAddrAll, err := client.GetAmountStatistic(ctx, countAddressType, &allPlatformType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_ADDRESSES, ALL, ALL_TIME, , , ", handleJsonAmountStat(statsAddrAll))
		statsTxAll, err := client.GetAmountStatistic(ctx, countTxType, &allPlatformType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, ALL, ALL_TIME, , , ", handleJsonAmountStat(statsTxAll))
		return nil
	})
	g.Go(func() error {
		statsVolBridge, err := client.GetAmountStatistic(ctx, totalVolumeType, &bridgeType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, BRIDGE, ALL_TIME, , , ", handleJsonAmountStat(statsVolBridge))
		statsFeeBridge, err := client.GetAmountStatistic(ctx, totalFeeType, &bridgeType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_FEE_USD, BRIDGE, ALL_TIME, , , ", handleJsonAmountStat(statsFeeBridge))
		statsAddrBridge, err := client.GetAmountStatistic(ctx, countAddressType, &bridgeType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_ADDRESSES, BRIDGE, ALL_TIME, , , ", handleJsonAmountStat(statsAddrBridge))
		statsTxBridge, err := client.GetAmountStatistic(ctx, countTxType, &bridgeType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, BRIDGE, ALL_TIME, , , ", handleJsonAmountStat(statsTxBridge))
		return nil
	})
	g.Go(func() error {
		statsVolSwap, err := client.GetAmountStatistic(ctx, totalVolumeType, &swapType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, SWAP, ALL_TIME, , , ", handleJsonAmountStat(statsVolSwap))
		statsFeeSwap, err := client.GetAmountStatistic(ctx, totalFeeType, &swapType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_FEE_USD, SWAP, ALL_TIME, , , ", handleJsonAmountStat(statsFeeSwap))
		statsAddrSwap, err := client.GetAmountStatistic(ctx, countAddressType, &swapType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_ADDRESSES, SWAP, ALL_TIME, , , ", handleJsonAmountStat(statsAddrSwap))
		statsTxSwap, err := client.GetAmountStatistic(ctx, countTxType, &swapType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, SWAP, ALL_TIME, , , ", handleJsonAmountStat(statsTxSwap))
		return nil
	})
	g.Go(func() error {
		statsFeeMsg, err := client.GetAmountStatistic(ctx, totalFeeType, &messagingType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, TOTAL_FEE_USD, MESSAGE_BUS, ALL_TIME, , , ", handleJsonAmountStat(statsFeeMsg))
		statsAddrMsg, err := client.GetAmountStatistic(ctx, countAddressType, &messagingType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_ADDRESSES, MESSAGE_BUS, ALL_TIME, , , ", handleJsonAmountStat(statsAddrMsg))
		statsTxMsg, err := client.GetAmountStatistic(ctx, countTxType, &messagingType, &allTimeType, nil, nil, nil)
		if err != nil {
			return err
		}
		service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, MESSAGE_BUS, ALL_TIME, , , ", handleJsonAmountStat(statsTxMsg))
		return nil
	})
	g.Go(func() error {
		dailyVolMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, ALL", handleJsonDailyStat(dailyVolMonth))
		dailyFeeMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, ALL", handleJsonDailyStat(dailyFeeMonth))
		dailyTxMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, ALL", handleJsonDailyStat(dailyTxMonth))
		dailyAddrMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, ALL", handleJsonDailyStat(dailyAddrMonth))
		return nil
	})
	g.Go(func() error {
		dailyVolYear, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_YEAR, ALL", handleJsonDailyStat(dailyVolYear))
		dailyFeeYear, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_YEAR, ALL", handleJsonDailyStat(dailyFeeYear))
		dailyTxYear, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_YEAR, ALL", handleJsonDailyStat(dailyTxYear))
		dailyAddrYear, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_YEAR, ALL", handleJsonDailyStat(dailyAddrYear))
		return nil
	})
	g.Go(func() error {
		dailyVolAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, ALL_TIME, ALL", handleJsonDailyStat(dailyVolAllTime))
		dailyFeeAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, ALL_TIME, ALL", handleJsonDailyStat(dailyFeeAllTime))
		dailyTxAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, ALL_TIME, ALL", handleJsonDailyStat(dailyTxAllTime))
		dailyAddrAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &allPlatformType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, ALL_TIME, ALL", handleJsonDailyStat(dailyAddrAllTime))
		return nil
	})

	g.Go(func() error {
		dailyVolMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, BRIDGE", handleJsonDailyStat(dailyVolMonthBridge))
		dailyFeeMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, BRIDGE", handleJsonDailyStat(dailyFeeMonthBridge))
		dailyTxMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, BRIDGE", handleJsonDailyStat(dailyTxMonthBridge))
		dailyAddrMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, BRIDGE", handleJsonDailyStat(dailyAddrMonthBridge))
		return nil
	})
	g.Go(func() error {
		dailyVolYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_YEAR, BRIDGE", handleJsonDailyStat(dailyVolYearBridge))
		dailyFeeYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_YEAR, BRIDGE", handleJsonDailyStat(dailyFeeYearBridge))
		dailyTxYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_YEAR, BRIDGE", handleJsonDailyStat(dailyTxYearBridge))
		dailyAddrYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_YEAR, BRIDGE", handleJsonDailyStat(dailyAddrYearBridge))
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, ALL_TIME, BRIDGE", handleJsonDailyStat(dailyVolAllTimeBridge))
		dailyFeeAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, ALL_TIME, BRIDGE", handleJsonDailyStat(dailyFeeAllTimeBridge))
		dailyTxAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, ALL_TIME, BRIDGE", handleJsonDailyStat(dailyTxAllTimeBridge))
		dailyAddrAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &bridgeType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, ALL_TIME, BRIDGE", handleJsonDailyStat(dailyAddrAllTimeBridge))
		return nil
	})

	g.Go(func() error {
		dailyVolMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, SWAP", handleJsonDailyStat(dailyVolMonthSwap))
		dailyFeeMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, SWAP", handleJsonDailyStat(dailyFeeMonthSwap))
		dailyTxMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, SWAP", handleJsonDailyStat(dailyTxMonthSwap))
		dailyAddrMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, SWAP", handleJsonDailyStat(dailyAddrMonthSwap))
		return nil
	})
	g.Go(func() error {
		dailyVolYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_YEAR, SWAP", handleJsonDailyStat(dailyVolYearSwap))
		dailyFeeYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_YEAR, SWAP", handleJsonDailyStat(dailyFeeYearSwap))
		dailyTxYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_YEAR, SWAP", handleJsonDailyStat(dailyTxYearSwap))
		dailyAddrYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_YEAR, SWAP", handleJsonDailyStat(dailyAddrYearSwap))
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , VOLUME, ALL_TIME, SWAP", handleJsonDailyStat(dailyVolAllTimeSwap))
		dailyFeeAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, ALL_TIME, SWAP", handleJsonDailyStat(dailyFeeAllTimeSwap))
		dailyTxAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, ALL_TIME, SWAP", handleJsonDailyStat(dailyTxAllTimeSwap))
		dailyAddrAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &swapType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, ALL_TIME, SWAP", handleJsonDailyStat(dailyAddrAllTimeSwap))
		return nil
	})

	g.Go(func() error {
		dailyFeeMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, MESSAGE_BUS", handleJsonDailyStat(dailyFeeMonthMessageBus))
		dailyTxMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, MESSAGE_BUS", handleJsonDailyStat(dailyTxMonthMessageBus))
		dailyAddrMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, MESSAGE_BUS", handleJsonDailyStat(dailyAddrMonthMessageBus))
		return nil
	})
	g.Go(func() error {
		dailyFeeYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_YEAR, MESSAGE_BUS", handleJsonDailyStat(dailyFeeYearMessageBus))
		dailyTxYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_YEAR, MESSAGE_BUS", handleJsonDailyStat(dailyTxYearMessageBus))
		dailyAddrYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_YEAR, MESSAGE_BUS", handleJsonDailyStat(dailyAddrYearMessageBus))
		return nil
	})
	g.Go(func() error {
		dailyFeeAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , FEE, ALL_TIME, MESSAGE_BUS", handleJsonDailyStat(dailyFeeAllTimeMessageBus))
		dailyTxAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, ALL_TIME, MESSAGE_BUS", handleJsonDailyStat(dailyTxAllTimeMessageBus))
		dailyAddrAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &messagingType)
		if err != nil {
			return err
		}
		service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, ALL_TIME, MESSAGE_BUS", handleJsonDailyStat(dailyAddrAllTimeMessageBus))
		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}
	return nil
}

func handleJsonAmountStat(r *gqlClient.GetAmountStatistic) *model.ValueResult {
	var res *model.ValueResult
	jsonRes, _ := json.Marshal(r.Response)
	json.Unmarshal(jsonRes, &res)
	return res
}

func handleJsonDailyStat(r *gqlClient.GetDailyStatisticsByChain) []*model.DateResultByChain {
	var res []*model.DateResultByChain
	jsonRes, _ := json.Marshal(r.Response)
	json.Unmarshal(jsonRes, &res)
	return res
}
