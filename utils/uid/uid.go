package uid

import (
	"strings"

	"github.com/google/uuid"
)

func Generate() (uid string) {
	uid = strings.ReplaceAll(uuid.New().String(), "-", "")
	return
}
