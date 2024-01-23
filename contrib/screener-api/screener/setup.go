package screener

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/synapsecns/sanguine/contrib/screener-api/config"
	"github.com/synapsecns/sanguine/contrib/screener-api/screener/internal"
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

// Set is a struct for the screener set.
type Set struct {
	Enabled    string `csv:"Enabled"`
	ID         int    `csv:"ID"`
	Category   string `csv:"Category"`
	Name       string `csv:"Name"`
	TypeOfRisk string `csv:"Type of risk"`
	Severity   string `csv:"Severity"`
}

func parseCsv(file string) (risks map[string]bool, err error) {
	//nolint: gosec
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open blacklist file: %w", err)
	}

	defer func() {
		_ = fileHandle.Close()
	}()

	var screeners []Set
	if err := gocsv.UnmarshalFile(fileHandle, &screeners); err != nil { // Load clients from file
		return nil, fmt.Errorf("could not unmarshal blacklist file: %w", err)
	}

	risks = make(map[string]bool)
	for _, screener := range screeners {
		risks[strings.ToLower(fmt.Sprintf("%s_%s", screener.Category, screener.TypeOfRisk))] = screener.Enabled == "true"
	}

	return risks, nil
}
