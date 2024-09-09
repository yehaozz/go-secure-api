package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yehaozz/go-secure-api/models"
	"k8s.io/utils/ptr"
)

var (
	testSongs     map[string]models.Song
	idOne         string
	idTwo         string
	idThree       string
	nonExistingID string
)

func init() {
	idOne = generateID()
	idTwo = generateID()
	idThree = generateID()
	nonExistingID = generateID()

	testSongs = map[string]models.Song{
		idOne:   {ID: idOne, Title: "September", Artist: "Earth, Wind & Fire", Rating: ptr.To(4.6)},
		idTwo:   {ID: idTwo, Title: "Fun For All", Artist: "Vinida", Rating: ptr.To(4.7)},
		idThree: {ID: idThree, Title: "Coco Elva Tia", Artist: "MaSiWei", Rating: ptr.To(4.7)},
	}

	fmt.Printf("test songs' IDs are: %v\n", []string{idOne, idTwo, idThree})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/songs", GetSongs)
	r.GET("/songs/:id", GetSong)
	r.POST("/songs", PostSong)
	r.PUT("/songs/:id", UpdateSong)
	r.DELETE("/songs/:id", DeleteSong)
	return r
}

func populateSongs() {
	for id, song := range testSongs {
		songs[id] = song
	}
}

func TestGetSongs(t *testing.T) {
	router := setupRouter()
	populateSongs()

	req, _ := http.NewRequest("GET", "/songs", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Parse the response body
	var responseSongs []models.Song
	if err := json.Unmarshal(w.Body.Bytes(), &responseSongs); err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}

	// Convert testSongs map to a slice for comparison
	var expectedSongs []models.Song
	for _, song := range testSongs {
		expectedSongs = append(expectedSongs, song)
	}

	// Compare the response with the expected songs
	if len(responseSongs) != len(expectedSongs) {
		t.Errorf("Expected %d songs, got %d", len(expectedSongs), len(responseSongs))
	}
	for _, expectedSong := range expectedSongs {
		fmt.Printf("finding song: %s\n", expectedSong.ID)
		found := false
		for _, responseSong := range responseSongs {
			if responseSong.ID == expectedSong.ID {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected song %v not found in response", expectedSong)
		}
	}
}

func TestGetSong(t *testing.T) {
	router := setupRouter()
	populateSongs()

	// Test getting the song with idOne
	req, _ := http.NewRequest("GET", fmt.Sprintf("/songs/%s", idOne), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Parse the response body
	var responseSong models.Song
	if err := json.Unmarshal(w.Body.Bytes(), &responseSong); err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}

	// Verify the details of the responseSong
	if !models.IsSameSong(songs[idOne], responseSong) {
		t.Fatalf("Expect song: %v, got song: %v", songs[idOne], responseSong)
	}

	// Test getting a non-existing song
	req, _ = http.NewRequest("GET", fmt.Sprintf("/songs/%s", nonExistingID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestPostSong(t *testing.T) {
	router := setupRouter()
	populateSongs()

	// Prepare new song data
	newSong := models.Song{Title: "New Song", Artist: "New Artist", Rating: ptr.To(4.5)}
	jsonData, _ := json.Marshal(newSong)

	req, _ := http.NewRequest("POST", "/songs", bytes.NewReader(jsonData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	// Parse the response body to check the newly created song
	var createdSong models.Song
	if err := json.Unmarshal(w.Body.Bytes(), &createdSong); err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}

	if createdSong.Title != newSong.Title || createdSong.Artist != newSong.Artist || *createdSong.Rating != *newSong.Rating {
		t.Errorf("Expected song %v, got %v", newSong, createdSong)
	}

	// Check if the song was added to the songs map
	if _, exists := songs[createdSong.ID]; !exists {
		t.Errorf("Expected song not found in songs map")
	}
}

func TestUpdateSong(t *testing.T) {
	router := setupRouter()
	populateSongs()

	// Prepare updated song data
	updatedTitle := "Updated Title"
	updatedArtist := "Updated Artist"
	updatedRating := 4.8
	updatedSong := models.Song{ID: idOne, Title: updatedTitle, Artist: updatedArtist, Rating: ptr.To(updatedRating)}
	jsonData, _ := json.Marshal(updatedSong)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/songs/%s", idOne), bytes.NewReader(jsonData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Parse the response body to check the updated song
	var responseSong models.Song
	if err := json.Unmarshal(w.Body.Bytes(), &responseSong); err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}

	// Check if the song was updated correctly
	if !models.IsSameSong(updatedSong, responseSong) {
		t.Errorf("Expected song %v, got %v", updatedSong, responseSong)
	}

	// Test updating a non-existing song
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/songs/%s", nonExistingID), bytes.NewReader(jsonData))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestDeleteSong(t *testing.T) {
	router := setupRouter()
	populateSongs()

	// Test deleting an existing song
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/songs/%s", idOne), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Check if the song has been deleted
	if _, exists := songs[idOne]; exists {
		t.Errorf("Expected song to be deleted, but found in songs map")
	}

	// Test deleting a non-existing song
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/songs/%s", nonExistingID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}
