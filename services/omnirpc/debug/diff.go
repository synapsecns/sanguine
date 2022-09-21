package debug

import (
	"encoding/json"
	"fmt"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"runtime/debug"
)

// HashDiff generates a diff
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

	uniqueResponses := make([]json.RawMessage, 2)
	i := 0

	// get first two responses
OUTER:
	for _, resps := range errResp.Hashes {
		uniqueResponses[i] = resps[0].Raw
		i++

		if i >= 2 {
			break OUTER
		}
	}

	debug.SetGCPercent(-1)

	differ := diff.New()
	respDiff, err := differ.Compare(uniqueResponses[0], uniqueResponses[1])
	if err != nil {
		return fmt.Errorf("could not compare: %w", err)
	}

	config := formatter.AsciiFormatterConfig{
		ShowArrayIndex: true,
		Coloring:       true,
	}

	var aJson map[string]interface{}
	err = json.Unmarshal(uniqueResponses[0], &aJson)

	r := formatter.NewAsciiFormatter(aJson, config)
	diffString, err := r.Format(respDiff)
	if err != nil {
		panic(err)
	}

	fmt.Println(diffString)

	return nil
}
