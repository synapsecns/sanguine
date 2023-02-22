package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/services/explorer/api/cache"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"

	"net/http"

	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/graphql/client"
	gqlServer "github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"golang.org/x/sync/errgroup"
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

const cacheRehydrationInterval = 1800

var logger = log.Logger("explorer-api")

// Start starts the api server.
func Start(ctx context.Context, cfg Config) error {
	router := ginhelper.New(logger)
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("could not get hostname %w", err)
	}
	// initialize the database
	consumerDB, err := InitDB(ctx, cfg.Address, true)
	if err != nil {
		return fmt.Errorf("could not initialize database: %w", err)
	}

	// get the fetcher
	fetcher := fetcher.NewFetcher(client.NewClient(http.DefaultClient, cfg.ScribeURL))

	// response cache
	responseCache, err := cache.NewAPICacheService()
	if err != nil {
		return fmt.Errorf("error creating api cache service, %w", err)
	}
	gqlServer.EnableGraphql(router, consumerDB, *fetcher, responseCache)

	fmt.Printf("started graphiql gqlServer on port: http://%s:%d/graphiql\n", hostname, cfg.HTTPPort)

	ticker := time.NewTicker(cacheRehydrationInterval * time.Second)
	defer ticker.Stop()
	first := make(chan bool, 1)
	first <- true
	g, ctx := errgroup.WithContext(ctx)
	url := fmt.Sprintf("http://%s:%d/graphql", hostname, cfg.HTTPPort)
	client := gqlClient.NewClient(http.DefaultClient, url)

	// refill cache
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				err = rehydrateCache(ctx, client, responseCache)
				if err != nil {
					logger.Warnf("rehydration failed: %s", err)
				}
			case <-first:
				// buffer to wait for everything to get initialized
				time.Sleep(10 * time.Second)
				err = rehydrateCache(ctx, client, responseCache)
				if err != nil {
					logger.Errorf("initial rehydration failed: %s", err)
				}
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

// TODO make this nicer. make a yaml of the queries needed for rehydration w/refresh rate and iterate on that.
//
// nolint:dupl,gocognit,cyclop,maintidx
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
	yearType := model.DurationPast3Months
	allTimeType := model.DurationPast6Months

	//dontUseMv := false
	useMv := true

	g, ctx := errgroup.WithContext(parentCtx)
	g.Go(func() error {
		statsVolAll, err := client.GetAmountStatistic(ctx, totalVolumeType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, ALL, ALL_TIME, , , ", handleJSONAmountStat(statsVolAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeAll, err := client.GetAmountStatistic(ctx, totalFeeType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, ALL, ALL_TIME, , , ", handleJSONAmountStat(statsFeeAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrAll, err := client.GetAmountStatistic(ctx, countAddressType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, ALL, ALL_TIME, , , ", handleJSONAmountStat(statsAddrAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxAll, err := client.GetAmountStatistic(ctx, countTxType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, ALL, ALL_TIME, , , ", handleJSONAmountStat(statsTxAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		statsVolBridge, err := client.GetAmountStatistic(ctx, totalVolumeType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, BRIDGE, ALL_TIME, , , ", handleJSONAmountStat(statsVolBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeBridge, err := client.GetAmountStatistic(ctx, totalFeeType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, BRIDGE, ALL_TIME, , , ", handleJSONAmountStat(statsFeeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrBridge, err := client.GetAmountStatistic(ctx, countAddressType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, BRIDGE, ALL_TIME, , , ", handleJSONAmountStat(statsAddrBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxBridge, err := client.GetAmountStatistic(ctx, countTxType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, BRIDGE, ALL_TIME, , , ", handleJSONAmountStat(statsTxBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		statsVolSwap, err := client.GetAmountStatistic(ctx, totalVolumeType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, SWAP, ALL_TIME, , , ", handleJSONAmountStat(statsVolSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeSwap, err := client.GetAmountStatistic(ctx, totalFeeType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, SWAP, ALL_TIME, , , ", handleJSONAmountStat(statsFeeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrSwap, err := client.GetAmountStatistic(ctx, countAddressType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, SWAP, ALL_TIME, , , ", handleJSONAmountStat(statsAddrSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxSwap, err := client.GetAmountStatistic(ctx, countTxType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, SWAP, ALL_TIME, , , ", handleJSONAmountStat(statsTxSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		statsFeeMsg, err := client.GetAmountStatistic(ctx, totalFeeType, &messagingType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, MESSAGE_BUS, ALL_TIME, , , ", handleJSONAmountStat(statsFeeMsg))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrMsg, err := client.GetAmountStatistic(ctx, countAddressType, &messagingType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, MESSAGE_BUS, ALL_TIME, , , ", handleJSONAmountStat(statsAddrMsg))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxMsg, err := client.GetAmountStatistic(ctx, countTxType, &messagingType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, MESSAGE_BUS, ALL_TIME, , , ", handleJSONAmountStat(statsTxMsg))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, ALL", handleJSONDailyStat(dailyVolMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, ALL", handleJSONDailyStat(dailyFeeMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, ALL", handleJSONDailyStat(dailyTxMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, ALL", handleJSONDailyStat(dailyAddrMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYear, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, ALL", handleJSONDailyStat(dailyVolYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYear, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, ALL", handleJSONDailyStat(dailyFeeYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYear, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, ALL", handleJSONDailyStat(dailyTxYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYear, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, ALL", handleJSONDailyStat(dailyAddrYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, ALL", handleJSONDailyStat(dailyVolAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, ALL", handleJSONDailyStat(dailyFeeAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, ALL", handleJSONDailyStat(dailyTxAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, ALL", handleJSONDailyStat(dailyAddrAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		dailyVolMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, BRIDGE", handleJSONDailyStat(dailyVolMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, BRIDGE", handleJSONDailyStat(dailyFeeMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, BRIDGE", handleJSONDailyStat(dailyTxMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, BRIDGE", handleJSONDailyStat(dailyAddrMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, BRIDGE", handleJSONDailyStat(dailyVolYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, BRIDGE", handleJSONDailyStat(dailyFeeYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, BRIDGE", handleJSONDailyStat(dailyTxYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, BRIDGE", handleJSONDailyStat(dailyAddrYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, BRIDGE", handleJSONDailyStat(dailyVolAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, BRIDGE", handleJSONDailyStat(dailyFeeAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, BRIDGE", handleJSONDailyStat(dailyTxAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, BRIDGE", handleJSONDailyStat(dailyAddrAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		dailyVolMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, SWAP", handleJSONDailyStat(dailyVolMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, SWAP", handleJSONDailyStat(dailyFeeMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, SWAP", handleJSONDailyStat(dailyTxMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}

		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, SWAP", handleJSONDailyStat(dailyAddrMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &yearType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, SWAP", handleJSONDailyStat(dailyVolYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, SWAP", handleJSONDailyStat(dailyFeeYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, SWAP", handleJSONDailyStat(dailyTxYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, SWAP", handleJSONDailyStat(dailyAddrYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &allTimeType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, SWAP", handleJSONDailyStat(dailyVolAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, SWAP", handleJSONDailyStat(dailyFeeAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, SWAP", handleJSONDailyStat(dailyTxAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, SWAP", handleJSONDailyStat(dailyAddrAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		dailyFeeMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, MESSAGE_BUS", handleJSONDailyStat(dailyFeeMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, MESSAGE_BUS", handleJSONDailyStat(dailyTxMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, MESSAGE_BUS", handleJSONDailyStat(dailyAddrMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyFeeYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &yearType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyFeeYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &yearType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyTxYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &yearType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyAddrYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyFeeAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &allTimeType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyFeeAllTimeMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &allTimeType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyTxAllTimeMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &allTimeType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, MESSAGE_BUS", handleJSONDailyStat(dailyAddrAllTimeMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error rehyrdrating cache, %w", err)
	}
	return nil
}

func handleJSONAmountStat(r *gqlClient.GetAmountStatistic) *model.ValueResult {
	var res *model.ValueResult
	jsonRes, err := json.Marshal(r.Response)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		return nil
	}
	return res
}

func handleJSONDailyStat(r *gqlClient.GetDailyStatisticsByChain) []*model.DateResultByChain {
	var res []*model.DateResultByChain
	jsonRes, err := json.Marshal(r.Response)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		return nil
	}
	return res
}
