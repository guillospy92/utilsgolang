package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath, targetFile string) error {
	// check to ensure file does not already exist
	data, err := templateFS.ReadFile(templatePath)

	if err != nil {
		exitGraceFully(err)
	}

	err = copyDateToFile(data, targetFile)

	if err != nil {
		exitGraceFully(err)
	}

	return nil
}

func copyDateToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)

	return err
}
