package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/guillospy92/utilsgolang/go-version-laravel/celeritas"
	"os"
)

var cel *celeritas.Accelerator

var errMake = errors.New("make require a subcommand: (migration|model|handler)")

const version = "1.0.0"

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()

	path, err := os.Getwd()

	if err != nil {
		exitGraceFully(err)
	}

	cel = celeritas.NewAccelerator(path)

	// init configuration .env vars
	cel.StartConfig()

	if err != nil {
		exitGraceFully(err)
	}

	switch arg1 {
	case "help":
		showHelp()

	case "version":
		color.Yellow("Application version: " + version)

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGraceFully(err)
		}

		message = "Migrations Complete"

	case "make":
		if arg2 == "" {
			exitGraceFully(errMake)
		}
		err = doMake(arg2, arg3)
		if err != nil {
			exitGraceFully(err)
		}
	default:
		showHelp()
	}

	successMessage(message)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func exitGraceFully(err error, msg ...string) {
	color.Red(err.Error(), msg)
}

func successMessage(msg ...string) {
	if len(msg) > 0 {
		color.Green("message", msg)
	}
}

func showHelp() {
	color.Yellow(`Available commands:
	help                  - show the help commands
	version               - print application version
	migrate               - runs all up migrations that have not been run previously
	migrate down          - reverses the most recent migration
	migrate reset         - runs all down migrations in reverse order, and then all up migrations
	make migration <name> - creates two new up and down migrations in the migrations folder
	`)
}
