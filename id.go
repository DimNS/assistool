package main

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

const (
	oneSec = 1000
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

// ExtractTimestampFromULID extracts the timestamp from a ULID.
func (*IDTab) ExtractTimestampFromULID(id string) string {
	v, err := ulid.Parse(id)
	if err != nil {
		return fmt.Sprintf("Invalid ULID: %v", err)
	}

	millis := v.Time()
	tm := time.Unix(int64(millis)/oneSec, (int64(millis)%oneSec)*int64(time.Millisecond)) //nolint:gosec // it's ok

	return tm.UTC().Format("2006-01-02T15:04:05.000Z")
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
