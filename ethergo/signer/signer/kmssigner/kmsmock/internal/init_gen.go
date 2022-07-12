// Code copied from github.com/nsmithuk/local-kms/src:/init.go for testing by synapse modulecopier DO NOT EDIT."

package internal

import (
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {

	//logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

}
