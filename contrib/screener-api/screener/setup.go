package screener

import (
	"encoding/csv"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener/internal"
	"io"
	"os"
	"strings"
)

func setupScreener(rulesets map[string]config.RulesetConfig) (internal.RulesetManager, error) {
	mgr := internal.NewRulesetManager(map[string]map[string]bool{})
	for csvName, cfg := range rulesets {
		csvPath := cfg.Filename
		parsedCsv, err := parseCsv(csvPath)
		if err != nil {
			return nil, fmt.Errorf("could not parse csv %s: %w", csvName, err)
		}
		err = mgr.AddRuleset(csvName, parsedCsv)
		if err != nil {
			return nil, fmt.Errorf("could not add ruleset %s: %w", csvName, err)
		}
	}

	return mgr, nil
}

func parseCsv(file string) (risks map[string]bool, err error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open blacklist file: %w", err)
	}

	defer func() {
		_ = fileHandle.Close()
	}()

	r := csv.NewReader(fileHandle)
	risks = make(map[string]bool)
	i := 0
	for {
		// skip first row
		if i == 0 {
			i++
			continue
		}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read blacklist file: %w", err)
		}

		shoudBlock := false
		if record[7] != "" {
			shoudBlock = true
		}

		i++
		// assumes record[2] and record[4] are uniform.
		risks[strings.ToLower(fmt.Sprintf("%s_%s", record[2], record[4]))] = shoudBlock
	}

	return risks, nil
}
