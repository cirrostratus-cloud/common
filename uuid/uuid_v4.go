package uuid

import "crypto/rand"

func NewV4() string {
	uuid := make([]byte, 16)
	rand.Read(uuid)
	return string(uuid)
}
