package debug

import (
	"encoding/json"
	"fmt"
	jd "github.com/josephburnett/jd/lib"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
)

// HashDiff generates a diff.
func HashDiff(fileContents []byte) error {
	var errResp proxy.ErrorResponse

	err := json.Unmarshal(fileContents, &errResp)
	if err != nil {
		return fmt.Errorf("could not unmarshall json: %w", err)
	}

	if len(errResp.Hashes) < 2 {
		return fmt.Errorf("found %d hash(es), cannot compare: %w", len(errResp.Hashes), err)
	}

	if len(errResp.Hashes) > 2 {
		fmt.Printf("found %d hashes, only comparing the first 2", len(errResp.Hashes))
	}

	uniqueResponses := make([]jd.JsonNode, 2)
	i := 0
	// get first two responses
OUTER:
	for _, resps := range errResp.Hashes {
		uniqueResponses[i], err = jd.ReadJsonString(string(resps[0].Raw))
		if err != nil {
			return fmt.Errorf("could not make json string: %w", err)
		}
		i++

		if i >= 2 {
			break OUTER
		}
	}

	diff := uniqueResponses[0].Diff(uniqueResponses[1])

	fmt.Println(diff.Render(jd.COLOR))
	return nil
}
