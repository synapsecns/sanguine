package tenderly

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/synapsecns/sanguine/ethergo/chain"
	contracts "github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/debug"
	"github.com/tenderly/tenderly-cli/providers"
	"github.com/tenderly/tenderly-cli/rest/payloads"
)

// VerifyContract verifies a contract on tenderly.
//
//nolint:staticcheck
func (t *Tenderly) VerifyContract(ctx context.Context, chn chain.Chain, contractType contracts.ContractType, contract contracts.DeployedContract) error {
	resultingContract, err := t.AddContract(ctx, chn, contractType, contract)
	if err != nil {
		return fmt.Errorf("error adding contract: %w", err)
	}

	var metadata debug.ContractMetadata
	err = json.Unmarshal([]byte(contractType.ContractInfo().Info.Metadata), &metadata)
	if err != nil {
		return fmt.Errorf("couild not parse metadata: %w", err)
	}

	res, err := t.rest.Contract.UploadContracts(payloads.UploadContractsRequest{
		Contracts: []providers.Contract{*resultingContract},
		Config: &payloads.Config{
			OptimizationsUsed:  &metadata.Settings.Optimizer.Enabled,
			OptimizationsCount: &metadata.Settings.Optimizer.Runs,
			EvmVersion:         &metadata.Settings.EvmVersion,
		},
		Tag: "uploads",
	}, t.projectSlug)

	if err != nil || (res != nil && res.Error != nil) {
		if res != nil && res.Error != nil {
			err = multierror.Append(err, res.Error)
		}
		return fmt.Errorf("could not upload contracts: %w", err)
	}

	_ = res

	return nil
}
