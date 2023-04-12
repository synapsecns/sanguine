package testhelper_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"github.com/synapsecns/sanguine/services/scribe/testhelper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/big"
)

func (s *TestHelperSuite) TestEmbeddedScribe() {
	testScribe := testhelper.NewTestScribe(s.GetTestContext(), s.T(), s.deployManager.GetDeployedContracts(), s.testBackends...)

	// let's send some messages on each domain
	g, gctx := errgroup.WithContext(s.GetTestContext())
	for _, backend := range s.testBackends {
		backend := backend // capture func literal
		_, testContract := s.deployManager.GetTestContract(gctx, backend)
		for i := 0; i < 10; i++ {
			g.Go(func() error {
				txContext := backend.GetTxContext(gctx, nil)
				tx, err := testContract.EmitEventAandB(txContext.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
				Nil(s.T(), err)

				backend.WaitForConfirmation(gctx, tx)

				return nil
			})
		}
	}

	Nil(s.T(), g.Wait())

	s.Eventually(func() bool {
		conn, err := grpc.DialContext(s.GetSuiteContext(), testScribe, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return false
		}

		grpcClient := pbscribe.NewScribeServiceClient(conn)

		healthCheck, err := grpcClient.Check(s.GetSuiteContext(), &pbscribe.HealthCheckRequest{})
		if err != nil {
			return false
		}

		if healthCheck.GetStatus() != pbscribe.HealthCheckResponse_SERVING {
			return false
		}

		// iterate through each backend and make sure there's at 20 logs
		for _, backend := range s.testBackends {
			logs, err := grpcClient.FilterLogs(s.GetTestContext(), &pbscribe.FilterLogsRequest{
				Filter: &pbscribe.LogFilter{
					ChainId: uint32(backend.GetChainID()),
				},
			})

			// no reason to error here except a bad request since we made sure the server was serving
			if err != nil {
				s.T().Error(err)
			}

			if len(logs.GetLogs()) < 20 {
				return false
			}

			if err != nil {
				return false
			}
		}
		return true
	})
}
