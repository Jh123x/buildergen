package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	tests := map[string]struct {
		goPath      string
		src         string
		expectedRes string
	}{
		"go mod pkg": {
			goPath:      "/home/user/go",
			src:         "github.com/Jh123x/buildergen",
			expectedRes: filepath.Join("/", "home", "user", "go", "pkg", "mod", "github.com", "Jh123x", "buildergen"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer func(s string) { os.Setenv(GOPATH, s) }(os.Getenv(GOPATH))
			os.Setenv(GOPATH, tc.goPath)

			res := GetFilenameFromGoPath(tc.src)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
