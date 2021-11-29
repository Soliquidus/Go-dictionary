package dictionary

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"time"
)

type Dictionary struct {
	db *badger.DB
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}
