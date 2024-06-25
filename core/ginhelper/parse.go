package ginhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// StringMaxLength is the maximum length of a string to display.
var StringMaxLength = 0

// Newline is the string to use for newlines.
var Newline = "\n"

// Indent is the number of spaces to indent.
var Indent = 4

// BeautifyJSONBytes beautifies json bytes.
func BeautifyJSONBytes(data []byte, hiddenFields []string) ([]byte, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	v = removeHiddenFields(v, hiddenFields)

	return []byte(format(v, 1)), nil
}

// FormatToBeautifulJSON dumps v to beautified json bytes.
func FormatToBeautifulJSON(v interface{}, hiddenFields []string) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}
	return BeautifyJSONBytes(data, hiddenFields)
}

func format(v interface{}, depth int) string {
	switch val := v.(type) {
	case string:
		return formatString(val)
	case float64:
		return fmt.Sprint(strconv.FormatFloat(val, 'f', -1, 64))
	case bool:
		return fmt.Sprint(strconv.FormatBool(val))
	case nil:
		return "null"
	case map[string]interface{}:
		return formatMap(val, depth)
	case []interface{}:
		return formatArray(val, depth)
	}

	return ""
}

func formatString(s string) string {
	r := []rune(s)
	if StringMaxLength != 0 && len(r) >= StringMaxLength {
		s = string(r[0:StringMaxLength]) + "..."
	}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	// TODO: check error
	//nolint:errchkjson
	_ = encoder.Encode(s)
	s = buf.String()
	s = strings.TrimSuffix(s, "\n")

	return fmt.Sprint(s)
}

func formatMap(m map[string]interface{}, depth int) string {
	if len(m) == 0 {
		return "{}"
	}

	currentIndent := generateIndent(depth - 1)
	nextIndent := generateIndent(depth)
	rows := []string{}
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		val := m[key]
		k := fmt.Sprintf(`"%s"`, key)
		v := format(val, depth+1)

		valueIndent := " "
		if Newline == "" {
			valueIndent = ""
		}
		row := fmt.Sprintf("%s%s:%s%s", nextIndent, k, valueIndent, v)
		rows = append(rows, row)
	}

	return fmt.Sprintf("{%s%s%s%s}", Newline, strings.Join(rows, ","+Newline), Newline, currentIndent)
}

func formatArray(a []interface{}, depth int) string {
	if len(a) == 0 {
		return "[]"
	}

	currentIndent := generateIndent(depth - 1)
	nextIndent := generateIndent(depth)
	rows := []string{}

	for _, val := range a {
		c := format(val, depth+1)
		row := nextIndent + c
		rows = append(rows, row)
	}
	return fmt.Sprintf("[%s%s%s%s]", Newline, strings.Join(rows, ","+Newline), Newline, currentIndent)
}

func generateIndent(depth int) string {
	return strings.Repeat(" ", Indent*depth)
}

func removeHiddenFields(v interface{}, hiddenFields []string) interface{} {
	if _, ok := v.(map[string]interface{}); !ok {
		return v
	}

	// nolint: forcetypeassert
	m := v.(map[string]interface{})

	// case insensitive key deletion
	for _, hiddenField := range hiddenFields {
		for k := range m {
			if strings.EqualFold(k, hiddenField) {
				delete(m, k)
			}
		}
	}

	return m
}
