package main

import (
	"os"

	"github.com/pkg/errors"
)

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return errors.Wrap(err, "Could not create folder")
		}
	}
	return nil
}
