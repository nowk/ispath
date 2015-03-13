package ispath

import (
	. "path/filepath"
)

func IsPath(path string) bool {
	dir, _ := Split(path)

	return dir != ""
}
