package server

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"google.golang.org/grpc"
	"math"
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

func (s *server) StreamLogs(req *pbscribe.StreamLogsRequest, res pbscribe.ScribeService_StreamLogsServer) error {
	var retrievedLogs []*types.Log

	logFilter := req.Filter.ToNative()
	logFilter.ChainID = req.Filter.ChainId
	fromBlock, err := setBlock(req.FromBlock)
	if err != nil {
		return fmt.Errorf("could not set from block: %w", err)
	}

	toBlock, err := setBlock(req.ToBlock)
	if err != nil {
		return fmt.Errorf("could not set to block: %w", err)
	}

	page := 1

	for {
		logs, err := s.db.RetrieveLogsInRange(res.Context(), logFilter, fromBlock, toBlock, page)
		if err != nil {
			return fmt.Errorf("could not retrieve logs: %w", err)
		}

		retrievedLogs = append(sliceReverse(logs), retrievedLogs...)

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
			return fmt.Errorf("could not send log: %w", err)
		}
	}

	return nil
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

func setBlock(s string) (uint64, error) {
	switch s {
	case "latest":
		return math.MaxUint64, nil
	case "earliest":
		return 0, nil
	case "pending":
		return math.MaxUint64, nil
	default:
		block, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse block: %w", err)
		}

		return uint64(block), nil
	}
}

func sliceReverse(a []*types.Log) []*types.Log {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
