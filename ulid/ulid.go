package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func New() string {
	// Configura una fuente de números aleatorios
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crea un nuevo ULID
	ulid := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)

	// Convierte el ULID a una cadena
	return ulid.String()
}
