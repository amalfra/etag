package etag

import (
	"crypto/sha1"
	"fmt"
)

func getHash(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

// Generates an Etag for given string. Allows specifying whether to generate weak
// Etag or not as second parameter
func Generate(str string, weak bool) string {
	tag := fmt.Sprintf("\"%d-%s\"", len(str), getHash(str))
	if weak {
		tag = "W/" + tag
	}

	return tag
}
