package http

import (
	"database/sql"

	"github.com/saeedmdd/go-hexa/pkg/config"
	"github.com/saeedmdd/go-hexa/utils"
)

var databaseConnection = func () (*sql.DB, error) {
	connectionConfig := config.AppConfig.Database.Postgres
	return utils.PostgresConnection(
		connectionConfig.Host,
		connectionConfig.UserName,
		connectionConfig.Password,
		connectionConfig.Port,
		connectionConfig.Database,
		connectionConfig.Sslmode,
	)
}