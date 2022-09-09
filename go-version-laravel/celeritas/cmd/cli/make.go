package main

import (
	"errors"
	"fmt"
	"time"
)

var errorMigration = errors.New("you must give the migration a name")

func doMake(args2, args3 string) error {
	switch args2 {
	case "migration":
		dbType := cel.DB.DataType
		if args3 == "" {
			exitGraceFully(errorMigration)
		}

		// create folders migrations
		err := cel.CreateDirNotExists(cel.RootPath + "/migrations")

		if err != nil {
			exitGraceFully(err)
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), args3)
		upFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)

		if err != nil {
			exitGraceFully(err)
		}

		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)

		if err != nil {
			exitGraceFully(err)
		}
	}

	return nil
}
