// Package updater attests to updates
package updater

import (
	"context"
	"github.com/synapsecns/sanguine/core/config"
)

// Updater updates the updater.
type Updater struct {
}

// NewUpdater creates a new updater.
func NewUpdater(cfg config.Config) (Updater, error) {
	return Updater{}, nil
}

// Start starts the updater.
func (u Updater) Start(ctx context.Context) {

}
