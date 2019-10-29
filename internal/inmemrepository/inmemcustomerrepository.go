package inmemrepository

import (
	"scm-api/internal/domainobjects"
	"sync"

	"github.com/google/uuid"
)

type dbcontainer map[uuid.UUID]*domainobjects.Customer

type database struct {
	container dbcontainer
	syncRoot  sync.Mutex
}

var (
	db   *database
	once sync.Once
)

func getDatabase() *database {
	once.Do(func() {
		db = &database{}
		db.container = make(dbcontainer)

		id := uuid.New()
		c := &domainobjects.Customer{
			Id:        id,
			Firstname: "Andreas",
			Lastname:  "Mock",
		}

		db.container[c.Id] = c
	})

	return db
}

// InMemRepository ...
type InMemRepository struct {
}

// CreateRepository ...
func CreateRepository() *InMemRepository {
	return &InMemRepository{}
}

// Add ...
func (r InMemRepository) Add(c *domainobjects.Customer) domainobjects.Customer {
	var db = getDatabase()
	db.syncRoot.Lock()
	defer db.syncRoot.Unlock()
	db.container[c.Id] = c
	return *c
}

// Get ...
func (r InMemRepository) Get(id uuid.UUID) domainobjects.Customer {
	var db = getDatabase()
	db.syncRoot.Lock()
	defer db.syncRoot.Unlock()
	return *db.container[id]
}

// GetAll ...
func (r InMemRepository) GetAll() []domainobjects.Customer {
	db := getDatabase()

	db.syncRoot.Lock()
	defer db.syncRoot.Unlock()

	result := make([]domainobjects.Customer, 0, len(db.container))

	for _, value := range db.container {
		result = append(result, *value)
	}

	return result
}

// DeleteAll ...
func (r InMemRepository) DeleteAll() []domainobjects.Customer {
	db := getDatabase()

	db.syncRoot.Lock()
	defer db.syncRoot.Unlock()

	result := make([]domainobjects.Customer, 0, len(db.container))

	for _, value := range db.container {
		result = append(result, *value)
	}

	db.container = make(dbcontainer)

	return result
}
