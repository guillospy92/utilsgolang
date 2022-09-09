package celeritas

import "os"

// CreateDirNotExists create new structure project
func (a *Accelerator) CreateDirNotExists(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateFileNotExists create a new file .env
func (a *Accelerator) CreateFileNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file, err = os.Create(path)

		if err != nil {
			return err
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	return nil
}
