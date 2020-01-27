package postgres

import (
	"time"

	"github.com/confus1on/meetmeup/ent"
	_ "github.com/lib/pq"

	"github.com/facebookincubator/ent/dialect/sql"
)

var Client *ent.Client

func Open() (*ent.Client, error) {
	if Client != nil {
		NewMeetupsRepo(Client)
		NewUsersRepo(Client)
		return Client, nil
	}

	connStr := "postgres://mxwcpqiz:8Mkenq89dh4svTcRi6C5vcdlwtWEYPhE@rosie.db.elephantsql.com:5432/mxwcpqiz"

	driver, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db := driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	Client = ent.NewClient(ent.Driver(driver))

	NewMeetupsRepo(Client)
	NewUsersRepo(Client)

	return Client, nil
}
