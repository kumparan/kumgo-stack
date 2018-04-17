package connector

import (
	"context"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/config"
)

var ArangoDB driver.Database

func OpenArangoConnection() (db driver.Database) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.ArangoHost()},
	})
	if err != nil {
		logger.Err("Connection to Arango Server failed: ", err)
		return
	}
	logger.Info("Connection to Arango Server success...")

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.ArangoUsername(), config.ArangoPassword()),
	})
	if err != nil {
		logger.Err("Create new arango client failed: ", err)
	}

	db, err = client.Database(context.Background(), config.ArangoDatabase())
	if err != nil {
		logger.Err("Connection to database failed: ", err)
	}
	logger.Info("Connection to database success...")

	return
}
