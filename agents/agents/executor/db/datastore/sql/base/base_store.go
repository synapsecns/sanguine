package base

import (
	"fmt"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
	"reflect"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
	// modelNames is a cache that stores the table names for each model. We do this as a cache to avoid
	// having to catch errors when retrieving from dbcommon.GetModelName. One pleasant side effect is that
	// we don't have to use gorm.Parse() every time leading to slight speed up. The map goes from reflect.TypeOf(modelName)->dbcommon.GetModelName
	modelNames map[string]string
}

// NewStore creates a new store.
func NewStore(db *gorm.DB) (*Store, error) {
	modelNames, err := makeModelNames(db)
	if err != nil {
		return nil, fmt.Errorf("failed to make model names: %w", err)
	}
	return &Store{
		db:         db,
		modelNames: modelNames,
	}, nil
}

// makeModelNames creates a list of modelNames from reflect.TypeOf(modelName)->dbcommon.GetModelName.
func makeModelNames(db *gorm.DB) (modelNames map[string]string, err error) {
	modelNames = make(map[string]string)
	// iterate through each model making a map from reflected type to common name
	for _, model := range GetAllModels() {
		ifaceRefection := reflect.TypeOf(model).String()
		modelNames[ifaceRefection], err = dbcommon.GetModelName(db, model)
		if err != nil {
			return nil, fmt.Errorf("failed to get model name for %s: %w", ifaceRefection, err)
		}
	}
	return modelNames, nil
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

// GetModelName gets the precomputed model name of a model. This is guaranteed to work on any model in GetAllModels()
// and will fail w/ a log in the case that the model is not in GetAllModels() an empty name. Please note, like the rest of gorm
// you must pass a pointer to the model in.
func (s Store) GetModelName(model interface{}) (name string) {
	// note: since this is in base and not one of the sql drivers themselves, we can export this here
	// just as long as it's not used through the service interface
	name, ok := s.modelNames[reflect.TypeOf(model).String()]
	if !ok {
		warning := fmt.Sprintf("model %T not found", model)
		// one common mistake here isi passing in an interface so we're going to specifically warn the caller if there's an interface
		if reflect.TypeOf(model).Kind() != reflect.Pointer {
			warning += ". Please note, you must pass a pointer to the model in."
		}

		// note: there's a case to be made for stronger error handling here, but given gorm isn't explicit about this kind of thing either
		// see: db.Model(), we're not going to bother
		logger.Errorf(warning)
	}
	return name
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&Message{}, &Attestation{}, &State{},
	)
	return allModels
}

var _ db.ExecutorDB = Store{}
