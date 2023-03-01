package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
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
//
// nolint:cyclop
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
	url := fmt.Sprintf("http://%s/graphql", net.JoinHostPort(hostname, fmt.Sprintf("%d", cfg.HTTPPort)))
	client := gqlClient.NewClient(http.DefaultClient, url)

	// refill cache
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				err = RehydrateCache(ctx, client, responseCache)
				if err != nil {
					logger.Warnf("rehydration failed: %s", err)
				}
			case <-first:
				// buffer to wait for everything to get initialized
				time.Sleep(10 * time.Second)
				err = RehydrateCache(ctx, client, responseCache)
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

// RehydrateCache rehydrates the cache.
//
// nolint:dupl,gocognit,cyclop,maintidx
func RehydrateCache(parentCtx context.Context, client *gqlClient.Client, service cache.Service) error {
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
	threeMonthType := model.DurationPast3Months
	sixMonthType := model.DurationPast6Months
	allTimeType := model.DurationAllTime

	// dontUseMv := false
	useMv := true

	g, ctx := errgroup.WithContext(parentCtx)
	g.Go(func() error {
		statsVolAll, err := client.GetAmountStatistic(ctx, totalVolumeType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, ALL, ALL_TIME, , , ", HandleJSONAmountStat(statsVolAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeAll, err := client.GetAmountStatistic(ctx, totalFeeType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, ALL, ALL_TIME, , , ", HandleJSONAmountStat(statsFeeAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrAll, err := client.GetAmountStatistic(ctx, countAddressType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, ALL, ALL_TIME, , , ", HandleJSONAmountStat(statsAddrAll))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxAll, err := client.GetAmountStatistic(ctx, countTxType, &allPlatformType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, ALL, ALL_TIME, , , ", HandleJSONAmountStat(statsTxAll))
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
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, BRIDGE, ALL_TIME, , , ", HandleJSONAmountStat(statsVolBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeBridge, err := client.GetAmountStatistic(ctx, totalFeeType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, BRIDGE, ALL_TIME, , , ", HandleJSONAmountStat(statsFeeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrBridge, err := client.GetAmountStatistic(ctx, countAddressType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, BRIDGE, ALL_TIME, , , ", HandleJSONAmountStat(statsAddrBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxBridge, err := client.GetAmountStatistic(ctx, countTxType, &bridgeType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, BRIDGE, ALL_TIME, , , ", HandleJSONAmountStat(statsTxBridge))
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
		err = service.CacheResponse("amountStatistic, TOTAL_VOLUME_USD, SWAP, ALL_TIME, , , ", HandleJSONAmountStat(statsVolSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsFeeSwap, err := client.GetAmountStatistic(ctx, totalFeeType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, SWAP, ALL_TIME, , , ", HandleJSONAmountStat(statsFeeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrSwap, err := client.GetAmountStatistic(ctx, countAddressType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, SWAP, ALL_TIME, , , ", HandleJSONAmountStat(statsAddrSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxSwap, err := client.GetAmountStatistic(ctx, countTxType, &swapType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, SWAP, ALL_TIME, , , ", HandleJSONAmountStat(statsTxSwap))
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
		err = service.CacheResponse("amountStatistic, TOTAL_FEE_USD, MESSAGE_BUS, ALL_TIME, , , ", HandleJSONAmountStat(statsFeeMsg))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsAddrMsg, err := client.GetAmountStatistic(ctx, countAddressType, &messagingType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_ADDRESSES, MESSAGE_BUS, ALL_TIME, , , ", HandleJSONAmountStat(statsAddrMsg))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		statsTxMsg, err := client.GetAmountStatistic(ctx, countTxType, &messagingType, &allTimeType, nil, nil, nil, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("amountStatistic, COUNT_TRANSACTIONS, MESSAGE_BUS, ALL_TIME, , , ", HandleJSONAmountStat(statsTxMsg))
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
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, ALL", HandleJSONDailyStat(dailyVolMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, ALL", HandleJSONDailyStat(dailyFeeMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, ALL", HandleJSONDailyStat(dailyTxMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonth, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, ALL", HandleJSONDailyStat(dailyAddrMonth))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYear, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &threeMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, ALL", HandleJSONDailyStat(dailyVolYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYear, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &threeMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, ALL", HandleJSONDailyStat(dailyFeeYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYear, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &threeMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, ALL", HandleJSONDailyStat(dailyTxYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYear, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &threeMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, ALL", HandleJSONDailyStat(dailyAddrYear))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &sixMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, ALL", HandleJSONDailyStat(dailyVolAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &sixMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, ALL", HandleJSONDailyStat(dailyFeeAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &sixMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, ALL", HandleJSONDailyStat(dailyTxAllTime))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTime, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &sixMonthType, &allPlatformType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, ALL", HandleJSONDailyStat(dailyAddrAllTime))
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
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, BRIDGE", HandleJSONDailyStat(dailyVolMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, BRIDGE", HandleJSONDailyStat(dailyFeeMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, BRIDGE", HandleJSONDailyStat(dailyTxMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, BRIDGE", HandleJSONDailyStat(dailyAddrMonthBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &threeMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, BRIDGE", HandleJSONDailyStat(dailyVolYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &threeMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, BRIDGE", HandleJSONDailyStat(dailyFeeYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &threeMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, BRIDGE", HandleJSONDailyStat(dailyTxYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &threeMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, BRIDGE", HandleJSONDailyStat(dailyAddrYearBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &sixMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, BRIDGE", HandleJSONDailyStat(dailyVolAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &sixMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, BRIDGE", HandleJSONDailyStat(dailyFeeAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &sixMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, BRIDGE", HandleJSONDailyStat(dailyTxAllTimeBridge))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeBridge, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &sixMonthType, &bridgeType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, BRIDGE", HandleJSONDailyStat(dailyAddrAllTimeBridge))
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
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_MONTH, SWAP", HandleJSONDailyStat(dailyVolMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, SWAP", HandleJSONDailyStat(dailyFeeMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, SWAP", HandleJSONDailyStat(dailyTxMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}

		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, SWAP", HandleJSONDailyStat(dailyAddrMonthSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &threeMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_3_MONTHS, SWAP", HandleJSONDailyStat(dailyVolYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &threeMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, SWAP", HandleJSONDailyStat(dailyFeeYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &threeMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, SWAP", HandleJSONDailyStat(dailyTxYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &threeMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, SWAP", HandleJSONDailyStat(dailyAddrYearSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyVolAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &volumeType, &sixMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , VOLUME, PAST_6_MONTHS, SWAP", HandleJSONDailyStat(dailyVolAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyFeeAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &sixMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, SWAP", HandleJSONDailyStat(dailyFeeAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &sixMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, SWAP", HandleJSONDailyStat(dailyTxAllTimeSwap))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeSwap, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &sixMonthType, &swapType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, SWAP", HandleJSONDailyStat(dailyAddrAllTimeSwap))
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
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_MONTH, MESSAGE_BUS", HandleJSONDailyStat(dailyFeeMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &monthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_MONTH, MESSAGE_BUS", HandleJSONDailyStat(dailyTxMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrMonthMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &monthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_MONTH, MESSAGE_BUS", HandleJSONDailyStat(dailyAddrMonthMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyFeeYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &threeMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_3_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyFeeYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &threeMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_3_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyTxYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrYearMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &threeMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_3_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyAddrYearMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dailyFeeAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &feeType, &sixMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , FEE, PAST_6_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyFeeAllTimeMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyTxAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &txType, &sixMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , TRANSACTIONS, PAST_6_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyTxAllTimeMessageBus))
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		dailyAddrAllTimeMessageBus, err := client.GetDailyStatisticsByChain(ctx, nil, &addrType, &sixMonthType, &messagingType, &useMv)
		if err != nil {
			return fmt.Errorf("error rehydrating cache: %w", err)
		}
		err = service.CacheResponse("dailyStatisticsByChain, , ADDRESSES, PAST_6_MONTHS, MESSAGE_BUS", HandleJSONDailyStat(dailyAddrAllTimeMessageBus))
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

// HandleJSONAmountStat converts the gqlClient.GetAmountStatistic to model.ValueResul.
func HandleJSONAmountStat(r *gqlClient.GetAmountStatistic) *model.ValueResult {
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

// HandleJSONDailyStat converts the gqlClient.GetDailyStatisticsByChain to the []*model.DateResultByChain type.
func HandleJSONDailyStat(r *gqlClient.GetDailyStatisticsByChain) []*model.DateResultByChain {
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
