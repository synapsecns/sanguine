package screener

import (
	"encoding/csv"
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
	ID         string `csv:"ID"`
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
		risks[internal.MakeParam(screener.Category, screener.TypeOfRisk)] = strings.EqualFold(screener.Enabled, "true") || strings.EqualFold(screener.Enabled, "yes")
	}

	return risks, nil
}

func splitCSV(file string) (map[string][]Set, error) {
	//nolint: gosec
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("could not open blacklist file: %w", err)
	}

	defer func() {
		_ = fileHandle.Close()
	}()

	reader := csv.NewReader(fileHandle)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read csv file: %w", err)
	}

	header := records[0]

	// "ExtraColumn->Row w/ rewritten enabled based on column name"
	outFiles := make(map[string][]Set)

	for _, record := range records[1:] {
		screener := Set{
			Enabled:    record[0],
			ID:         record[1],
			Category:   record[2],
			Name:       record[3],
			TypeOfRisk: record[4],
			Severity:   record[5],
		}

		for i := 6; i < len(header); i++ {
			// skip empty records
			if record[i] == "" {
				continue
			}
			// reset the enabled field to the current record
			newRecord := screener
			newRecord.Enabled = record[i]

			// add the record to the set of rules
			outFiles[header[i]] = append(outFiles[header[i]], newRecord)
		}
	}

	return outFiles, nil
}

// SplitAndWriteCSV splits a csv file and writes the output to a directory.
func SplitAndWriteCSV(file string, outDir string) (files []string, err error) {
	outFiles, err := splitCSV(file)
	if err != nil {
		return nil, fmt.Errorf("could not split csv: %w", err)
	}

	for k, v := range outFiles {
		outFile, err := os.Create(fmt.Sprintf("%s/%s.csv", outDir, k))
		if err != nil {
			return nil, fmt.Errorf("could not create csv: %w", err)
		}
		defer func() {
			_ = outFile.Close()
		}()
		v := v
		err = gocsv.MarshalFile(&v, outFile)
		if err != nil {
			return nil, fmt.Errorf("could not marshal csv: %w", err)
		}
	}

	return files, nil
}
