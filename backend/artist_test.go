package backend

import (
	"testing"
)

func TestArtistsData(t *testing.T) {
	artists, err := ArtistsData("https://groupietrackers.herokuapp.com/api/artists/")
	if err != nil {
		t.Fatalf("Error fetching artists data: %v", err)
	}

	// Check if the slice of artists is not empty
	if len(artists) == 0 {
		t.Fatal("Expected some artists, but got none")
	}

	//Check if the first artist is Queen
	if artists[0].Name != "Queen" {
		t.Errorf("Expected artist name to be Queen, but got %s", artists[0].Name)
	}

	// Check if the artist ID is non-zero
	if artists[0].Id == 0 {
		t.Error("Expected artist ID to be non-zero, but got 0")
	}

	// Check if the third artist's CreationDate matches expectation
	if len(artists) > 2 && artists[2].CreationDate != 1965 {
		t.Errorf("Expected artist creation date to be 1965, but got %d", artists[2].CreationDate)
	}

	// Check if the first artist's first album date is as expected
	if artists[0].FirstAlbum != "14-12-1973" {
		t.Errorf("Expected first album date to be 14-12-1973, but got %s", artists[0].FirstAlbum)
	}
}
