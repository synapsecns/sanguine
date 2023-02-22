package server

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

// SetupGRPCServer sets up the grpc server.
func SetupGRPCServer(ctx context.Context, engine *gin.Engine, eventDB db.EventDB) (*grpc.Server, error) {
	s := grpc.NewServer()
	sImpl := server{
		db: eventDB,
	}

	mux := runtime.NewServeMux()
	pbscribe.RegisterScribeServiceServer(s, &sImpl)
	err := pbscribe.RegisterScribeServiceHandlerServer(ctx, mux, &sImpl)
	if err != nil {
		return nil, fmt.Errorf("could not register server")
	}

	engine.NoRoute(func(c *gin.Context) {
		c.Status(200)
		gin.WrapF(mux.ServeHTTP)(c)
	})

	return s, nil
}

type server struct {
	// db is the db to use for the server
	db db.EventDB
	pbscribe.UnimplementedScribeServiceServer
}

var logger = log.Logger("scribe-grpc")

func (s *server) FilterLogs(ctx context.Context, req *pbscribe.FilterLogsRequest) (*pbscribe.FilterLogsResponse, error) {
	logFilter := req.Filter.ToNative()
	logFilter.ChainID = req.Filter.ChainId

	logs, err := s.db.RetrieveLogsWithFilter(ctx, logFilter, int(req.Page))
	if err != nil {
		logger.Errorf("error retreiving logs: %v", err)
		return nil, fmt.Errorf("error retreiving logs: %w", err)
	}

	return &pbscribe.FilterLogsResponse{
		Logs: pbscribe.FromNativeLogs(logs),
	}, nil
}

//nolint:gocognit,cyclop
func (s *server) StreamLogs(req *pbscribe.StreamLogsRequest, res pbscribe.ScribeService_StreamLogsServer) error {
	streamNewBlocks := false
	fromBlock, toBlock, err := s.setBlocks(res.Context(), req)
	if err != nil {
		logger.Errorf("could not set blocks: %v", err)
		return fmt.Errorf("could not set blocks: %w", err)
	}

	wait := 0
	nextFromBlock := uint64(0)
	logFilter := req.Filter.ToNative()
	logFilter.ChainID = req.Filter.ChainId

	if req.ToBlock == "latest" {
		streamNewBlocks = true
	}

	for {
		var retrievedLogs []*types.Log

		if nextFromBlock > fromBlock {
			fromBlock = nextFromBlock
		}

		page := 1

		for {
			logs, err := s.db.RetrieveLogsInRangeAsc(res.Context(), logFilter, fromBlock, toBlock, page)
			if err != nil {
				logger.Errorf("could not retrieve logs: %v", err)
				return fmt.Errorf("could not retrieve logs: %w", err)
			}

			retrievedLogs = append(retrievedLogs, logs...)

			// See if we do not need to get the next page.
			if len(logs) < base.PageSize {
				break
			}

			page++
		}

		// Convert the logs to the protobuf format and send them through the stream.
		for _, log := range retrievedLogs {
			err = res.Send(&pbscribe.StreamLogsResponse{
				Log: pbscribe.FromNativeLog(log),
			})
			if err != nil {
				logger.Errorf("could not send log: %v", err)
				return fmt.Errorf("could not send log: %w", err)
			}
		}

		if !streamNewBlocks {
			return nil
		}

	STREAM:
		for {
			select {
			case <-res.Context().Done():
				return nil
			default:
				time.Sleep(time.Duration(wait) * time.Second)
				latestScribeBlock, err := s.db.RetrieveLastIndexed(res.Context(), common.HexToAddress(req.Filter.ContractAddress.GetData()), req.Filter.ChainId)
				if err != nil {
					logger.Errorf("could not retrieve last indexed block: %v", err)
					return fmt.Errorf("could not retrieve last indexed block: %w", err)
				}

				if latestScribeBlock > toBlock {
					nextFromBlock = toBlock + 1
					toBlock = latestScribeBlock
					wait = 0
					break STREAM
				}
				wait = 1
			}
		}
	}
}

func (s *server) Check(context.Context, *pbscribe.HealthCheckRequest) (*pbscribe.HealthCheckResponse, error) {
	return &pbscribe.HealthCheckResponse{Status: pbscribe.HealthCheckResponse_SERVING}, nil
}

func (s *server) Watch(a *pbscribe.HealthCheckRequest, res pbscribe.ScribeService_WatchServer) error {
	for {
		select {
		case <-res.Context().Done():
			err := res.Context().Err()
			if err != nil {
				return fmt.Errorf("context finished: %w", err)
			}
			return nil
		case <-time.After(time.Second):
			err := res.Send(&pbscribe.HealthCheckResponse{Status: pbscribe.HealthCheckResponse_SERVING})
			if err != nil {
				return fmt.Errorf("could not check response: %w", err)
			}
		}
	}
}

func (s *server) setBlocks(ctx context.Context, req *pbscribe.StreamLogsRequest) (uint64, uint64, error) {
	blocks := []string{req.FromBlock, req.ToBlock}
	resBlocks := make([]uint64, 2)

	for i, block := range blocks {
		switch block {
		case "latest":
			lastIndexed, err := s.db.RetrieveLastIndexed(ctx, common.HexToAddress(req.Filter.ContractAddress.GetData()), req.Filter.ChainId)
			if err != nil {
				logger.Errorf("could not retrieve last indexed block: %v", err)
				return 0, 0, fmt.Errorf("could not retrieve last indexed block: %w", err)
			}

			resBlocks[i] = lastIndexed
		case "earliest":
			resBlocks[i] = 0
		default:
			blockNum, err := strconv.ParseUint(block, 16, 64)
			if err != nil {
				logger.Errorf("could not parse %s block: %v", block, err)
				return 0, 0, fmt.Errorf("could not parse %s block: %w", block, err)
			}

			resBlocks[i] = blockNum
		}
	}

	return resBlocks[0], resBlocks[1], nil
}
