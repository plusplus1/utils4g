package str

import (
	"strings"
)

import (
	"github.com/satori/go.uuid"
)

func (u *Util) NewStrUUID() string {
	return strings.Replace(uuid.NewV1().String(), "-", "", -1)
}
