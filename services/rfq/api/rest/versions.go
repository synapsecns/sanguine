package rest

type ApiVersionHistory struct {
	Versions []ApiVersion `json:"versions"`
}

type ApiVersion struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Comment string `json:"_comment"`
	Alerts  string `json:"_alerts"`
}

// ApiVersionHistory contains version information for the Synapse RFQ API. Deprecation notices & other alerts, if any, will be listed here.
// Note: Items must be listed in descending chronological order (current version index 0).
var apiVersions = ApiVersionHistory{
	Versions: []ApiVersion{
		{
			Version: "1.0",
			Date:    "2024-01-01",
			Comment: "",
			Alerts:  "",
		},
	},
}
