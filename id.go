package main

import (
	"crypto/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// IDTab struct.
type IDTab struct {
	uuidFunc func() string
	ulidFunc func() string
}

// NewIDTab creates a new IDTab.
func NewIDTab() *IDTab {
	return &IDTab{
		uuidFunc: newUUID,
		ulidFunc: newULID,
	}
}

// GenerateUUID generates a new UUID.
func (t *IDTab) GenerateUUID(upperCase bool) string { //nolint:revive // it's ok
	id := t.uuidFunc()

	if upperCase {
		id = strings.ToUpper(id)
	} else {
		id = strings.ToLower(id)
	}

	return id
}

// GenerateULID generates a new ULID.
func (t *IDTab) GenerateULID(upperCase bool) string { //nolint:revive // it's ok
	id := t.ulidFunc()

	if upperCase {
		id = strings.ToUpper(id)
	} else {
		id = strings.ToLower(id)
	}

	return id
}

func newUUID() string {
	return uuid.New().String()
}

func newULID() string {
	return ulid.MustNew(
		ulid.Now(),
		ulid.Monotonic(rand.Reader, 0),
	).String()
}
