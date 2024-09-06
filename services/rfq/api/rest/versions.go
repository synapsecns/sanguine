package rest

// APIVersionHistory lists historical ApiVersion structs in descending chronological order (current version index 0), each containing further detail about the version.
type APIVersionHistory struct {
	Versions []APIVersion `json:"versions"`
}

// APIVersion specifies the date of a particular API version along with any comments & alerts for integrators to review.
type APIVersion struct {
	Version  string `json:"version"`
	Date     string `json:"date"`
	Comments string `json:"comments"`
	Alerts   string `json:"alerts"`
}

// APIversions contains version information for the Synapse RFQ API. Deprecation notices & other alerts, if any, will be listed here.
// Note: Items must be listed in descending chronological order (current version index 0).
var APIversions = APIVersionHistory{
	Versions: []APIVersion{
		{
			Version:  "1.0",
			Date:     "2024-01-01",
			Comments: "",
			Alerts:   "",
		},
	},
}
