package rest

// ApiVersionHistroy lists historical ApiVersion structs in descending chronological order (current version index 0), each containing further detail about the version.
type ApiVersionHistory struct {
	Versions []ApiVersion `json:"versions"`
}

// ApiVersion specifies the date of a particular API version along with any comments & alerts for integrators to review.
type ApiVersion struct {
	Version  string `json:"version"`
	Date     string `json:"date"`
	Comments string `json:"comments"`
	Alerts   string `json:"alerts"`
}

// ApiVersionHistory contains version information for the Synapse RFQ API. Deprecation notices & other alerts, if any, will be listed here.
// Note: Items must be listed in descending chronological order (current version index 0).
var apiVersions = ApiVersionHistory{
	Versions: []ApiVersion{
		{
			Version:  "1.0",
			Date:     "2024-01-01",
			Comments: "",
			Alerts:   "",
		},
	},
}
