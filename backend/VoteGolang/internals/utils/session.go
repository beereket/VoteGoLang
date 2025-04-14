package utils

import (
	"crypto/rand"
	"encoding/binary"
)

func GenerateSessionID() uint {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	sessionID := binary.LittleEndian.Uint64(bytes)

	return uint(sessionID)
}
