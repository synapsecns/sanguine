package tunnel

// here we copy the iap module from github.com/gartnera/gcloud/compute/iap

import _ "github.com/gartnera/gcloud/compute/iap"

//go:generate go run github.com/synapsecns/sanguine/tools/modulecopier --module-path github.com/gartnera/gcloud/compute/iap --package-name tunnel
