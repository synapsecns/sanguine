package base_test

import (
	"context"
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/core/metrics"
	"testing"
)

// TestGetModelName makes sure are getmodelname works correctly.
// Testing this on all models could be prone to breaking frequently
// as models change so we test n the first one
// since if the reflect is broken, the name will have gotten overwritten.
func TestGetModelName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	store, err := sqlite.NewSqliteStore(ctx, filet.TmpDir(t, ""), metrics.NewNullHandler(), false)
	Nil(t, err)

	messageName := store.GetModelName(&base.Message{})
	Equal(t, "messages", messageName)
}
