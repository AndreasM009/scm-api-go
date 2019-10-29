package cosmosrepository

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"scm-api/internal/domainobjects"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/google/uuid"

	"gopkg.in/mgo.v2"
)

// CosmosRepository ...
type CosmosRepository struct {
}

type database struct {
	session *mgo.Session
}

var (
	db   database
	once sync.Once
)

func getDatabase() *database {
	once.Do(func() {
		dialInfo := &mgo.DialInfo{
			Addrs:    []string{fmt.Sprintf("%s.documents.azure.com:10255", "dbname")},
			Timeout:  60 * time.Second,
			Database: "dbname",   // It can be anything
			Username: "username", // Username
			Password: "password", // PASSWORD
			DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
				return tls.Dial("tcp", addr.String(), &tls.Config{})
			},
		}
		session, err := mgo.DialWithInfo(dialInfo)

		if err != nil {
			fmt.Printf("Can not connect, go error %v\n", err)
			os.Exit(0)
		}

		session.SetSafe(&mgo.Safe{})
		db.session = session
	})

	return &db
}

// Add ...
func (r CosmosRepository) Add(c *domainobjects.Customer) domainobjects.Customer {
	database := getDatabase()
	session := database.session
	collection := session.DB("databasename").C("customers")

	c.Id = uuid.New()
	collection.Insert(c)
	return *c
}

// Get ...
func (r CosmosRepository) Get(id uuid.UUID) (domainobjects.Customer, error) {
	database := getDatabase()
	session := database.session
	collection := session.DB("databasename").C("customers")

	c := domainobjects.Customer{}
	err := collection.Find(bson.M{"id": id.String()}).One(&c)

	if err != nil {
		return domainobjects.Customer{}, err
	}

	return c, nil
}
