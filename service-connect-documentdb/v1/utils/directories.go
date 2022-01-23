package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadingFiles files get example
func ReadingFiles() error {
	// "bin"
	dirNow, err := os.Getwd()

	if err != nil {
		return err
	}

	filesW, _ := ioutil.ReadDir(dirNow)

	for _, file := range filesW {
		fmt.Println(file.Name(), "names 2")
	}

	files, err := ioutil.ReadDir("./")

	for _, file := range files {
		fmt.Println(file.Name(), "names")
	}

	// now file reading
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fmt.Println(dir, "dir now actual file")

	return nil
}
