package selector

import (
	"log"
	"testing"
)

func TestEmpty(t *testing.T) {
	e := Empty()

	if e.stype != NONE || e.name != "" {
		log.Fatalf("Selector Empty function should be empty, but is %v", e)
	}
}
