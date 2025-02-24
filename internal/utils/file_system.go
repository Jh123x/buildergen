package utils

import (
	"net/url"
	"os"
	"path/filepath"
)

const (
	GOPATH = "GOPATH"
)

func GetFilenameFromGoPath(src string) string {
	url, err := url.Parse(src)
	if err != nil {
		return ""
	}

	goPath := os.Getenv(GOPATH)
	return filepath.Join(goPath, "pkg", "mod", url.Path)
}
