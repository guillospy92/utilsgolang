package main

func doMigrate(arg2, arg3 string) error {
	dns := prepareDnsMigrations()

	switch arg2 {
	case "up":
		return cel.Migrate(dns)
	case "down":
		if arg3 == "all" {
			err := cel.MigrateDownAll(dns)

			if err != nil {
				return err
			}
			return nil
		}

		return cel.Steps(-1, dns)
	case "reset":
		err := cel.MigrateDownAll(dns)

		if err != nil {
			return err
		}

		return cel.Migrate(dns)
	default:
		showHelp()
	}

	return nil
}
