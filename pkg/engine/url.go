package engine

import (
	"path"
	"strings"
)

// ShiftPath trims the last / and returns the head
// which we can base our routing decision upon
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)

	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}

	return p[1:i], p[i:]
}
