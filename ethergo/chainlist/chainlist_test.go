package chainlist_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chainlist"
	"golang.org/x/sync/errgroup"
	"gopkg.in/resty.v1"
	"net/http"
	"os"
)

func (s *ChainlistSuite) TestChainList() {
	if os.Getenv("CI") == "" {
		mainnetChain := chainlist.ChainIDs.ChainByID(uint(params.MainnetChainConfig.ChainID.Uint64()))
		NotNil(s.T(), mainnetChain)

		client, err := ethclient.DialContext(s.GetTestContext(), "https://cloudflare-eth.com")
		Nil(s.T(), err)

		g, _ := errgroup.WithContext(s.GetTestContext())
		restClient := resty.New()

		// test block url
		g.Go(func() error {
			block, err := client.BlockByHash(s.GetTestContext(), params.MainnetGenesisHash)
			Nil(s.T(), err)

			url := mainnetChain.ExplorerURL(block)
			pointerURL := mainnetChain.ExplorerURL(*block)
			Equal(s.T(), pointerURL, url)

			resp, err := restClient.R().SetContext(s.GetTestContext()).Get(url)
			Nil(s.T(), err)

			True(s.T(), resp.StatusCode() != http.StatusNotFound)
			return errors.Wrap(err, "could not get block")
		})

		// tx test
		g.Go(func() error {
			// https://www.reddit.com/r/ethereum/comments/6qildp/what_is_the_first_ever_ethereum_transaction/dkxtddf?utm_source=share&utm_medium=web2x&context=3
			tx, err := client.TransactionInBlock(s.GetTestContext(), common.HexToHash("0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd"), 0)
			Nil(s.T(), err)

			url := mainnetChain.ExplorerURL(tx)
			pointerURL := mainnetChain.ExplorerURL(*tx)
			txURL := mainnetChain.TxHashURL(tx.Hash())
			Equal(s.T(), pointerURL, url, txURL)

			resp, err := restClient.R().SetContext(s.GetTestContext()).Get(url)
			Nil(s.T(), err)
			True(s.T(), resp.StatusCode() != http.StatusNotFound)
			return errors.Wrap(err, "could not get tx")
		})

		// address test
		g.Go(func() error {
			// https://www.reddit.com/r/ethereum/comments/6qildp/what_is_the_first_ever_ethereum_transaction/dkxtddf?utm_source=share&utm_medium=web2x&context=3
			address := params.MainnetCheckpointOracle.Address

			url := mainnetChain.ExplorerURL(address)
			pointerURL := mainnetChain.ExplorerURL(&address)
			Equal(s.T(), pointerURL, url)

			resp, err := restClient.R().SetContext(s.GetTestContext()).Get(url)
			Nil(s.T(), err)

			True(s.T(), resp.StatusCode() != http.StatusNotFound)
			return errors.Wrap(err, "could not get address")
		})

		Nil(s.T(), g.Wait())
	}
}
