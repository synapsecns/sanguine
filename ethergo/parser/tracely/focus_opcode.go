package tracely

var OpCodeFocus map[string]bool = map[string]bool{
	"CALL":         true,
	"DELEGATECALL": true,
	"STATICCALL":   true,
	"RETURN":       true,
	"CREATE":       true,
	"REVERT":       true,
}
