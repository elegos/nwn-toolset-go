package tools

import (
	"os"
	"path/filepath"
)

// RemoveDirectory removes a directory recursively
// https://stackoverflow.com/a/33451503
func RemoveDirectory(directory string) error {
	d, err := os.Open(directory)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(directory, name))
		if err != nil {
			return err
		}
	}

	os.RemoveAll(directory)

	return nil
}
