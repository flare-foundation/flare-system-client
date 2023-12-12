package utils

import "strings"

// Join (nonempty) path1 and path2 with inserting "/" between them if necessary.
// Use net/url JoinPath when upgrading to go 1.9.x
func JoinPaths(path1, path2 string) string {
	if strings.HasSuffix(path1, "/") {
		return path1 + path2
	} else {
		return path1 + "/" + path2
	}
}
