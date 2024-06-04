package imdb_test

import (
	"testing"
)

const (
	cillianMurphyID = "nm0614165"
	invalidPersonID = "invalidID123"
	keanuReevesID   = "nm0000206" // Example valid ID, you can replace it with other valid IDs
)

func TestGetPersons(t *testing.T) {
	personIDs := []string{
		cillianMurphyID,
		invalidPersonID,
		keanuReevesID,
	}

	for _, id := range personIDs {
		t.Run(id, func(t *testing.T) {
			res, err := c.GetPerson(id)
			if err != nil {
				t.Logf("Expected error for ID %s: %v", id, err)
				if id != invalidPersonID {
					t.Errorf("Unexpected error for valid ID %s: %v", id, err)
				}
			} else {
				if id == invalidPersonID {
					t.Errorf("Expected error for invalid ID %s, but got result: %+v", id, res)
				} else {
					t.Logf("Got person data for ID %s", id)
				}
			}
		})
	}
}
