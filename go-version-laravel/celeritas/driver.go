package celeritas

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
)

// BuildDns interpreter database type
func (a *Accelerator) BuildDns() string {
	var dns string
	switch a.Config.database.dbType {
	case "postgres", "postgresql":
		dns = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			a.Config.database.dbHost,
			a.Config.database.dbPort,
			a.Config.database.dbUser,
			a.Config.database.dbName,
			a.Config.database.dbSslMode,
		)

		if a.Config.database.dbPassword != "" {
			dns = fmt.Sprintf("%s password=%s", dns, a.Config.database.dbPassword)
		}

	default:

	}

	return dns
}

func (a *Accelerator) openDB(dns string) (*sql.DB, error) {
	dbType := a.Config.database.dbType
	if dbType == "postgres" || dbType == "postgresql" {
		dbType = "pgx"
	}

	db, err := sql.Open(dbType, dns)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
