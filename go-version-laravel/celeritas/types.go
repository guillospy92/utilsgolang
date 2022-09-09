package celeritas

import "database/sql"

// initPath path init
type initPath struct {
	rootPath    string
	folderNames []string
}

// cookieConfig cookie config params system
type cookieConfig struct {
	name     string
	lifeTime string
	persists string
	secure   string
	domain   string
}

type dataBaseConfig struct {
	dns        string
	dbType     string
	dbHost     string
	dbName     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbSslMode  string
}

type Database struct {
	DataType string
	Pool     *sql.DB
}
