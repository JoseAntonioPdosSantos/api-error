package apierror

import (
	"os"
	"strings"
)

func getByKey(key string) string {
	property := os.Getenv(key)
	if len(strings.TrimSpace(property)) == 0 {
		return ""
	}
	return property
}
