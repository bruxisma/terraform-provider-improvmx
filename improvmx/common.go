package improvmx

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

// This is based off of (but not a direct copy of!) how cloudflare combines ids
func generateID(ids []string) string {
	sort.Strings(ids)
	hash := sha256.New()
	hash.Write([]byte(strings.Join(ids, "")))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
