package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func importPathFromFile(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	if ip, err := importPathFromModule(absPath); err == nil {
		return ip, nil
	}

	if ip, err := importPathFromGopath(absPath); err == nil {
		return ip, nil
	}

	return "", fmt.Errorf("unable to determine import path for: %s", filePath)
}

func importPathFromModule(file string) (string, error) {
	abs, _ := filepath.Abs(file)
	dir := filepath.Dir(abs)

	modRoot, err := findModuleRoot(dir)
	if err != nil {
		return "", err
	}

	modPath, err := readModulePath(modRoot)
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(modRoot, dir)
	if err != nil {
		return "", err
	}

	if rel == "." {
		return modPath, nil
	}
	return path.Join(modPath, filepath.ToSlash(rel)), nil
}

func readModulePath(modRoot string) (string, error) {
	data, err := os.ReadFile(filepath.Join(modRoot, "go.mod"))
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", fmt.Errorf("module path not found")
}

func findModuleRoot(dir string) (string, error) {
	cur := dir
	for {
		if _, err := os.Stat(filepath.Join(cur, "go.mod")); err == nil {
			return cur, nil
		}

		parent := filepath.Dir(cur)
		if parent == cur {
			break
		}
		cur = parent
	}
	return "", fmt.Errorf("no module root found")
}

func importPathFromGopath(file string) (string, error) {
	abs, _ := filepath.Abs(file)
	for _, gopath := range filepath.SplitList(build.Default.GOPATH) {
		src := filepath.Join(gopath, "src") + string(os.PathSeparator)
		if strings.HasPrefix(abs, src) {
			rel := strings.TrimPrefix(filepath.Dir(abs), src)
			return filepath.ToSlash(rel), nil
		}
	}
	return "", fmt.Errorf("file is not in a module or GOPATH")
}
