package content

import (
	"testing"

	"github.com/Nirespire/swn/content/background"
	"github.com/Nirespire/swn/content/name"
)

func TestPcRandomRoll(t *testing.T) {
	pc := PlayerCharacter{}
	pc.RandomRoll(true, name.Any, name.None)

	if pc.Name == "" {
		t.Error("PC Name not initialized", pc.Name)
	}

	if pc.Stats == (Stats{}) {
		t.Error("PC Stats not initialized", pc.Stats)
	}

	if pc.Background == (background.Background{}) {
		t.Error("PC Background not initialized", pc.Background)
	}
}
