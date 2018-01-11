package mongo

import "strings"

func formatName(name string) string {
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	return name
}
