package str

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
