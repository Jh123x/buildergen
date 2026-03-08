package writer

import (
	"os"
)

func WriteToSingleFile(destination, builderCode string) error {
	file, err := os.Create(destination)
	if err != nil {
		return err
	}

	file.WriteString(builderCode)
	if err := file.Close(); err != nil {
		return err
	}

	return nil
}
