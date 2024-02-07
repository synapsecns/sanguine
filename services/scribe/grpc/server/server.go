package server

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

// SetupGRPCServer sets up the grpc server.
func SetupGRPCServer(ctx context.Context, engine *gin.Engine, eventDB db.EventDB, handler metrics.Handler) (*grpc.Server, error) {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	sImpl := server{
		db:      eventDB,
		handler: handler,
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
	handler metrics.Handler
}

func (s *server) FilterLogs(ctx context.Context, req *pbscribe.FilterLogsRequest) (*pbscribe.FilterLogsResponse, error) {
	logFilter := req.Filter.ToNative()
	logFilter.ChainID = req.Filter.ChainId

	logs, err := s.db.RetrieveLogsWithFilter(ctx, logFilter, int(req.Page))
	if err != nil {
		return nil, fmt.Errorf("error retreiving logs: %w", err)
	}

	return &pbscribe.FilterLogsResponse{
		Logs: pbscribe.FromNativeLogs(logs),
	}, nil
}

//nolint:gocognit,cyclop
func (s *server) StreamLogs(req *pbscribe.StreamLogsRequest, res pbscribe.ScribeService_StreamLogsServer) error {
	streamNewBlocks := false
	retrieveLogsBackoff := 3
	fromBlock, toBlock, err := s.setBlocks(res.Context(), req)
	if err != nil {
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

		ctx, span := s.handler.Tracer().Start(res.Context(), "grpc.StreamLogsLoop", trace.WithAttributes(
			attribute.Int(metrics.ChainID, int(req.Filter.ChainId)),
			attribute.String(metrics.ContractAddress, req.Filter.ContractAddress.GetData()),
			attribute.Int("fromBlock", int(fromBlock)),
			attribute.Int("toBlock", int(toBlock)),
		))

		page := 1

		for {
			logs, err := s.db.RetrieveLogsInRangeAsc(ctx, logFilter, fromBlock, toBlock, page)
			if err != nil {
				time.Sleep(time.Duration(retrieveLogsBackoff) * time.Second)
				continue
			}

			retrievedLogs = append(retrievedLogs, logs...)

			// See if we do not need to get the next page.
			if len(logs) < base.PageSize {
				break
			}

			span.AddEvent("Getting next page. Page: " + strconv.Itoa(page))

			page++
		}

		// Convert the logs to the protobuf format and send them through the stream.
		for _, log := range retrievedLogs {
			err = res.Send(&pbscribe.StreamLogsResponse{
				Log: pbscribe.FromNativeLog(log),
			})
			if err != nil {
				return fmt.Errorf("could not send log: %w", err)
			}

			span.AddEvent("Sending log.", trace.WithAttributes(
				attribute.String(metrics.TxHash, log.TxHash.String()),
			))
		}

		span.AddEvent("Got logs. Count: " + strconv.Itoa(len(retrievedLogs)))

		if !streamNewBlocks {
			go func() {
				span.End()
			}()
			return nil
		}

	STREAM:
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				// TODO: Make wait time configurable (?).
				time.Sleep(time.Duration(wait) * time.Second)
				wait = 1
				latestScribeBlock, err := s.db.RetrieveLastIndexed(ctx, common.HexToAddress(req.Filter.ContractAddress.GetData()), req.Filter.ChainId, false)
				if err != nil {
					continue
				}

				if latestScribeBlock > toBlock {
					nextFromBlock = toBlock + 1
					toBlock = latestScribeBlock
					wait = 0

					span.AddEvent("New block. From: " + strconv.Itoa(int(nextFromBlock)) + " To: " + strconv.Itoa(int(toBlock)))

					go func() {
						span.End()
					}()

					break STREAM
				}
			}
		}
	}
}

func (s *server) Check(context.Context, *pbscribe.HealthCheckRequest) (*pbscribe.HealthCheckResponse, error) {
	return &pbscribe.HealthCheckResponse{Status: pbscribe.HealthCheckResponse_SERVING}, nil
}

func (s *server) Watch(_ *pbscribe.HealthCheckRequest, res pbscribe.ScribeService_WatchServer) error {
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
			lastIndexed, err := s.db.RetrieveLastIndexed(ctx, common.HexToAddress(req.Filter.ContractAddress.GetData()), req.Filter.ChainId, false)
			if err != nil {
				return 0, 0, fmt.Errorf("could not retrieve last indexed block: %w", err)
			}

			resBlocks[i] = lastIndexed
		case "earliest":
			resBlocks[i] = 0
		default:
			blockNum, err := strconv.ParseUint(block, 16, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("could not parse %s block: %w", block, err)
			}

			resBlocks[i] = blockNum
		}
	}

	return resBlocks[0], resBlocks[1], nil
}
