package etag

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func getHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// Generate an Etag for given sring. Allows specifying whether to generate weak
// Etag or not as second parameter
func Generate(str string, weak bool) string {
	tag := fmt.Sprintf("%d-%s", len(str), getHash(str))
	if weak {
		tag = "W/" + tag
	}

	return tag
}
