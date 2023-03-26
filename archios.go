package archio_domis

import (
	"encoding/binary"
	"os"
	"sync"
)

var (
	fileLocks map[string]*sync.Mutex
)

func init() {
	fileLocks = make(map[string]*sync.Mutex)
}

// Save() saves any struct into a file on file system. This file can later be loaded into a struct.
// Save() function is safe to use concurrently.
func Save(path string, v interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, v)
	if err != nil {
		return err
	}
	return nil
}

// Load() loads the data back from file into struct. If file not found it returns back FileNotFound Error.
func Load(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = binary.Read(file, binary.LittleEndian, v)
	return err
}
