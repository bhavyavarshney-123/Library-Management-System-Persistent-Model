package badgerdb

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const dbFolder = "data"

//Exists return a bool indicating if the database directory is already initialized
func Exists() bool {
	//create path to MANIFEST file in database directory.
	//this MANIFEST file is good indication of whether the database is initialized
	manifest := filepath.Join(dir(), "MANIFEST")

	//Check if the MANIFEST file exists
	if _, err := os.Stat(manifest); errors.Is(err, os.ErrNotExist) {
		//File does not exist,therefore no DB
		return false
	}
	return true
}

//dir returns the path of the directory that contains the database contents.
//always in the same directory as of the project
func dir() string {
	// Get path to executable
	executable, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("database directory detection failure: exec path detection failure: %w", err))
	}

	// Get directory of the executable
	execdir := filepath.Dir(executable)
	// Add dbFolder to return directory of database
	return filepath.Join(execdir, dbFolder)
}
